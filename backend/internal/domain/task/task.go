// Package task contains the smallest piece of business logic in the lab.
package task

import "errors"

// Status is a named string type, so task statuses are not arbitrary strings.
type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
)

var ErrInvalidTransition = errors.New("invalid task transition")

// Task is the domain object used by every layer.
type Task struct {
	ID     string `json:"id"`
	Status Status `json:"status"`
}

// Complete protects the rule that only an in-progress task can be completed.
func (t *Task) Complete() error {
	if t.Status != StatusInProgress {
		return ErrInvalidTransition
	}

	t.Status = StatusCompleted
	return nil
}
