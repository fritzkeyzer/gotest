package main

import (
	varis "github.com/Xamber/Varis"
	"github.com/dathoangnd/gonet"
)

var (
	varisNet varis.Perceptron
	goNet    gonet.NN

	gonetTrainingData = [][][]float64{
		{{0, 0}, {0}},
		{{0, 1}, {1}},
		{{1, 0}, {1}},
		{{1, 1}, {0}},
		{{0.5, 0.5}, {0.5}},
		{{0.25, 0.25}, {0.75}},
	}

	varisTrainingData = varis.Dataset{
		{varis.Vector{0.00, 0.00}, varis.Vector{0.00}},
		{varis.Vector{0.00, 1.00}, varis.Vector{1.00}},
		{varis.Vector{1.00, 0.00}, varis.Vector{1.00}},
		{varis.Vector{1.00, 1.00}, varis.Vector{0.00}},
		{varis.Vector{0.50, 0.50}, varis.Vector{0.50}},
		{varis.Vector{0.25, 0.25}, varis.Vector{0.75}},
	}

	gonetTestData = [][]float64{
		{0.1, 0.9},
		{0.2, 0.8},
		{0.3, 0.7},
		{0.4, 0.6},
		{0.5, 0.5},
		{0.6, 0.4},
		{0.7, 0.3},
		{0.8, 0.2},
		{0.9, 0.1},
	}

	varisTestData = []varis.Vector{
		{0.1, 0.9},
		{0.2, 0.8},
		{0.3, 0.7},
		{0.4, 0.6},
		{0.5, 0.5},
		{0.6, 0.4},
		{0.7, 0.3},
		{0.8, 0.2},
		{0.9, 0.1},
	}
)

func init() {

	varisNet = varis.CreatePerceptron(2, 10, 10, 1)
	goNet = gonet.New(2, []int{10, 10}, 1, false)
}

func TrainVaris(iterations int) {
	trainer := varis.PerceptronTrainer{
		Network: &varisNet,
		Dataset: varisTrainingData,
	}

	trainer.BackPropagation(iterations)

}

func TrainGoNet(iterations int) {
	goNet.Train(gonetTrainingData, iterations, 0.4, 0.2, true)
}

func RunGonet(iterarions int) {
	for _, input := range gonetTestData {
		goNet.Predict(input)
	}
}

func RunVaris(iterarions int) {
	for _, input := range varisTestData {
		varisNet.Calculate(input)
	}
}
