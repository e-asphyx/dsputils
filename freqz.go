package dsputils

import (
	"math"
	"math/cmplx"
)

type PoleZero struct {
	P []complex128
	Z []complex128
}

func (p PoleZero) FreqZ(points int) [][2]float64 {
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

		data[i] = [2]float64{relPos * 0.5, zProd / pProd}
	}

	return data
}
