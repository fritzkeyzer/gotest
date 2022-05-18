package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("ERR rpc.DialHTTP: ", err)
	}

	// RPC ARITH:
	args := Args{17, 8}
	var reply float64
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("ERR Arith.Multiply: ", err)
	}
	log.Printf("Arith.Multiply: %v * %v = %v\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("ERR Arith.Divide: ", err)
	}
	log.Printf("Arith.Divide: %v / %v = %d rem %d\n", args.A, args.B, quot.Quo, quot.Rem)


	// RPC DB:
	dbArgs := "whatever"
	var connString string
	err = client.Call("DB.GetConnString", dbArgs, &connString)
	if err != nil {
		log.Fatal("ERR DB.GetConnString: ", err)
	}
	log.Printf("DB.GetConnString: '%s'\n", connString)

	var recs []Record
	err = client.Call("DB.GetRecords", dbArgs, &recs)
	if err != nil {
		log.Fatal("ERR DB.GetRecords: ", err)
	}
	log.Printf("DB.GetRecords: %+v\n", recs)

	//var recs []Record
	err = client.Call("DB.GetRecordsError", dbArgs, &recs)
	if err != nil {
		log.Fatal("ERR DB.GetRecordsError: ", err)
	}
	log.Printf("DB.GetRecordsError: %v\n", recs)
}