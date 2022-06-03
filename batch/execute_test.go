package batch

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {

	rand.Seed(time.Now().Unix())

	type person struct {
		name   string
		age    int
		friend *person
	}

	data := []*person{
		{
			name: "bob 😃",
			age:  0,
		}, {
			name: "alice",
			age:  0,
		}, {
			name: "charlie",
			age:  0,
		}, {
			name: "jack",
			age:  0,
		}, {
			name: "sparrow",
			age:  0,
			friend: &person{
				name: "frankie",
				age:  0,
			},
		}, {
			name: "jack",
			age:  0,
		}, {
			name: "jack",
			age:  0,
		}, {
			name: "jack",
			age:  0,
		}, {
			name: "jack",
			age:  0,
		}, {
			name: "jack",
			age:  0,
		},
	}

	fn := func(p *person) error {
		p.age = len(p.name)

		if rand.Float64() < 0.02 {
			panic("unlucky 2 percent")
		}

		time.Sleep(200 * time.Millisecond)

		if rand.Float64() < 0.3 {
			time.Sleep(500 * time.Millisecond)
			return fmt.Errorf("unlucky 30 percent %f", rand.Float64())
		}

		time.Sleep(300 * time.Millisecond)

		return nil
	}

	err := Execute(data, fn, 1, true)
	if err != nil {
		t.Failed()

		e, isTaskError := err.(TaskError)
		if isTaskError {
			fmt.Println("we had task errors:", e)

			return
		}

		e2, isTaskPanic := err.(TaskPanic)
		if isTaskPanic {
			fmt.Println("we had a task panic:", e2)

			return
		}

		fmt.Println("oh no:", err)

		return
	}

	//pretty.Print(data)

	log.Println(data)
}
