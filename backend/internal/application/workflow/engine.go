// Package workflow coordinates task business rules and persistence.
package workflow

import (
	"context"
	"fmt"

	"mini-governata/backend/internal/domain/task"
	"mini-governata/backend/internal/platform/tracectx"
)

// TaskStore describes only the storage behavior the Engine needs.
// A concrete store satisfies this interface implicitly by having these methods.
type TaskStore interface {
	Find(ctx context.Context, id string) (*task.Task, error)
	Save(ctx context.Context, value *task.Task) error
}

// Engine contains business workflow behavior.
type Engine struct {
	store TaskStore
}

// NewEngine injects a TaskStore implementation into a new Engine.
func NewEngine(store TaskStore) *Engine {
	return &Engine{store: store}
}

// CompleteTask finds a task, applies the domain rule, and saves the result.
func (e *Engine) CompleteTask(ctx context.Context, id string) (*task.Task, error) {
	tracectx.Log(ctx, "Engine", fmt.Sprintf("finding task %q", id))

	value, err := e.store.Find(ctx, id) // Interface dispatch happens here.
	if err != nil {
		return nil, fmt.Errorf("find task: %w", err)
	}

	tracectx.Log(ctx, "Engine", "checking the completion rule")
	if err := value.Complete(); err != nil { // Direct method call.
		return nil, fmt.Errorf("complete task %q: %w", id, err)
	}

	tracectx.Log(ctx, "Engine", "saving completed task")
	if err := e.store.Save(ctx, value); err != nil { // Interface dispatch again.
		return nil, fmt.Errorf("save task %q: %w", id, err)
	}

	return value, nil
}
