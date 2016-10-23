package sujana

import (
	"math"
	"math/rand"

	"github.com/gonum/matrix/mat64"
)

func randMat(x, y int) *mat64.Dense {
	data := make([]float64, x*y)
	for i := range data {
		data[i] = rand.NormFloat64()
	}
	return mat64.NewDense(x, y, data)
}

type Network struct {
	Sizes   []int
	Layers  int
	Biases  []*mat64.Dense
	Weights []*mat64.Dense
}

func (n *Network) GetLayers() {
	n.Layers = len(n.Sizes)
}

func (n *Network) GetBiases() {
	var ret []*mat64.Dense
	for _, vals := range n.Sizes[1:] {
		ret = append(ret, randMat(vals, 1))
	}
	n.Biases = ret
}

func (n *Network) GetWeights() {
	var ret []*mat64.Dense
	for i := range n.Sizes[:n.Layers-1] {
		ret = append(ret, randMat(n.Sizes[i], n.Sizes[i+1]))
	}
	n.Weights = ret
}

func sigmoid(z float64) float64 {
	return 1.0 / (1.0 + math.Exp(-z))
}
