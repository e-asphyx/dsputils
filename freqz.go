package dsputils

import (
	"math"
	"math/cmplx"
)

type PoleZero struct {
	P []complex128
	Z []complex128
}

type FilterCoef struct {
	A []float64
	B []float64
}

func (p PoleZero) FreqZ(points int, fs ...int) [][2]float64 {
	var fFs float64 = 1
	if len(fs) != 0 {
		fFs = float64(fs[0])
	}

	data := make([][2]float64, points)

	for i := 0; i < points; i++ {
		relPos := float64(i) / float64(points-1)
		dst := cmplx.Rect(1, relPos*math.Pi)

		var pProd float64 = 1
		for _, v := range p.P {
			pProd *= cmplx.Abs(dst - v)
		}

		var zProd float64 = 1
		for _, v := range p.Z {
			zProd *= cmplx.Abs(dst - v)
		}

		var mag float64
		if pProd != 0 {
			mag = zProd / pProd
		} else {
			mag = math.MaxFloat64
		}

		data[i] = [2]float64{relPos * 0.5 * fFs, mag}
	}

	return data
}

func (p PoleZero) Coef() FilterCoef {
	return FilterCoef{
		A: rootsToCoefficients(p.P),
		B: rootsToCoefficients(p.Z),
	}
}

func rootsToCoefficients(roots []complex128) []float64 {
	p := make([]complex128, len(roots)+1)
	p1 := make([]complex128, len(p))

	p[0] = complex(1, 0)

	for _, r := range roots {
		for i, _ := range p {
			p1[i] = p[i] * r
		}

		copy(p[1:], p)
		p[0] = 0

		for i, _ := range p {
			p[i] -= p1[i]
		}
	}

	result := make([]float64, len(p))
	for i, c := range p {
		result[i] = real(c)
	}
	return result
}
