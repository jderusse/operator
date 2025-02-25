package gcs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	api "kubevault.dev/operator/apis/kubevault/v1alpha1"
)

func TestOptions_GetStorageConfig(t *testing.T) {
	opts, err := NewOptions(api.GcsSpec{
		Bucket:      "test",
		MaxParallel: 128,
		ChunkSize:   "256",
		HAEnabled:   true,
	})
	assert.Nil(t, err)

	out := `
storage "gcs" {
bucket = "test"
ha_enabled = "true"
chunk_size = "256"
max_parallel = 128
}
`
	t.Run("Gcs storage config", func(t *testing.T) {
		got, err := opts.GetStorageConfig()
		assert.Nil(t, err)
		if !assert.Equal(t, out, got) {
			fmt.Println("expected:", out)
			fmt.Println("got:", got)
		}
	})
}
