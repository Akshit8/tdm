// Package internal defines the API entities and their corresponding attributes/methods.
package internal

import (
	"context"
	"errors"
	"fmt"
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

	return errors.New("unknow value")
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
		return errors.New("start date should be before end data")
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
		return errors.New("description is required")
	}

	if err := t.Priority.Validate(); err != nil {
		return fmt.Errorf("priority invalid: %w", err)
	}

	if err := t.Dates.Validate(); err != nil {
		return fmt.Errorf("dates invalid: %w", err)
	}

	return nil
}

// TaskRepository defines the datastore handling persisting Task records.
type TaskRepository interface {
	Create(ctx context.Context, description string, priority Priority, dates Dates) (Task, error)
	Find(ctx context.Context, id string) (Task, error)
	Update(ctx context.Context, id string, description string, priority Priority, dates Dates, isDone bool) error
}

// TaskService defines available operation on Task Service
type TaskService interface {
	Create(ctx context.Context, description string, priority Priority, dates Dates) (Task, error)
	Task(ctx context.Context, id string) (Task, error)
	Update(ctx context.Context, id string, description string, priority Priority, dates Dates, isDone bool) error
}
