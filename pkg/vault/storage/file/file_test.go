package file

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	api "kubevault.dev/operator/apis/kubevault/v1alpha1"
)

func TestOptions_GetStorageConfig(t *testing.T) {
	opts, err := NewOptions(api.FileSpec{
		Path: "/test",
	})
	assert.Nil(t, err)

	out := `
storage "file" {
path = "/test"
}
`
	t.Run("file system storage config", func(t *testing.T) {
		got, err := opts.GetStorageConfig()
		assert.Nil(t, err)
		if !assert.Equal(t, out, got) {
			fmt.Println("expected:", out)
			fmt.Println("got:", got)
		}
	})
}
