package batch

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

// Execute will run taskFn for each element in tasks.
// Number of parallel threads can be specified or left default by using 0. (default is runtime.NumCPU threads)
// Errors returned by taskFn are collected and returned as a TaskError - remaining tasks are processed.
// Panics will cancel remaining executions and return with an error of type TaskPanic.
// (Panics can be used to cancel remaining executions.)
func Execute[T any](tasks []T, taskFn func(v T) error, parallel int, showProgressBar bool) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("execute panic: %v: %v", r, string(debug.Stack()))
		}
	}()

	// create a channel of booleans, to limit parallel executions (similar to a thread pool)
	if parallel < 1 {
		parallel = runtime.NumCPU()
	}
	queue := make(chan bool, parallel)

	// create a sync / wait group, with a size equal to the size of the tasks
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))

	// create a thread safe cancellation flag
	var status int32 = 1

	// initialise error variables
	errLock := &sync.Mutex{}

	// progress bar:
	var progressBar *ProgressBar
	if showProgressBar {
		progressBar = NewProgressBar(len(tasks))
		defer progressBar.Cancel()
		progressBar.Start()
	}

	// start task executions
	for i, t := range tasks {
		// block here if all threads are busy
		queue <- true

		// start execution on new thread
		// pass values into closure - preventing thread misreads
		go func(task T, taskI int) {
			defer func() {
				if r := recover(); r != nil {
					errLock.Lock()
					defer errLock.Unlock()

					// cancel further executions
					atomic.StoreInt32(&status, -1)

					// only save panic if err is nil
					if err == nil {
						err = TaskPanic{
							Index:      taskI,
							Task:       task,
							Message:    fmt.Sprint(r),
							StackTrace: string(debug.Stack()),
						}
					}
				}

				// decrement counter and pop value out of channel (freeing up an execution "slot")
				wg.Done()
				<-queue
				if atomic.LoadInt32(&status) == 1 && progressBar != nil {
					progressBar.Increment()
				}
			}()

			// if not cancelled, run task
			if atomic.LoadInt32(&status) == 1 {
				// execute task
				if err2 := taskFn(task); err2 != nil {
					errLock.Lock()
					defer errLock.Unlock()

					// cancel further executions
					atomic.StoreInt32(&status, -1)

					// only save error if err is nil
					if err == nil {
						err = TaskError{
							Index:   taskI,
							Task:    task,
							Message: err2.Error(),
						}
					}
				}
			}
		}(t, i)
	}

	// block further execution until all tasks have completed
	wg.Wait()

	return err
}
