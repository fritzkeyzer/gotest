package main

import (
	"log"
	"net/http"
	"net/rpc"
)

func main() {

	rpcArith := new(Arith)
	err := rpc.Register(rpcArith)
	if err != nil {
		log.Fatalln(err)
	}

	rpcDB := new(DB)
	rpcDB.connString = "123.456.789"
	rpcDB.genTestData()
	err = rpc.Register(rpcDB)
	if err != nil {
		log.Fatalln(err)
	}

	rpc.HandleHTTP()

	err = http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatalln(err)
	}
}



