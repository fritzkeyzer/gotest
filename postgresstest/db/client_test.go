package db

import (
	"flag"
	"github.com/fritzkeyzer/test/postgresstest/types"
	"math/rand"
	"testing"
)

var (
	connString = flag.String("connstring", "postgres://postgres:password@localhost:5432/postgres", "DB connection string")
	debug      = flag.Bool("debug", false, "debug mode")
)

func TestDBClient(t *testing.T) {

	cl := NewDBClient(DBConfig{
		ConnString: *connString,
		Debug:      *debug,
	})

	err := cl.Ping()
	if err != nil {
		t.Fatal(err)
	}

	err = cl.DropTableStudents()
	if err != nil {
		t.Error(err)
	}

	err = cl.DropTableStudents()
	if err != nil {
		t.Error(err)
	}

	err = cl.CreateTableStudents()
	if err != nil {
		t.Error(err)
	}

	err = cl.DeleteAllStudents()
	if err != nil {
		t.Error(err)
	}

	err = cl.InsertStudent(&types.Student{
		Name: "Bob",
		Age:  int(100 * rand.Float64()),
	})
	if err != nil {
		t.Error(err)
	}

	var students []types.Student
	err = cl.GetAllStudents(&students)
	if err != nil {
		t.Error(err)
	}
}
