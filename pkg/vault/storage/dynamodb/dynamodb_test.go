package dynamodb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	api "kubevault.dev/operator/apis/kubevault/v1alpha1"
)

func TestOptions_GetStorageConfig(t *testing.T) {
	opts, err := NewOptions(api.DynamoDBSpec{
		EndPoint:      "endpoint",
		HaEnabled:     true,
		MaxParallel:   128,
		Region:        "us",
		ReadCapacity:  1,
		WriteCapacity: 1,
		Table:         "vault",
	})
	assert.Nil(t, err)

	out := `
storage "dynamodb" {
endpoint = "endpoint"
ha_enabled = "true"
region = "us"
read_capacity = 1
write_capacity = 1
table = "vault"
max_parallel = 128
}
`
	t.Run("DynamoDB storage config", func(t *testing.T) {
		got, err := opts.GetStorageConfig()
		assert.Nil(t, err)
		if !assert.Equal(t, out, got) {
			fmt.Println("expected:", out)
			fmt.Println("got:", got)
		}
	})
}
