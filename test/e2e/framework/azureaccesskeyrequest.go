package framework

import (
	"kubevault.dev/operator/apis"
	api "kubevault.dev/operator/apis/engine/v1alpha1"
	patchutil "kubevault.dev/operator/client/clientset/versioned/typed/engine/v1alpha1/util"
)

func (f *Framework) UpdateAzureAccessKeyRequestStatus(status *api.AzureAccessKeyRequestStatus, azureAReq *api.AzureAccessKeyRequest) error {
	_, err := patchutil.UpdateAzureAccessKeyRequestStatus(f.CSClient.EngineV1alpha1(), azureAReq, func(s *api.AzureAccessKeyRequestStatus) *api.AzureAccessKeyRequestStatus {
		s = status
		return s
	}, apis.EnableStatusSubresource)
	return err
}
