package env_test

import (
	"errors"
	"os"
	"path"
	"testing"

	"github.com/Akshit8/tdm/internal/env"
	"github.com/Akshit8/tdm/internal/env/envtesting"
	"github.com/stretchr/testify/require"
)

func TestConfigurationGet(t *testing.T) {
	t.Parallel()

	type output struct {
		val     string
		withErr bool
	}

	tests := []struct {
		name   string
		setup  func(p *envtesting.FakeProvider) (teardown func())
		input  string
		output output
		arg    string
	}{
		{
			"OK: from env file",
			func(_ *envtesting.FakeProvider) func() {
				return func() {
					os.Setenv("KEY_1", "")
				}
			},
			"KEY_1",
			output{
				val: "VALUE_1",
			},
			"",
		},
		{
			"OK: from provider",
			func(p *envtesting.FakeProvider) func() {
				os.Setenv("KEY_2_SECURE", "/secret/value")

				p.GetReturns("provider value", nil)

				return func() {
					os.Setenv("KEY_2", "")
					os.Setenv("KEY_2_SECURE", "")
				}
			},
			"KEY_2",
			output{
				val: "provider value",
			},
			"/secret/value",
		},
		{
			"ERR: provider failed",
			func(p *envtesting.FakeProvider) func() {
				os.Setenv("KEY_ERR_SECURE", "/failed")

				p.GetReturns("", errors.New("failed"))

				return func() {
					os.Setenv("KEY_ERR_SECURE", "")
				}
			},
			"KEY_ERR",
			output{
				withErr: true,
			},
			"/failed",
		},
	}

	err := env.Load(path.Join("fixtures", "no file"))
	require.Error(t, err)

	err = env.Load(path.Join("fixtures", ".env"))
	require.NoError(t, err)

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			provider := envtesting.FakeProvider{}

			teardown := test.setup(&provider)
			t.Cleanup(teardown)

			val, err := env.NewConfiguration(&provider).Get(test.input)

			isErr := err != nil
			require.Equal(t, test.output.withErr, isErr)

			require.Equal(t, test.output.val, val)

			if provider.GetCallCount() > 0 {
				arg := provider.GetArgsForCall(0)
				require.Equal(t, test.arg, arg)
			}
		})
	}
}
