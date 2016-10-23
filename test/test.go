package main

import (
	"fmt"
	"sujana"
)

func main() {
	size := []int{3, 4, 5}
	nets := sujana.Network{Sizes: size}
	nets.GetLayers()
	nets.GetBiases()
	nets.GetWeights()
	zip := sujana.ZipMat{A: nets.Biases, B: nets.Weights}
	zip.Zip()
	fmt.Println(zip.Result)
	fmt.Println(nets.Biases, nets.Weights)
	zipi := sujana.ZipInt{A: []int{1, 2, 3, 4}, B: []int{5, 6, 7, 8}}
	zipi.Zip()
	fmt.Println(zipi.Result)
}
