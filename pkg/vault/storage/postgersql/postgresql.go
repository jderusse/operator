package postgresql

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	api "kubevault.dev/operator/apis/kubevault/v1alpha1"
)

var postgresqlStorageFmt = `
storage "postgresql" {
%s
}
`

type Options struct {
	api.PostgreSQLSpec

	// the connection string to use to authenticate and connect to PostgreSQL.
	ConnectionUrl string
}

func NewOptions(kubeClient kubernetes.Interface, ns string, s api.PostgreSQLSpec) (*Options, error) {
	var (
		url string
	)

	if s.ConnectionUrlSecret != "" {
		sr, err := kubeClient.CoreV1().Secrets(ns).Get(s.ConnectionUrlSecret, metav1.GetOptions{})
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get connection url secret(%s)", s.ConnectionUrlSecret)
		}

		if value, ok := sr.Data["connection_url"]; ok {
			url = string(value)
		} else {
			return nil, errors.Errorf("connection_url not found in secret(%s)", s.ConnectionUrlSecret)
		}

	}
	return &Options{
		s,
		url,
	}, nil
}

func (o *Options) Apply(pt *core.PodTemplateSpec) error {
	return nil
}

// vault doc: https://www.vaultproject.io/docs/configuration/storage/postgresql.html
//
// GetGcsConfig creates postgresql storae config from GcsSpec
func (o *Options) GetStorageConfig() (string, error) {
	params := []string{}
	if o.ConnectionUrlSecret != "" {
		params = append(params, fmt.Sprintf(`connection_url = "%s"`, o.ConnectionUrl))
	}
	if o.Table != "" {
		params = append(params, fmt.Sprintf(`table = "%s"`, o.Table))
	}
	if o.MaxParallel != 0 {
		params = append(params, fmt.Sprintf(`max_parallel = "%d"`, o.MaxParallel))
	}

	storageCfg := fmt.Sprintf(postgresqlStorageFmt, strings.Join(params, "\n"))
	return storageCfg, nil
}
