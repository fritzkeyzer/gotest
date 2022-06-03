package batch

import (
	"fmt"
)

// TaskError implements the error interface.
// It contains details about individual task errors during execution.
type TaskError struct {
	ErrMap map[int]string
}

func (e TaskError) Error() string {
	errString := fmt.Sprintf("Task errors: %d\n", len(e.ErrMap))

	for i, s := range e.ErrMap {
		errString += fmt.Sprintf("    task[%d]: error: %s\n", i, s)
	}

	return errString
}

// IsError returns true if the TaskError is not nil.
func (e TaskError) IsError() bool {
	if e.ErrMap == nil {
		return false
	}
	if len(e.ErrMap) > 0 {
		return true
	}
	return false
}

// TaskPanic implements the error interface.
// The Execute function will return a TaskPanic if any of the tasks cause a panic.
type TaskPanic struct {
	ErrIndex int
	ErrMsg   string
}

func (p TaskPanic) Error() string {
	return fmt.Sprintf("task[%d]: panic: %s", p.ErrIndex, p.ErrMsg)
}

// IsError returns true if the TaskPanic is not nil.
func (p TaskPanic) IsError() bool {
	if p.ErrIndex == 0 && p.ErrMsg == "" {
		return false
	}
	return true
}
