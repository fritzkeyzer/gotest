package main

import (
	"fmt"
	"github.com/Xamber/Varis"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//net := varis.CreatePerceptron(2, 10, 1)
	//
	//dataset := varis.Dataset{
	//	{varis.Vector{0.0, 0.0}, varis.Vector{1.0}},
	//	{varis.Vector{1.0, 0.0}, varis.Vector{0.0}},
	//	{varis.Vector{0.0, 1.0}, varis.Vector{0.0}},
	//	{varis.Vector{1.0, 1.0}, varis.Vector{1.0}},
	//	{varis.Vector{0.5, 0.5}, varis.Vector{0.5}},
	//}
	//
	//trainer := varis.PerceptronTrainer{
	//	Network: &net,
	//	Dataset: dataset,
	//}
	//
	//log.Println("training...")
	//trainer.BackPropagation(10000)
	//log.Println("training complete")
	//
	//varis.PrintCalculation = true
	//net.Calculate(varis.Vector{0.0, 0.0}) // Output: [0.9816677167418877]
	//net.Calculate(varis.Vector{1.0, 0.0}) // Output: [0.02076530509106318]
	//net.Calculate(varis.Vector{0.0, 1.0}) // Output: [0.018253250887023762]
	//net.Calculate(varis.Vector{1.0, 1.0}) // Output: [0.9847884089930481]\
	//net.Calculate(varis.Vector{0.5, 0.5}) // Output: [0.9847884089930481]\
	//
	//err := SaveNet(net, "varisNet.json")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	net2, err := LoadNet("varisNet.json")
	if err != nil {
		log.Fatalln(err)
	}

	varis.PrintCalculation = true
	net2.Calculate(varis.Vector{0.0, 0.0}) // Output: [0.9816677167418877]
	net2.Calculate(varis.Vector{1.0, 0.0}) // Output: [0.02076530509106318]
	net2.Calculate(varis.Vector{0.0, 1.0}) // Output: [0.018253250887023762]
	net2.Calculate(varis.Vector{1.0, 1.0}) // Output: [0.9847884089930481]\
	net2.Calculate(varis.Vector{0.5, 0.5}) // Output: [0.9847884089930481]\
}

func SaveNet(net varis.Perceptron, filename string) error {
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		return fmt.Errorf("create directory failed: %v", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file failed: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(varis.ToJSON(net))
	if err != nil {
		return fmt.Errorf("save file failed: %v", err)
	}
	return nil
}

func LoadNet(filename string) (varis.Perceptron, error) {

	jsonString, err := ioutil.ReadFile(filename)
	if err != nil {
		return varis.Perceptron{}, fmt.Errorf("read file failed: %v", err)
	}

	return varis.FromJSON(string(jsonString)), nil
}
