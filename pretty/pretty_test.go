package pretty

import (
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

	Print(Person{})
}
