package rest

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Akshit8/tdm/internal"
)

// Priority indicates how important a Task is.
type Priority string

const (
	priorityNone   Priority = "none"
	priorityLow    Priority = "low"
	priorityMedium Priority = "medium"
	priorityHigh   Priority = "high"
)

// NewPriority converts the received domain type to a rest type,
// when the argument is unknown "none" is used.
func NewPriority(p internal.Priority) Priority {
	switch p {
	case internal.PriorityNone:
		return priorityNone
	case internal.PriorityLow:
		return priorityLow
	case internal.PriorityMedium:
		return priorityMedium
	case internal.PriorityHigh:
		return priorityHigh
	}

	return priorityNone
}

// Convert returns the domain type defining the internal representation,
// when priority is unknown "none" is used.
func (p Priority) Convert() internal.Priority {
	switch p {
	case priorityNone:
		return internal.PriorityNone
	case priorityLow:
		return internal.PriorityLow
	case priorityMedium:
		return internal.PriorityMedium
	case priorityHigh:
		return internal.PriorityHigh
	}

	return internal.PriorityNone
}

// Validate the priority
func (p Priority) Validate() error {
	switch p {
	case "none", "low", "medium", "high":
		return nil
	}

	return errors.New("unknown value")
}

// MarshalJSON ...
// This function MarshalJSOn on type Priority will be invoked by json.Marshal due to internal working of json pkg
// search about it more!
func (p Priority) MarshalJSON() ([]byte, error) {
	err := p.Validate()
	if err != nil {
		return nil, fmt.Errorf("conver: %w", err)
	}

	b, err := json.Marshal(string(p))
	if err != nil {
		return nil, fmt.Errorf("json marshal: %w", err)
	}

	return b, nil
}

// UnmarshalJSON ...
func (p *Priority) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("json unmarshal: %w", err)
	}

	err = Priority(s).Validate()
	if err != nil {
		return fmt.Errorf("convert: %w", err)
	}

	*p = Priority(s)

	return nil
}
