package rest

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Akshit8/tdm/internal"
)

// Time represents an instant in time, JSON are strings using RFC3339.
type Time time.Time

// Dates indicates a point in time where a task starts or completes, dates are not enforced on Tasks.
type Dates struct {
	Start Time `json:"start"`
	Due   Time `json:"due"`
}

// NewDates convert internal.Dates to rest.Dates
func NewDates(d internal.Dates) Dates {
	return Dates{
		Start: Time(d.Start),
		Due:   Time(d.Due),
	}
}

// MarshalJSON ...
func (t Time) MarshalJSON() ([]byte, error) {
	str := time.Time(t).Format(time.RFC3339)

	b, err := json.Marshal(str)
	if err != nil {
		return nil, fmt.Errorf("json marshal: %w", err)
	}

	return b, nil
}

// UnmarshalJSON ...
func (t *Time) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("json unmarshal: %w", err)
	}

	tt, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return fmt.Errorf("convert: %w", err)
	}

	*t = Time(tt)

	return nil
}
