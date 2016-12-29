package dsputils

type FilterCoef struct {
	A []float64
	B []float64
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
