package continuous

import (
	"math"
	"strconv"
	"testing"
)

func TestFProbability(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		x        float64
		d1, d2   int
		expected float64
	}{
		{1, 1, 1, 0.1591549},
		{2, 1, 1, 0.07502636},
		{3, 1, 1, 0.04594407},
		{4, 1, 1, 0.03183099},
		{5, 1, 1, 0.02372542},
		{10, 1, 1, 0.009150766},

		{1, 2, 1, 0.1924501},
		{2, 2, 1, 0.08944272},
		{3, 2, 1, 0.05399492},
		{4, 2, 1, 0.03703704},
		{5, 2, 1, 0.02741012},
		{10, 2, 1, 0.01039133},

		{1, 1, 2, 0.1924501},
		{2, 1, 2, 0.08838835},
		{3, 1, 2, 0.05163978},
		{4, 1, 2, 0.03402069},
		{5, 1, 2, 0.02414726},
		{10, 1, 2, 0.007607258},

		{1, 2, 2, 0.25},
		{2, 2, 2, 0.1111111},
		{3, 2, 2, 0.0625},
		{4, 2, 2, 0.04},
		{5, 2, 2, 0.02777778},
		{10, 2, 2, 0.008264463},

		{5, 3, 7, 0.01667196},
		{7, 6, 2, 0.016943},
		{7, 20, 14, 0.0002263343},
		{45, 2, 3, 0.0001868942},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := F{d1: c.d1, d2: c.d2}

			res := f.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestFDistribution(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		x        float64
		d1, d2   int
		expected float64
	}{
		{0, 1, 1, 0},
		{0, 1, 2, 0},
		{0, 2, 1, 0},
		{0, 2, 2, 0},
		{0, 2, 3, 0},

		{1, 1, 1, 0.5},
		{2, 1, 1, 0.6081734},
		{3, 1, 1, 0.6666667},
		{4, 1, 1, 0.7048328},
		{5, 1, 1, 0.7322795},
		{10, 1, 1, 0.8050178},

		{1, 2, 1, 0.4226497},
		{2, 2, 1, 0.5527864},
		{3, 2, 1, 0.6220355},
		{4, 2, 1, 0.6666667},
		{5, 2, 1, 0.6984887},
		{10, 2, 1, 0.7817821},

		{1, 1, 2, 0.5773503},
		{2, 1, 2, 0.7071068},
		{3, 1, 2, 0.7745967},
		{4, 1, 2, 0.8164966},
		{5, 1, 2, 0.8451543},
		{10, 1, 2, 0.9128709},

		{1, 2, 2, 0.5},
		{2, 2, 2, 0.6666667},
		{3, 2, 2, 0.75},
		{4, 2, 2, 0.8},
		{5, 2, 2, 0.8333333},
		{10, 2, 2, 0.9090909},

		{5, 3, 7, 0.9633266},
		{7, 6, 2, 0.8697408},
		{7, 20, 14, 0.9997203},
		{45, 2, 3, 0.9942063},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := F{d1: c.d1, d2: c.d2}

			res := f.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestFInverse(t *testing.T) {

	tol := 0.001

	cases := []struct {
		x        float64
		d1, d2   int
		expected float64
	}{
		{0, 1, 1, 0},
		{0, 1, 2, 0},
		{0, 2, 1, 0},
		{0, 2, 2, 0},
		{0, 2, 3, 0},

		{0.5, 1, 1, 1},
		{0.6081734, 1, 1, 2},
		{0.6666667, 1, 1, 3},
		{0.7048328, 1, 1, 4},
		{0.7322795, 1, 1, 5},
		{0.8050178, 1, 1, 10},

		{0.4226497, 2, 1, 1},
		{0.5527864, 2, 1, 2},
		{0.6220355, 2, 1, 3},
		{0.6666667, 2, 1, 4},
		{0.6984887, 2, 1, 5},
		{0.7817821, 2, 1, 10},

		{0.5773503, 1, 2, 1},
		{0.7071068, 1, 2, 2},
		{0.7745967, 1, 2, 3},
		{0.8164966, 1, 2, 4},
		{0.8451543, 1, 2, 5},
		{0.9128709, 1, 2, 10},

		{0.5, 2, 2, 1},
		{0.6666667, 2, 2, 2},
		{0.75, 2, 2, 3},
		{0.8, 2, 2, 4},
		{0.8333333, 2, 2, 5},
		{0.9090909, 2, 2, 10},

		{0.9633266, 3, 7, 5},
		{0.8697408, 6, 2, 7},
		{0.9997203, 20, 14, 7},
		{0.9942063, 2, 3, 45},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := F{d1: c.d1, d2: c.d2}

			res := f.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestFMean(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		d1, d2   int
		expected float64
	}{
		{1, 3, 3},
		{1, 4, 2},
		{1, 5, 1.66666667},
		{1, 6, 1.5},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := F{d1: c.d1, d2: c.d2}

			res := f.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestFMeanNaN(t *testing.T) {

	cases := []struct {
		d1, d2 int
	}{
		{1, 1},
		{1, 2},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := F{d1: c.d1, d2: c.d2}

			res := f.Mean()
			if !math.IsNaN(res) {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, math.NaN(), res)
			}

		})
	}
}

func TestFMode(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		d1, d2   int
		expected float64
	}{
		{3, 1, 0.11111111},
		{3, 2, 0.16666667},
		{3, 3, 0.2},
		{3, 4, 0.22222222},
		{4, 1, 0.16666667},
		{4, 2, 0.25},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := F{d1: c.d1, d2: c.d2}

			res := f.Mode()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestFModeNaN(t *testing.T) {

	cases := []struct {
		d1, d2 int
	}{
		{1, 1},
		{1, 5},
		{2, 2},
		{2, 4},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := F{d1: c.d1, d2: c.d2}

			res := f.Mode()
			if !math.IsNaN(res) {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, math.NaN(), res)
			}

		})
	}
}

func TestFVariance(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		d1, d2   int
		expected float64
	}{
		{1, 5, 22.22222222},
		{2, 5, 13.88888889},
		{3, 5, 11.11111111},
		{4, 5, 9.72222222},
		{5, 5, 8.88888889},
		{6, 5, 8.33333333},
		{5, 7, 2.61333333},
		{9, 8, 1.48148148},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := F{d1: c.d1, d2: c.d2}

			res := f.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestFVarianceNaN(t *testing.T) {

	cases := []struct {
		d1, d2 int
	}{
		{1, 1},
		{1, 2},
		{2, 3},
		{5, 4},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := F{d1: c.d1, d2: c.d2}

			res := f.Variance()
			if !math.IsNaN(res) {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, math.NaN(), res)
			}

		})
	}
}

func TestFSkewness(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		d1, d2   int
		expected float64
	}{
		{5, 52, 1.54132},
		{3, 1764, 1.63995},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := F{d1: c.d1, d2: c.d2}

			res := f.Skewness()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestFExKurtosis(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		d1, d2   int
		expected float64
	}{
		{5, 52, 3.9982},
		{3, 1764, 4.0456},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := F{d1: c.d1, d2: c.d2}

			res := f.ExKurtosis()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
