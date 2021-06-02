package rest_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Akshit8/tdm/internal/rest"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
)

// func TestNewDates(t *testing.T) {
// 	t.Parallel()

// 	tests := []struct {
// 		name   string
// 		input  internal.Dates
// 		output rest.Dates
// 	}{
// 		{
// 			"OK",
// 			internal.Dates{
// 				Start: time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC),
// 				Due:   time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC).Add(time.Hour),
// 			},
// 			rest.Dates{
// 				Start: rest.Time(time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)),
// 				Due:   rest.Time(time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC).Add(time.Hour)),
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test

// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()

// 			res := rest.NewDates(test.input)

// 			if !cmp.Equal(test.output, res, cmpopts.IgnoreUnexported(rest.Time{})) {
// 				t.Fatalf("expected output do not match\n%s", cmp.Diff(test.output, res, cmpopts.IgnoreUnexported(rest.Time{})))
// 			}
// 		})
// 	}
// }

func TestTimeMarshal(t *testing.T) {
	t.Parallel()

	type output struct {
		res     []byte
		withErr bool
	}

	tests := []struct {
		name   string
		input  rest.Time
		output output
	}{
		{
			"OK",
			rest.Time(time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)),
			output{
				res: []byte(`"2009-11-10T23:00:00Z"`),
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

func TestTimeUnmarshal(t *testing.T) {
	t.Parallel()

	type output struct {
		res     rest.Time
		withErr bool
	}

	tests := []struct {
		name   string
		input  []byte
		output output
	}{
		{
			"OK",
			[]byte(`"2009-11-10T23:00:00Z"`),
			output{
				res: rest.Time(time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)),
			},
		},
		{
			"ERR",
			[]byte(`"2009-"`),
			output{
				withErr: true,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var res rest.Time

			err := json.Unmarshal(test.input, &res)

			isErr := err != nil
			require.Equal(t, test.output.withErr, isErr)

			if !cmp.Equal(test.output.res, res, cmpopts.IgnoreUnexported(rest.Time{})) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(test.output.res, res, cmpopts.IgnoreUnexported(rest.Time{})))
			}
		})
	}
}

func TestDateMarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  rest.Dates
		output []byte
	}{
		{
			"OK",
			rest.Dates{
				Start: rest.Time(time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)),
			},
			[]byte(`{"start":"2009-11-10T23:00:00Z","due":"0001-01-01T00:00:00Z"}`),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			res, err := json.Marshal(test.input)
			require.NoError(t, err)

			if !cmp.Equal(test.output, res) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(test.output, res))
			}
		})
	}
}
