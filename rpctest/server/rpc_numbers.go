package main

import "errors"

type Arith struct{
	// we dont need anything here
	//id string
	//val float64
}

type Args struct{
	A, B float64
}

type Quotient struct {
	Quo, Rem int
}

func (a *Arith) Multiply(args Args, reply *float64) error {
	*reply = args.A * args.B
	return nil
}

func (a *Arith) Divide(args Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = int(args.A) / int(args.B)
	quo.Rem = int(args.A) % int(args.B)
	return nil
}
