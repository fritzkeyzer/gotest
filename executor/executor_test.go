package executor

import (
	"CodeSnippits/pretty"
	"log"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {

	type person struct {
		name   string
		age    int
		friend *person
	}

	data := []*person{
		{
			name: "bob 😃",
			age:  0,
		},
		{
			name: "alice",
			age:  0,
		},
		{
			name: "charlie",
			age:  0,
		},
		{
			name: "jack",
			age:  0,
		},
		{
			name: "sparrow",
			age:  0,
			friend: &person{
				name: "frankie",
				age:  0,
			},
		},
	}

	fn := func(v any) {
		entity := v.(*person)
		entity.age = len(entity.name)

		time.Sleep(500 * time.Millisecond)
	}

	err := Execute(data, fn, 1)
	if err != nil {
		log.Fatalln("oh no:", err)
	}

	pretty.Print(data)
}
