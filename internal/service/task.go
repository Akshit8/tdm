// Package service implements logical operation available on API entities.
package service

import (
	"context"
	"fmt"

	"github.com/Akshit8/tdm/internal"
)

//go:generate counterfeiter -o servicetesting/task_service.gen.go . TaskService

// TaskRepository defines the datastore handling persisting Task records.
type TaskRepository interface {
	Create(ctx context.Context, description string, priority internal.Priority, dates internal.Dates) (internal.Task, error)
	Find(ctx context.Context, id string) (internal.Task, error)
	Update(ctx context.Context, id string, description string, priority internal.Priority, dates internal.Dates, isDone bool) error
	Delete(ctx context.Context, id string) error
}

// TaskService defines available operation on Task Service
type TaskService interface {
	Create(ctx context.Context, description string, priority internal.Priority, dates internal.Dates) (internal.Task, error)
	Task(ctx context.Context, id string) (internal.Task, error)
	Update(ctx context.Context, id string, description string, priority internal.Priority, dates internal.Dates, isDone bool) error
	Delete(ctx context.Context, id string) error
}

type task struct {
	repo TaskRepository
}

// NewTask creates new instance of Task.
func NewTask(repo TaskRepository) TaskService {
	return &task{
		repo: repo,
	}
}

// Create stores a new record.
func (t *task) Create(ctx context.Context, description string, priority internal.Priority, dates internal.Dates) (internal.Task, error) {
	task, err := t.repo.Create(ctx, description, priority, dates)
	if err != nil {
		return internal.Task{}, fmt.Errorf("repo create: %w", err)
	}

	return task, nil
}

// Task gets an existing Task from the datastore.
func (t *task) Task(ctx context.Context, id string) (internal.Task, error) {
	task, err := t.repo.Find(ctx, id)
	if err != nil {
		return internal.Task{}, fmt.Errorf("repo find: %w", err)
	}

	return task, nil
}

// Update updates an existing Task in the datastore.
func (t *task) Update(ctx context.Context, id string, description string, priority internal.Priority, dates internal.Dates, isDone bool) error {
	err := t.repo.Update(ctx, id, description, priority, dates, isDone)
	if err != nil {
		return fmt.Errorf("repo update: %w", err)
	}

	return nil
}

// Delete removes an existing Task from the datastore.
func (t *task) Delete(ctx context.Context, id string) error {
	err := t.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("repo delete: %w", err)
	}

	return nil
}
