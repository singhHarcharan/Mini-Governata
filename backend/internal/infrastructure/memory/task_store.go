// Package memory provides the Phase 0 in-memory TaskStore implementation.
package memory

import (
	"context"
	"errors"
	"fmt"

	"mini-governata/backend/internal/application/workflow"
	"mini-governata/backend/internal/domain/task"
	"mini-governata/backend/internal/platform/tracectx"
)

var ErrTaskNotFound = errors.New("task not found")

// This compile-time assertion documents that *Store implements workflow.TaskStore.
var _ workflow.TaskStore = (*Store)(nil)

// Store stands in for the real database adapter introduced in a later phase.
type Store struct {
	tasks map[string]task.Task
}

// NewStore creates a store and copies the supplied tasks into it.
func NewStore(initialTasks ...*task.Task) *Store {
	store := &Store{tasks: make(map[string]task.Task, len(initialTasks))}
	for _, value := range initialTasks {
		if value != nil {
			store.tasks[value.ID] = *value
		}
	}

	return store
}

// Find returns a copy so changes are not persisted until Save is called.
func (s *Store) Find(ctx context.Context, id string) (*task.Task, error) {
	tracectx.Log(ctx, "MemoryStore", fmt.Sprintf("looking up task %q", id))

	value, ok := s.tasks[id]
	if !ok {
		return nil, ErrTaskNotFound
	}

	return &value, nil
}

// Save copies the task into the map. It returns only an error, not a task.
func (s *Store) Save(ctx context.Context, value *task.Task) error {
	if value == nil {
		return errors.New("cannot save a nil task")
	}

	tracectx.Log(ctx, "MemoryStore", fmt.Sprintf("persisting task %q", value.ID))
	s.tasks[value.ID] = *value
	return nil
}
