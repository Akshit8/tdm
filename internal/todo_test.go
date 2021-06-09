package internal_test

import (
	"testing"
	"time"

	"github.com/Akshit8/tdm/internal"
	"github.com/stretchr/testify/require"
)

func TestPriorityValidator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   internal.Priority
		withErr bool
	}{
		{
			"OK: PriorityNone",
			internal.PriorityNone,
			false,
		},
		{
			"OK: PriorityLow",
			internal.PriorityLow,
			false,
		},
		{
			"OK: PriorityMedium",
			internal.PriorityMedium,
			false,
		},
		{
			"OK: PriorityHigh",
			internal.PriorityHigh,
			false,
		},
		{
			"ERR: unknown value",
			internal.Priority(-1),
			true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actualErr := test.input.Validate()
			isErr := actualErr != nil
			require.Equal(t, test.withErr, isErr)

			var ierr *internal.Error
			if test.withErr {
				require.ErrorAs(t, actualErr, &ierr)
			}
		})
	}
}

func TestDateValidator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   internal.Dates
		withErr bool
	}{
		{
			"OK: Start is zero",
			internal.Dates{
				Due: time.Now(),
			},
			false,
		},
		{
			"OK: Due is zero",
			internal.Dates{
				Start: time.Now(),
			},
			false,
		},
		{
			"OK: Start < Due",
			internal.Dates{
				Start: time.Now(),
				Due:   time.Now().Add(2 * time.Hour),
			},
			false,
		},
		{
			"OK: Start > Due",
			internal.Dates{
				Start: time.Now().Add(2 * time.Hour),
				Due:   time.Now(),
			},
			true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actualErr := test.input.Validate()
			isErr := actualErr != nil
			require.Equal(t, test.withErr, isErr)

			var ierr *internal.Error
			if test.withErr {
				require.ErrorAs(t, actualErr, &ierr)
			}
		})
	}
}

func TestTaskValidator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   internal.Task
		withErr bool
	}{
		{
			"OK",
			internal.Task{
				Description: "sample description",
				Priority:    internal.PriorityHigh,
				Dates: internal.Dates{
					Start: time.Now(),
					Due:   time.Now().Add(time.Hour),
				},
			},
			false,
		},
		{
			"ERR: Description",
			internal.Task{
				Priority: internal.PriorityHigh,
				Dates: internal.Dates{
					Start: time.Now(),
					Due:   time.Now().Add(time.Hour),
				},
			},
			true,
		},
		{
			"ERR: Priority",
			internal.Task{
				Description: "sample description",
				Priority:    internal.Priority(-1),
				Dates: internal.Dates{
					Start: time.Now(),
					Due:   time.Now().Add(time.Hour),
				},
			},
			true,
		},
		{
			"ERR: Dates(Start < Due)",
			internal.Task{
				Description: "sample description",
				Priority:    internal.PriorityHigh,
				Dates: internal.Dates{
					Start: time.Now().Add(time.Hour),
					Due:   time.Now(),
				},
			},
			true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actualErr := test.input.Validate()
			isErr := actualErr != nil
			require.Equal(t, test.withErr, isErr)

			var ierr *internal.Error
			if test.withErr {
				require.ErrorAs(t, actualErr, &ierr)
			}
		})
	}
}
