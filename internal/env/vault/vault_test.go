package vault_test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/Akshit8/tdm/internal/env"
	"github.com/Akshit8/tdm/internal/env/vault"
	"github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/require"
)

type vaultClient struct {
	Token   string
	Address string
	path    string
	Client  *api.Client
}

func newVaultClient(t *testing.T) *vaultClient {
	err := env.Load(path.Join("../../../", ".test.env"))
	require.NoError(t, err)

	token := os.Getenv("VAULT_TOKEN_TEST")
	addr := os.Getenv("VAULT_ADDRESS_TEST")
	path := os.Getenv("VAULT_PATH_TEST")

	config := &api.Config{
		Address: addr,
	}

	client, err := api.NewClient(config)
	require.NoError(t, err)

	client.SetToken(token)

	return &vaultClient{
		Token:   token,
		Address: addr,
		path:    path,
		Client:  client,
	}
}

func TestGet(t *testing.T) {
	t.Parallel()

	type output struct {
		res     string
		withErr bool
	}

	tests := []struct {
		name   string
		setup  func(*testing.T, *vaultClient)
		input  string
		output output
	}{
		{
			"OK",
			func(t *testing.T, vc *vaultClient) {
				_, err := vc.Client.Logical().Write(
					fmt.Sprintf("%s/data/ok", vc.path),
					map[string]interface{}{
						"data": map[string]interface{}{
							"one": "1",
							"two": "2",
						},
					},
				)
				require.NoError(t, err)
			},
			"/ok:one",
			output{
				res: "1",
			},
		},
		{
			"OK: cached",
			func(t *testing.T, vc *vaultClient) {},
			"/ok:two",
			output{
				res: "2",
			},
		},
		{
			"ERR: key not found in cached data",
			func(t *testing.T, vc *vaultClient) {},
			"/ok:three",
			output{
				withErr: true,
			},
		},
		{
			"ERR: secret not found",
			func(t *testing.T, vc *vaultClient) {},
			"/not:found",
			output{
				withErr: true,
			},
		},
		{
			"ERR: key not found in retreives data",
			func(t *testing.T, vc *vaultClient) {
				_, err := vc.Client.Logical().Write(
					fmt.Sprintf("%s/data/err", vc.path),
					map[string]interface{}{
						"data": map[string]interface{}{
							"hello": "world",
						},
					},
				)
				require.NoError(t, err)
			},
			"/err:something",
			output{
				withErr: true,
			},
		},
	}

	// Provider is not local to the subtest because we want to test the local caching logic

	client := newVaultClient(t)
	provider, err := vault.NewVaultProvider(client.Token, client.Address, client.path)
	require.NoError(t, err)

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			// Not calling t.Parallel() because vault.Provider is not goroutine safe.

			test.setup(t, client)

			res, err := provider.Get(test.input)

			isErr := err != nil
			require.Equal(t, test.output.withErr, isErr)
			require.Equal(t, test.output.res, res)
		})
	}
}
