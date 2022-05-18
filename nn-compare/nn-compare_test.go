package main

import "testing"

func BenchmarkTrainGoNet(b *testing.B) {

	for n := 0; n < b.N; n++ {
		TrainGoNet(100)
	}
}

func BenchmarkTrainVaris(b *testing.B) {

	for n := 0; n < b.N; n++ {
		TrainVaris(100)
	}
}

func BenchmarkRunGonet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RunGonet(n)
	}
}

func BenchmarkRunVaris(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RunVaris(n)
	}
}
