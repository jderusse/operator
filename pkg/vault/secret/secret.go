package secret

import (
	vaultapi "github.com/hashicorp/vault/api"
)

const (
	RoleKey      = "role"
	PathKey      = "path"
	SecretKey    = "secret"
	KeyAlgorithm = "key_algorithm"
	KeyType      = "key_type"
)

type SecretManager interface {
	SecretGetter
	SetOptions(*vaultapi.Client, map[string]string) error
}

type SecretGetter interface {
	GetSecret() (*vaultapi.Secret, error)
}
