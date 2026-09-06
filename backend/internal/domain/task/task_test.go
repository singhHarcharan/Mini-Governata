package task

import (
	"errors"
	"testing"
)

func TestComplete(t *testing.T) {
	tests := []struct {
		name       string
		initial    Status
		wantStatus Status
		wantErr    error
	}{
		{
			name:       "in-progress task can be completed",
			initial:    StatusInProgress,
			wantStatus: StatusCompleted,
		},
		{
			name:       "pending task cannot be completed",
			initial:    StatusPending,
			wantStatus: StatusPending,
			wantErr:    ErrInvalidTransition,
		},
		{
			name:       "completed task cannot be completed again",
			initial:    StatusCompleted,
			wantStatus: StatusCompleted,
			wantErr:    ErrInvalidTransition,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := &Task{ID: "task-001", Status: test.initial}

			err := value.Complete()

			if !errors.Is(err, test.wantErr) {
				t.Fatalf("Complete() error = %v, want %v", err, test.wantErr)
			}
			if value.Status != test.wantStatus {
				t.Fatalf("task status = %q, want %q", value.Status, test.wantStatus)
			}
		})
	}
}
