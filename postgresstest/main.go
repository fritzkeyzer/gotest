package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"github.com/natefinch/lumberjack"

	"github.com/fritzkeyzer/test/postgresstest/db"
	"github.com/fritzkeyzer/test/postgresstest/types"
)

var (
	connString  = flag.String("connstring", "postgres://postgres:password@localhost:5432/postgres", "DB connection string")
	outputFile  = flag.String("logfile", "", "log file output")
	debug = flag.Bool("debug", false, "debug mode")
)

func main() {
	flag.Parse()

	if len(*outputFile) > 0{
		log.SetOutput(&lumberjack.Logger{
			Filename:   *outputFile,
			MaxSize:    500, // megabytes
			MaxAge:     28, //days
			//Compress:   true,
		})
		fmt.Println("logging to file")
	}

	cl := db.NewDBClient(db.DBConfig{
		ConnString: *connString,
		Debug:      *debug,
	})

	err := cl.Ping()
	if err != nil {
		log.Fatal(err)
	}

	err = cl.DeleteAllStudents()
	if err != nil {
		log.Fatal(err)
	}

	err = cl.InsertStudent(&types.Student{
		Name: "Bob",
		Age:  int(100*rand.Float64()),
	})
	if err != nil {
		log.Fatal(err)
	}

	var students []types.Student
	err = cl.GetAllStudents(&students)
	if err != nil {
		log.Fatal(err)
	}

}
