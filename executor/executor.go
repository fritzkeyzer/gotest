package executor

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

// Execute will run the function defined by fn for each element in the provided data.
// The data argument will be converted to type []any - if this fails
// Number of parallel executions can be specified or left automatic by using 0
// If the slice cannot be converted, the function will panic.
func Execute(data any, fn func(v any), parallel int) (err error) {
	defer func() {
		if err2 := recover(); err2 != nil {
			err = fmt.Errorf("execution panic: %v", err2)
		}
	}()

	// convert data argument of type any, to type of []any
	sliceOfData, err := castToSlice(data)
	if err != nil {
		return fmt.Errorf("cannot convert data to type []any")
	}

	// progress bar:
	progressBar := NewBar(len(sliceOfData))
	defer progressBar.Cancel()

	// create a channel of bools, to limit parallel executions
	if parallel < 1 {
		parallel = runtime.NumCPU()
	}
	queue := make(chan bool, parallel)

	// create a sync / wait group, with a size equal to the size of the data
	wg := sync.WaitGroup{}
	wg.Add(len(sliceOfData))

	// start the executions
	progressBar.Start()
	for _, v := range sliceOfData {
		if err != nil {
			return err
		}
		queue <- true
		x := v
		go func() {
			defer func() {
				if err2 := recover(); err2 != nil {
					err = fmt.Errorf("execution panic: %v", err2)
				} else {
					progressBar.Increment()
				}

				wg.Done()
				<-queue
			}()
			fn(x)
		}()

	}
	wg.Wait()

	return err
}

// castToSlice casts any -> []any.
// Returns error if data is not of type slice
func castToSlice(data any) ([]any, error) {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		return nil, fmt.Errorf("not a data")
	}

	sliceLen := val.Len()

	out := make([]interface{}, sliceLen)

	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}

	return out, nil
}
