package pretty

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {

	type Person struct {
		Married bool
		married bool
		Age     uint
		age     uint
		Id      int
		id      int
		Worth   float64
		worth   float32
		Name    string
		name    string

		Spouse *Person
		spouse *Person

		Children []Person
		children []Person

		Friends map[string]Person
		friends map[string]Person
	}

	s := String(Person{}, -1)
	fmt.Println(s)

	Print(Person{
		Spouse:   &Person{},
		spouse:   &Person{},
		Children: []Person{{}, {}},
		children: []Person{{}, {}},
		Friends:  nil,
		friends:  nil,
	})

	s = String(Person{
		Spouse:   &Person{},
		spouse:   &Person{},
		Children: []Person{{}, {}},
		children: []Person{{}, {}},
		Friends:  nil,
		friends:  nil,
	}, -1)
	fmt.Println(s)
}
