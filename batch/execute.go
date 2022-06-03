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
	var taskError TaskError
	var taskPanic TaskPanic

	// progress bar:
	var progressBar *ProgressBar
	if showProgressBar {
		progressBar = NewProgressBar(len(tasks))
		defer progressBar.Cancel()
		progressBar.Start()
	}

	// start task executions
	for i, t := range tasks {
		// copy data
		task := t
		taskI := i

		// block here if all threads are busy
		queue <- true

		// start execution on new thread
		go func() {
			defer func() {
				if r := recover(); r != nil {
					taskPanic = TaskPanic{
						ErrIndex: taskI,
						ErrMsg:   fmt.Sprintf("%v: %v", r, string(debug.Stack())),
					}

					// cancel further executions - better than breaking early, because of wg.Wait()
					atomic.StoreInt32(&status, -1)

					if progressBar != nil {
						progressBar.Cancel()
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
				e := taskFn(task)

				// collect errors
				if e != nil {
					if taskError.ErrMap == nil {
						taskError.ErrMap = make(map[int]string)
					}
					taskError.ErrMap[taskI] = e.Error()
				}
			}
		}()
	}

	// block further execution until all tasks have completed
	wg.Wait()

	// check errors and return
	if taskPanic.IsError() {
		return taskPanic
	}
	if taskError.IsError() {
		return taskError
	}
	return nil
}
