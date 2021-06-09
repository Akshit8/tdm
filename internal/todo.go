// Package internal defines the API entities and their corresponding attributes/methods.
package internal

import (
	"time"
)

// Priority indicates how important a Task is.
type Priority int8

const (
	// PriorityNone indicates the task needs to be prioritized.
	PriorityNone Priority = iota

	// PriorityLow indicates a non urgent task.
	PriorityLow

	// PriorityMedium indicates a task that should be completed soon.
	PriorityMedium

	// PriorityHigh indicates an urgent task that must be completed as soon as possible.
	PriorityHigh
)

// Validate the Priority field on Task.
func (p Priority) Validate() error {
	switch p {
	case PriorityNone, PriorityLow, PriorityMedium, PriorityHigh:
		return nil
	}

	return NewErrorf(ErrorCodeInvalidArgument, "unknown value")
}

// Category is human readable value meant to be used to organize your tasks. Category values are unique.
type Category string

// Dates indicates a point in time where a task starts or completes, dates are not enforced on Tasks.
type Dates struct {
	Start time.Time
	Due   time.Time
}

// Validate the Dates field on Task.
func (d Dates) Validate() error {
	if !d.Start.IsZero() && !d.Due.IsZero() && d.Start.After(d.Due) {
		return NewErrorf(ErrorCodeInvalidArgument, "start date should be before due date")
	}

	return nil
}

// Task is an activity that needs to be completed within a period of time.
type Task struct {
	ID          string
	Description string
	Priority    Priority
	Dates       Dates
	SubTasks    []Task
	Categories  []Category
	IsDone      bool
}

// Validate the Task object.
func (t Task) Validate() error {
	if t.Description == "" {
		return NewErrorf(ErrorCodeInvalidArgument, "description is required")
	}

	if err := t.Priority.Validate(); err != nil {
		return WrapErrorf(err, ErrorCodeInvalidArgument, "priority is invalid")
	}

	if err := t.Dates.Validate(); err != nil {
		return WrapErrorf(err, ErrorCodeInvalidArgument, "dates are invalid")
	}

	return nil
}
