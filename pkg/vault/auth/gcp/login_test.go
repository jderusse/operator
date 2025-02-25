package gcp

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
)

func TestLogin(t *testing.T) {
	addr := os.Getenv("VAULT_ADDR")
	credentialaddr := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	role := os.Getenv("VAULT_GCP_ROLE_NAME")
	if addr == "" || credentialaddr == "" || role == "" {
		t.Skip()
	}

	jsonBytes, err := ioutil.ReadFile(credentialaddr)
	if err != nil {
		log.Fatal(err)
	}

	au, err := New(&appcat.AppBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "gcp",
			Namespace: "default",
		},
		Spec: appcat.AppBindingSpec{
			ClientConfig: appcat.ClientConfig{
				URL:                   &addr,
				InsecureSkipTLSVerify: true,
			},
			Secret: &corev1.LocalObjectReference{
				Name: "gcp",
			},
			Parameters: &runtime.RawExtension{
				Raw: []byte(fmt.Sprintf(`{ "PolicyControllerRole" : "%s" }`, role)),
			},
		},
	}, &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "gcp",
			Namespace: "default",
		},
		Data: map[string][]byte{
			"sa.json": []byte(jsonBytes),
		},
	})

	if err != nil {
		log.Println("New failed!")
	}
	if au == nil {
		log.Println("au nil!")
		t.Skip()
	}

	if au.signedJwt == "" || au.role == "" {
		t.Skip()
	}
	token, err := au.Login()
	if assert.Nil(t, err) {
		fmt.Println(token)
	}
}
