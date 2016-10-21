package main

import (
	"fmt"
	"math/rand"

	"github.com/gonum/matrix/mat64"
)

type Network struct {
	sizes  []int
	layers int
	biases *mat64.Dense
	weights *mat64.Dense
}

func main() {
	data := make([]float64, 36)
	for i := range data {
		data[i] = rand.NormFloat64()
	}

	a := mat64.NewDense(6, 6, data)
	fmt.Println(a)
}
