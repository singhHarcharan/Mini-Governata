// Package handler represents the application boundary that will become HTTP in Phase 3.
package handler

import (
	"context"
	"fmt"

	"mini-governata/backend/internal/application/workflow"
	"mini-governata/backend/internal/domain/task"
	"mini-governata/backend/internal/platform/tracectx"
)

// TaskHandler receives an action and hands it to the workflow engine.
type TaskHandler struct {
	engine *workflow.Engine
}

// NewTaskHandler injects the workflow engine into the handler.
func NewTaskHandler(engine *workflow.Engine) *TaskHandler {
	return &TaskHandler{engine: engine}
}

// Complete simulates handling POST /tasks/:id/complete.
func (h *TaskHandler) Complete(
	ctx context.Context, 
	id string,
) (*task.Task, error) {
	tracectx.Log(ctx, "Handler", fmt.Sprintf("received Complete for task %q", id))
	return h.engine.CompleteTask(ctx, id) // Direct call from handler to engine.
}
