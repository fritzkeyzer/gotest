package main

// RPC Arith:

type Args struct {
	A, B float64
}

type Quotient struct {
	Quo, Rem int
}


// RPC DB:

type Record struct{
	ID string
	ValA float64
	ValB float64
}