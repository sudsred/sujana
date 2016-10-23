package sujana

import "github.com/gonum/matrix/mat64"

type zipper interface {
	Zip()
}

type zipMat64 struct {
	A, B *mat64.Dense
}

type ZipMat struct {
	A, B   []*mat64.Dense
	Result []zipMat64
}

func (z *ZipMat) Zip() {
	if len(z.A) != len(z.B) {
		panic("Zip lengths do not match")
	}

	r := make([]zipMat64, len(z.A))
	for i, e := range z.A {
		r[i] = zipMat64{e, z.B[i]}
	}

	z.Result = r

}

type zipInt struct {
	A, B int
}

type ZipInt struct {
	A, B   []int
	Result []zipInt
}

func (z *ZipInt) Zip() {
	if len(z.A) != len(z.B) {
		panic("Zip lengths do not match")
	}

	r := make([]zipInt, len(z.A))
	for i, e := range z.A {
		r[i] = zipInt{e, z.B[i]}
	}

	z.Result = r

}
