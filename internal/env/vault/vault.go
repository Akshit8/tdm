// Package vault implements types and methods to interact with vault server
package vault

import "github.com/hashicorp/vault/api"

type vaultProvider struct {
	path string
	client *api.Logical
}