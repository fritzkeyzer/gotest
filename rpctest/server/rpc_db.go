package main

import "fmt"

type DB struct{
	connString string

	// faking a database by storing slice of Records here
	Records []Record
}

type Record struct{
	ID string
	ValA float64
	ValB float64
}

func (db *DB) genTestData(){
	someRandomTestRecords := []Record{
		Record{
			ID:   "001",
			ValA: 0,
			ValB: 0,
		},
		Record{
			ID:   "002",
			ValA: 1.23,
			ValB: 45.6,
		},
	}

	db.Records = someRandomTestRecords
}


func (db *DB) GetRecords(args string, records *[]Record) error{
	*records = db.Records
	return nil
}

func (db *DB) GetRecordsError(args string, records *[]Record) error{
	*records = db.Records
	return fmt.Errorf("couldnt get records because this is a test")
}

func (db *DB) GetConnString(args string, connString *string) error{
	*connString = db.connString
	return nil
}