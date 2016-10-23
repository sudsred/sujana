package main

import (
	"fmt"
	"sujana"

	"github.com/gonum/matrix/mat64"
)

func main() {
	size := []int{3, 4, 5}
	nets := sujana.Network{Sizes: size}
	nets.GetLayers()
	nets.GetBiases()
	nets.GetWeights()
	zip := sujana.ZipMat{A: nets.Biases, B: nets.Weights}
	zip.Zip()
	for _, vals := range zip.Result {
		fmt.Println(mat64.Dot(vals.A.ColView(0), vals.B.RowView(0)))

	}
}
