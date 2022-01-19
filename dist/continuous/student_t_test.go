package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dt(t, ν) from package stats
func TestStudentTProbability(t *testing.T) {
	tol := 0.0000001

	cases := []struct {
		t, ν, expected float64
	}{
		{-4, 1, 0.01872411},
		{-3, 1, 0.03183099},
		{-2, 1, 0.06366198},
		{-1, 1, 0.1591549},
		{0, 1, 0.3183099},
		{1, 1, 0.1591549},
		{2, 1, 0.06366198},
		{3, 1, 0.03183099},
		{4, 1, 0.01872411},
		{5, 1, 0.01224269},
		{10, 1, 0.003151583},

		{-4, 2, 0.01309457},
		{-3, 2, 0.02741012},
		{-2, 2, 0.06804138},
		{-1, 2, 0.1924501},
		{0, 2, 0.3535534},
		{1, 2, 0.1924501},
		{2, 2, 0.06804138},
		{3, 2, 0.02741012},
		{4, 2, 0.01309457},
		{5, 2, 0.007127781},
		{10, 2, 0.0009707329},

		{-4, 6, 0.004054578},
		{-3, 6, 0.01549193},
		{-2, 6, 0.06403612},
		{-1, 6, 0.2231423},
		{0, 6, 0.3827328},
		{1, 6, 0.2231423},
		{2, 6, 0.06403612},
		{3, 6, 0.01549193},
		{4, 6, 0.004054578},
		{5, 6, 0.001220841},
		{10, 6, 1.651408e-05},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := StudentT{c.ν, nil}

			res := b.Probability(c.t)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R pt(t, ν) from package stats
func TestStudentTDistribution(t *testing.T) {
	tol := 0.0000001

	cases := []struct {
		t, ν, expected float64
	}{
		{-4, 1, 0.07797913},
		{-3, 1, 0.1024164},
		{-2, 1, 0.1475836},
		{-1, 1, 0.25},
		{0, 1, 0.5},
		{1, 1, 0.75},
		{2, 1, 0.8524164},
		{3, 1, 0.8975836},
		{4, 1, 0.9220209},
		{5, 1, 0.937167},
		{10, 1, 0.9682745},

		{-4, 2, 0.02859548},
		{-3, 2, 0.04773298},
		{-2, 2, 0.09175171},
		{-1, 2, 0.2113249},
		{0, 2, 0.5},
		{1, 2, 0.7886751},
		{2, 2, 0.9082483},
		{3, 2, 0.952267},
		{4, 2, 0.9714045},
		{5, 2, 0.9811252},
		{10, 2, 0.9950738},

		{-4, 6, 0.003559489},
		{-3, 6, 0.0120041},
		{-2, 6, 0.04621316},
		{-1, 6, 0.1779588},
		{0, 6, 0.5},
		{1, 6, 0.8220412},
		{2, 6, 0.9537868},
		{3, 6, 0.9879959},
		{4, 6, 0.9964405},
		{5, 6, 0.9987738},
		{10, 6, 0.999971},

		{-2, 3, 0.06966298},
		{0.1, 2, 0.5352673},
		{2.9, 2, 0.9494099},
		{3.9, 6, 0.996008},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := StudentT{c.ν, nil}

			res := b.Distribution(c.t)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestStudentTInverseDistribution(t *testing.T) {
	tol := 0.0000001

	cases := []struct {
		t, ν, expected float64
	}{
		{-4, 1, 0.07797913},
		{-3, 1, 0.1024164},
		{-2, 1, 0.1475836},
		{-1, 1, 0.25},
		{0, 1, 0.5},
		{1, 1, 0.75},
		{2, 1, 0.8524164},
		{3, 1, 0.8975836},
		{4, 1, 0.9220209},
		{5, 1, 0.937167},
		{10, 1, 0.9682745},
		{-4, 2, 0.02859548},
		{-3, 2, 0.04773298},
		{-2, 2, 0.09175171},
		{-1, 2, 0.2113249},
		{0, 2, 0.5},
		{1, 2, 0.7886751},
		{2, 2, 0.9082483},
		{3, 2, 0.952267},
		{4, 2, 0.9714045},
		{5, 2, 0.9811252},
		{10, 2, 0.9950738},
		{-4, 6, 0.003559489},
		{-3, 6, 0.0120041},
		{-2, 6, 0.04621316},
		{-1, 6, 0.1779588},
		{0, 6, 0.5},
		{1, 6, 0.8220412},
		{2, 6, 0.9537868},
		{3, 6, 0.9879959},
		{4, 6, 0.9964405},
		{5, 6, 0.9987738},
		{10, 6, 0.999971},
		{-2, 3, 0.06966298},
		{0.1, 2, 0.5352673},
		{2.9, 2, 0.9494099},
		{3.9, 6, 0.996008},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := StudentT{c.ν, nil}

			res := b.Distribution(c.t)
			inverse := b.Inverse(res)
			if math.Abs(c.t-inverse) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.t, inverse)
			}

		})
	}
}

func TestStudentTInverse(t *testing.T) {
	tol := 0.000001

	cases := []struct {
		p, ν, expected float64
	}{
		{0.6, 1, 0.3249196962},
		{0.6, 2, 0.2886751346},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := StudentT{c.ν, nil}

			res := b.Inverse(c.p)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestStudentTMean(t *testing.T) {

	cases := []struct {
		ν, expected float64
	}{
		{2, 0},
		{3, 0},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := StudentT{c.ν, nil}

			res := b.Mean()
			if res != c.expected {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestStudentTMeanNaN(t *testing.T) {
	b := StudentT{1, nil}
	res := b.Mean()
	if !math.IsNaN(res) {
		t.Errorf("Mismatch. want: %v, got: %v", math.NaN(), res)
	}
}

func TestStudentTMedian(t *testing.T) {

	cases := []struct {
		ν, expected float64
	}{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := StudentT{c.ν, nil}

			res := b.Median()
			if res != c.expected {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestStudentTMode(t *testing.T) {

	cases := []struct {
		ν, expected float64
	}{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := StudentT{c.ν, nil}

			res := b.Mode()
			if res != c.expected {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestStudentTVariance(t *testing.T) {

	cases := []struct {
		ν, expected float64
	}{
		{1.1, math.Inf(1)},
		{1.5, math.Inf(1)},
		{2, math.Inf(1)},
		{3, 3},
		{4, 2},
		{5, 5. / 3},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := StudentT{c.ν, nil}

			res := b.Variance()
			if res != c.expected {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestStudentTVarianceNaN(t *testing.T) {

	cases := []struct {
		ν float64
	}{
		{0.1},
		{0.5},
		{0.9},
		{1},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := StudentT{c.ν, nil}

			res := b.Variance()
			if !math.IsNaN(res) {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, math.NaN(), res)
			}

		})
	}
}
