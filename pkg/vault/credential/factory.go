package credential

import (
	dbapi "github.com/kubedb/apimachinery/apis/authorization/v1alpha1"
	dbcrd "github.com/kubedb/apimachinery/client/clientset/versioned"
	engineapi "github.com/kubevault/operator/apis/engine/v1alpha1"
	vaultcrd "github.com/kubevault/operator/client/clientset/versioned"
	"github.com/kubevault/operator/pkg/vault/credential/aws"
	"github.com/kubevault/operator/pkg/vault/credential/database"
	"github.com/kubevault/operator/pkg/vault/credential/gcp"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	appcat_cs "kmodules.xyz/custom-resources/client/clientset/versioned/typed/appcatalog/v1alpha1"
)

func NewCredentialManagerForDatabase(kubeClient kubernetes.Interface, appClient appcat_cs.AppcatalogV1alpha1Interface, cr dbcrd.Interface, dbAR *dbapi.DatabaseAccessRequest) (CredentialManager, error) {
	dbCM, err := database.NewDatabaseCredentialManager(kubeClient, appClient, cr, dbAR)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get database credential manager")
	}
	return &CredManager{
		kubeClient:   kubeClient,
		secretEngine: dbCM,
		vaultClient:  dbCM.VaultClient,
	}, nil
}

func NewCredentialManagerForAWS(kubeClient kubernetes.Interface, appClient appcat_cs.AppcatalogV1alpha1Interface, cr vaultcrd.Interface, awsAKReq *engineapi.AWSAccessKeyRequest) (CredentialManager, error) {
	awsCM, err := aws.NewAWSCredentialManager(kubeClient, appClient, cr, awsAKReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get aws credential manager")
	}
	return &CredManager{
		kubeClient:   kubeClient,
		secretEngine: awsCM,
		vaultClient:  awsCM.VaultClient,
	}, nil
}

func NewCredentialManagerForGCP(kubeClient kubernetes.Interface, appClient appcat_cs.AppcatalogV1alpha1Interface, cr vaultcrd.Interface, gcpAKReq *engineapi.GCPAccessKeyRequest) (CredentialManager, error) {
	gcpCM, err := gcp.NewGCPCredentialManager(kubeClient, appClient, cr, gcpAKReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get gcp credential manager")
	}
	return &CredManager{
		kubeClient:   kubeClient,
		secretEngine: gcpCM,
		vaultClient:  gcpCM.VaultClient,
	}, nil
}
