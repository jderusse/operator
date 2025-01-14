package controller

import (
	"time"

	pcm "github.com/coreos/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	core "k8s.io/api/core/v1"
	crd_cs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	reg_util "kmodules.xyz/client-go/admissionregistration/v1beta1"
	"kmodules.xyz/client-go/discovery"
	appcat_cs "kmodules.xyz/custom-resources/client/clientset/versioned/typed/appcatalog/v1alpha1"
	db_cs "kubedb.dev/apimachinery/client/clientset/versioned"
	dbinformers "kubedb.dev/apimachinery/client/informers/externalversions"
	cs "kubevault.dev/operator/client/clientset/versioned"
	vaultinformers "kubevault.dev/operator/client/informers/externalversions"
	"kubevault.dev/operator/pkg/eventer"
)

const (
	mutatingWebhook   = "mutators.kubevault.com"
	validatingWebhook = "validators.kubevault.com"
)

type config struct {
	EnableRBAC              bool
	DockerRegistry          string
	MaxNumRequeues          int
	NumThreads              int
	ResyncPeriod            time.Duration
	EnableValidatingWebhook bool
	EnableMutatingWebhook   bool
}

type Config struct {
	config

	ClientConfig     *rest.Config
	KubeClient       kubernetes.Interface
	ExtClient        cs.Interface
	CRDClient        crd_cs.ApiextensionsV1beta1Interface
	AppCatalogClient appcat_cs.AppcatalogV1alpha1Interface
	PromClient       pcm.MonitoringV1Interface
	DbClient         db_cs.Interface
}

func NewConfig(clientConfig *rest.Config) *Config {
	return &Config{
		ClientConfig: clientConfig,
	}
}

func (c *Config) New() (*VaultController, error) {
	if err := discovery.IsDefaultSupportedVersion(c.KubeClient); err != nil {
		return nil, err
	}

	ctrl := &VaultController{
		config:           c.config,
		clientConfig:     c.ClientConfig,
		ctxCancels:       make(map[string]CtxWithCancel),
		finalizerInfo:    NewMapFinalizer(),
		authMethodCtx:    make(map[string]CtxWithCancel),
		kubeClient:       c.KubeClient,
		extClient:        c.ExtClient,
		crdClient:        c.CRDClient,
		promClient:       c.PromClient,
		appCatalogClient: c.AppCatalogClient,
		dbClient:         c.DbClient,
		kubeInformerFactory: informers.NewSharedInformerFactoryWithOptions(
			c.KubeClient,
			c.ResyncPeriod,
			informers.WithNamespace(core.NamespaceAll)),
		extInformerFactory: vaultinformers.NewSharedInformerFactory(c.ExtClient, c.ResyncPeriod),
		dbInformerFactory:  dbinformers.NewSharedInformerFactory(c.DbClient, c.ResyncPeriod),
		recorder:           eventer.NewEventRecorder(c.KubeClient, "vault-operator"),
	}

	if err := ctrl.ensureCustomResourceDefinitions(); err != nil {
		return nil, err
	}
	if c.EnableMutatingWebhook {
		if err := reg_util.UpdateMutatingWebhookCABundle(c.ClientConfig, mutatingWebhook); err != nil {
			return nil, err
		}
	}
	if c.EnableValidatingWebhook {
		if err := reg_util.UpdateValidatingWebhookCABundle(c.ClientConfig, validatingWebhook); err != nil {
			return nil, err
		}
	}

	// For VaultServer
	ctrl.initVaultServerWatcher()
	// For VaultPolicy
	ctrl.initVaultPolicyWatcher()
	// For VaultPolicyBinding
	ctrl.initVaultPolicyBindingWatcher()

	// For DB manager
	ctrl.initPostgresRoleWatcher()
	ctrl.initMySQLRoleWatcher()
	ctrl.initMongoDBRoleWatcher()
	ctrl.initDatabaseAccessWatcher()

	// For AWSRole
	ctrl.initAWSRoleWatcher()
	ctrl.initAWSAccessKeyWatcher()

	// For GCPRole
	ctrl.initGCPRoleWatcher()
	ctrl.initGCPAccessKeyWatcher()

	// For AzureRole
	ctrl.initAzureRoleWatcher()
	ctrl.initAzureAccessKeyWatcher()

	return ctrl, nil
}
