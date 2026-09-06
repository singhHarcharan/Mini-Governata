// Command trace-lab runs one complete task flow through every Phase 0 layer.
package main

import (
	"context"
	"fmt"
	"log"

	"mini-governata/backend/internal/application/workflow"
	"mini-governata/backend/internal/domain/task"
	"mini-governata/backend/internal/infrastructure/memory"
	"mini-governata/backend/internal/platform/tracectx"
	"mini-governata/backend/internal/transport/handler"
)

func main() {
	// Composition root: concrete dependencies are created and connected here.
	store := memory.NewStore(&task.Task{
		ID:     "task-001",
		Status: task.StatusInProgress,
	})
	engine := workflow.NewEngine(store)
	taskHandler := handler.NewTaskHandler(engine)

	// Simulate one incoming request carrying a Trace ID.
	ctx := tracectx.WithID(context.Background(), "trace-phase-0")

	// Actual Implementation / Code Flow Starts Here.
	completedTask, err := taskHandler.Complete(ctx, "task-001")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nResult: task %q is now %q\n", completedTask.ID, completedTask.Status)
}
