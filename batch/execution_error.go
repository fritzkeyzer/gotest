package batch

import (
	"fmt"
)

// TaskError implements the error interface.
// The Execute function will return a TaskError if any of the tasks return a non-nil error.
type TaskError struct {
	Index   int
	Task    any
	Message string
}

func (e TaskError) Error() string {
	return fmt.Sprintf("task[%d] error: %s", e.Index, e.Message)
}

// TaskPanic implements the error interface.
// The Execute function will return a TaskPanic if any of the tasks panic.
type TaskPanic struct {
	Index      int
	Task       any
	Message    string
	StackTrace string
}

func (p TaskPanic) Error() string {
	return fmt.Sprintf("task[%d] panic: %s: %s", p.Index, p.Message, p.StackTrace)
}
