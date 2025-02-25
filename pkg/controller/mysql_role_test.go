package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/appscode/go/encoding/json/types"
	"github.com/appscode/pat"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kfake "k8s.io/client-go/kubernetes/fake"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	api "kubedb.dev/apimachinery/apis/authorization/v1alpha1"
	dbfake "kubedb.dev/apimachinery/client/clientset/versioned/fake"
	dbinformers "kubedb.dev/apimachinery/client/informers/externalversions"
	"kubevault.dev/operator/pkg/vault/role/database"
)

func setupVaultServerForMysql() *httptest.Server {
	m := pat.New()

	m.Del("/v1/database/roles/m-read", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	m.Del("/v1/database/roles/error", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error"))
	}))

	return httptest.NewServer(m)
}

func TestUserManagerController_reconcileMySQLRole(t *testing.T) {
	mRole := api.MySQLRole{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "pg-role",
			Namespace:  "pg",
			Generation: 0,
		},
		Spec: api.MySQLRoleSpec{
			DatabaseRef: &corev1.LocalObjectReference{
				Name: "test",
			},
			AuthManagerRef: &appcat.AppReference{},
		},
	}

	testData := []struct {
		testName           string
		mRole              api.MySQLRole
		dbRClient          database.DatabaseRoleInterface
		hasStatusCondition bool
		expectedErr        bool
	}{
		{
			testName:           "initial stage, no error",
			mRole:              mRole,
			dbRClient:          &fakeDRole{},
			expectedErr:        false,
			hasStatusCondition: false,
		},
		{
			testName: "initial stage, failed to enable database",
			mRole:    mRole,
			dbRClient: &fakeDRole{
				errorOccurredInEnableDatabase: true,
			},
			expectedErr:        true,
			hasStatusCondition: true,
		},
		{
			testName: "initial stage, failed to create database connection config",
			mRole:    mRole,
			dbRClient: &fakeDRole{
				errorOccurredInCreateConfig: true,
			},
			expectedErr:        true,
			hasStatusCondition: true,
		},
		{
			testName: "initial stage, failed to create database role",
			mRole:    mRole,
			dbRClient: &fakeDRole{
				errorOccurredInCreateRole: true,
			},
			expectedErr:        true,
			hasStatusCondition: true,
		},
		{
			testName: "update role, successfully updated database role",
			mRole: func(p api.MySQLRole) api.MySQLRole {
				p.Generation = 2
				types.IntHashForGeneration(1)
				return p
			}(mRole),
			dbRClient:          &fakeDRole{},
			expectedErr:        false,
			hasStatusCondition: false,
		},
		{
			testName: "update role, failed to update database role",
			mRole: func(p api.MySQLRole) api.MySQLRole {
				p.Generation = 2
				p.Status.ObservedGeneration = types.IntHashForGeneration(1)
				return p
			}(mRole),
			dbRClient: &fakeDRole{
				errorOccurredInCreateRole: true,
			},
			expectedErr:        true,
			hasStatusCondition: true,
		},
	}

	for _, test := range testData {
		t.Run(test.testName, func(t *testing.T) {
			c := &VaultController{
				kubeClient: kfake.NewSimpleClientset(),
				dbClient:   dbfake.NewSimpleClientset(),
			}
			c.dbInformerFactory = dbinformers.NewSharedInformerFactory(c.dbClient, time.Minute*10)

			_, err := c.dbClient.AuthorizationV1alpha1().MySQLRoles(test.mRole.Namespace).Create(&test.mRole)
			if !assert.Nil(t, err) {
				return
			}

			err = c.reconcileMySQLRole(test.dbRClient, &test.mRole)
			if test.expectedErr {
				if assert.NotNil(t, err) {
					if test.hasStatusCondition {
						p, err2 := c.dbClient.AuthorizationV1alpha1().MySQLRoles(test.mRole.Namespace).Get(test.mRole.Name, metav1.GetOptions{})
						if assert.Nil(t, err2) {
							assert.Condition(t, func() (success bool) {
								if len(p.Status.Conditions) == 0 {
									return false
								}
								return true
							}, "should have status.conditions")
						}
					}
				}
			} else {
				if assert.Nil(t, err) {
					p, err2 := c.dbClient.AuthorizationV1alpha1().MySQLRoles(test.mRole.Namespace).Get(test.mRole.Name, metav1.GetOptions{})
					if assert.Nil(t, err2) {
						assert.Condition(t, func() (success bool) {
							if len(p.Status.Conditions) != 0 {
								return false
							}
							return true
						}, "should not have status.conditions")
					}
				}
			}
		})
	}

}
