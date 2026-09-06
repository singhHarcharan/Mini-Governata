package workflow

import (
	"context"
	"testing"

	"mini-governata/backend/internal/domain/task"
)

// fakeStore satisfies TaskStore without using the memory adapter or a database.
type fakeStore struct {
	value *task.Task
	saved *task.Task
}

func (f *fakeStore) Find(context.Context, string) (*task.Task, error) {
	copyOfTask := *f.value
	return &copyOfTask, nil
}

func (f *fakeStore) Save(_ context.Context, value *task.Task) error {
	copyOfTask := *value
	f.saved = &copyOfTask
	return nil
}

func TestEngineCompleteTask(t *testing.T) {
	store := &fakeStore{value: &task.Task{
		ID:     "task-001",
		Status: task.StatusInProgress,
	}}
	engine := NewEngine(store)

	completed, err := engine.CompleteTask(context.Background(), "task-001")
	if err != nil {
		t.Fatalf("CompleteTask() error = %v", err)
	}
	if completed.Status != task.StatusCompleted {
		t.Fatalf("returned status = %q, want %q", completed.Status, task.StatusCompleted)
	}
	if store.saved == nil || store.saved.Status != task.StatusCompleted {
		t.Fatal("completed task was not passed to TaskStore.Save")
	}
}
