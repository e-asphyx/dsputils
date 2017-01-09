package dsputils

import (
	"math"
)

type FilterCoef struct {
	A []float64
	B []float64
}

type IntFilterCoef struct {
	A []int64
	B []int64
}

func (f FilterCoef) Norm() FilterCoef {
	b := make([]float64, len(f.B))
	for i, c := range f.B {
		b[i] = c / f.B[0]
	}

	return FilterCoef{
		A: f.A,
		B: b,
	}
}

func round(x float64) float64 {
	if x >= 0 {
		return math.Trunc(x + 0.5)
	} else {
		return math.Trunc(x - 0.5)
	}
}

func quantize(val float64, bits uint) int64 {
	return int64(round(val * float64(int64(1)<<bits)))
}

func (f FilterCoef) Quantize(bits uint) IntFilterCoef {
	b := make([]int64, len(f.B))
	a := make([]int64, len(f.A))

	for i, val := range f.B {
		b[i] = quantize(val, bits)
	}

	for i, val := range f.A {
		a[i] = quantize(val, bits)
	}

	return IntFilterCoef{
		A: a,
		B: b,
	}
}
