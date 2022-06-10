package batch

import (
	"fmt"
	"testing"
	"time"
)

func fn(d int) error {
	time.Sleep(10 * time.Microsecond)

	if d < 0 {
		return fmt.Errorf("we dislike negative numbers")
	}
	if d == 13 {
		panic("we really dislike 13")
	}

	return nil
}

func TestParallel(t *testing.T) {
	var data []int

	for i := 0; i < 100000; i++ {
		val := i
		if val == 13 {
			val = 12 // 13 panics!
		}
		data = append(data, val)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("parallel : %d\n", i)
		err := Execute(data, fn, i, true)
		if err != nil {
			t.Error(err)
		}
	}

	//fmt.Println("parallel : 0")
	//err = Execute(data, fn, 0, true)
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Println("parallel : 1")
	//err = Execute(data, fn, 1, true)
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Println("parallel : 4")
	//err = Execute(data, fn, 4, true)
	//if err != nil {
	//	t.Error(err)
	//}
}

func TestError(t *testing.T) {
	for i := -1; i < 10; i++ {
		fmt.Printf("parallel : %d - with errors: \n", i)
		data := []int{0, 1, 2, 3, 4, -5, 6, 7, 8, 9}
		err := Execute(data, fn, i, true)
		if err == nil {
			t.Error("expected an error")
			return
		}

		te, isTaskError := err.(TaskError)
		if !isTaskError {
			t.Error("expected err to be a *TaskError:", err)
			return
		}
		if te.Index != 5 {
			t.Error("expected Index to be 5")
		}
		if te.Task == nil {
			t.Error("expected Task != nil")
		}
		if te.Message == "" {
			t.Error("expected Message != \"\"")
		}

		if t.Failed() {
			t.Log(err)
		}
	}
}

func TestPanic(t *testing.T) {

	for i := -1; i < 10; i++ {
		fmt.Printf("parallel : %d - with panic\n", i)
		data := []int{0, 1, 2, 3, 13, 5, 6, 7, 8, 9}
		err := Execute(data, fn, i, true)
		if err == nil {
			t.Error("expected an error")
			return
		}

		te, isTaskPanic := err.(TaskPanic)
		if !isTaskPanic {
			t.Error("expected err to be a *TaskPanic:", err)
			return
		}
		if te.Index != 4 {
			t.Error("expected Index to be 4")
		}
		if te.Task == nil {
			t.Error("expected Task != nil")
		}
		if te.Message == "" {
			t.Error("expected Message != \"\"")
		}
		if te.StackTrace == "" {
			t.Error("expected StackTrace != \"\"")
		}

		if t.Failed() {
			t.Log(err)
		}
	}

}
