package dsputils

import (
	"testing"
)

func TestRootsToCoefficients(t *testing.T) {
	var coef = []float64{1, 5, 5, -5, -6}
	roots := []complex128{complex(-3, 0), complex(-2, 0), complex(-1, 0), complex(1, 0)}
	c := rootsToCoefficients(roots)
	for i, cc := range c {
		if cc != coef[i] {
			t.Error("%d: %f != %f", i, cc, coef[i])
		}
	}
}
