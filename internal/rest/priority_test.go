package rest_test

import (
	"encoding/json"
	"testing"

	"github.com/Akshit8/tdm/internal"
	"github.com/Akshit8/tdm/internal/rest"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestNewPriority(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  internal.Priority
		output rest.Priority
	}{
		{
			"OK: none",
			internal.PriorityNone,
			rest.Priority("none"),
		},
		{
			"OK: low",
			internal.PriorityLow,
			rest.Priority("low"),
		},
		{
			"OK: medium",
			internal.PriorityMedium,
			rest.Priority("medium"),
		},
		{
			"OK: high",
			internal.PriorityHigh,
			rest.Priority("high"),
		},
		{
			"OK: unknown",
			internal.Priority(-1),
			rest.Priority("none"),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			res := rest.NewPriority(test.input)

			if !cmp.Equal(test.output, res) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(test.output, res))
			}

		})
	}
}

func TestConvert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		input rest.Priority
		output internal.Priority
	}{
		{
			"OK: none",
			rest.Priority("none"),
			internal.PriorityNone,
		},
		{
			"OK: low",
			rest.Priority("low"),
			internal.PriorityLow,
		},
		{
			"OK: medium",
			rest.Priority("medium"),
			internal.PriorityMedium,
		},
		{
			"OK: high",
			rest.Priority("high"),
			internal.PriorityHigh,
		},
		{
			"OK: unknown",
			rest.Priority("unknown"),
			internal.PriorityNone,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			res := test.input.Convert()

			if !cmp.Equal(test.output, res) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(test.output, res))
			}

		})
	}
}

func TestPriorityMarshalJSON(t *testing.T) {
	t.Parallel()

	type output struct {
		res []byte
		withErr bool
	}

	tests := []struct {
		name string
		input rest.Priority
		output output
	}{
		{
			"OK",
			rest.Priority("none"),
			output{
				res: []byte(`"none"`),
			},
		},
		{
			"ERR",
			rest.Priority("unknown"),
			output{
				withErr: true,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			res, err := json.Marshal(test.input)

			isErr := err != nil
			require.Equal(t, test.output.withErr, isErr)

			if !cmp.Equal(test.output.res, res) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(test.output.res, res))
			}

		})
	}
}

func TestPriorityUnmarshalJSON(t *testing.T) {
	t.Parallel()

	type output struct {
		res rest.Priority
		withErr bool
	}

	tests := []struct {
		name string
		input []byte
		output output
	}{
		{
			"OK",
			[]byte(`"none"`),
			output{
				res: rest.Priority("none"),
			},
		},
		{
			"ERR",
			[]byte(`"unknown`),
			output{
				withErr: true,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var res rest.Priority

			err := json.Unmarshal(test.input, &res)

			isErr := err != nil
			require.Equal(t, test.output.withErr, isErr)

			if !cmp.Equal(test.output.res, res) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(test.output.res, res))
			}

		})
	}
}