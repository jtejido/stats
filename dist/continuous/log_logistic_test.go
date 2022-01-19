package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dllogis(x, scale = α, shape = β) [From eha package]
func TestLogLogisticProbability(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, α, β  float64
		expected float64
	}{
		{0, 1, 1, 1},
		{1, 1, 1, 0.25},
		{2, 1, 1, 0.1111111},
		{3, 1, 1, 0.0625},
		{4, 1, 1, 0.04},
		{5, 1, 1, 0.02777778},
		{10, 1, 1, 0.008264463},

		{0, 1, 2, 0},
		{1, 1, 2, 0.5},
		{2, 1, 2, 0.16},
		{3, 1, 2, 0.06},
		{4, 1, 2, 0.02768166},
		{5, 1, 2, 0.0147929},
		{10, 1, 2, 0.001960592},

		{0, 2, 2, 0},
		{1, 2, 2, 0.32},
		{2, 2, 2, 0.25},
		{3, 2, 2, 0.1420118},
		{4, 2, 2, 0.08},
		{5, 2, 2, 0.04756243},
		{10, 2, 2, 0.00739645},

		{4, 2, 3, 0.07407407},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogLogistic{scale: c.α, shape: c.β, location: 0}

			res := e.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R pllogis(x, scale = α, shape = β) [From eha package]
func TestLogLogisticDistribution(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, α, β  float64
		expected float64
	}{
		{0, 1, 1, 0},
		{1, 1, 1, 0.5},
		{2, 1, 1, 0.6666667},
		{3, 1, 1, 0.75},
		{4, 1, 1, 0.8},
		{5, 1, 1, 0.8333333},
		{10, 1, 1, 0.9090909},

		{0, 1, 2, 0},
		{1, 1, 2, 0.5},
		{2, 1, 2, 0.8},
		{3, 1, 2, 0.9},
		{4, 1, 2, 0.9411765},
		{5, 1, 2, 0.9615385},
		{10, 1, 2, 0.990099},

		{0, 2, 2, 0},
		{1, 2, 2, 0.2},
		{2, 2, 2, 0.5},
		{3, 2, 2, 0.6923077},
		{4, 2, 2, 0.8},
		{5, 2, 2, 0.862069},
		{10, 2, 2, 0.9615385},

		{4, 2, 3, 0.8888889},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogLogistic{scale: c.α, shape: c.β, location: 0}

			res := e.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLogLogisticInverse(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, α, β  float64
		expected float64
	}{
		{0, 1, 1, 0},
		{1, 1, 1, 0.5},
		{2, 1, 1, 0.6666667},
		{3, 1, 1, 0.75},
		{4, 1, 1, 0.8},
		{5, 1, 1, 0.8333333},
		{10, 1, 1, 0.9090909},

		{0, 1, 2, 0},
		{1, 1, 2, 0.5},
		{2, 1, 2, 0.8},
		{3, 1, 2, 0.9},
		{4, 1, 2, 0.9411765},
		{5, 1, 2, 0.9615385},
		{10, 1, 2, 0.990099},

		{0, 2, 2, 0},
		{1, 2, 2, 0.2},
		{2, 2, 2, 0.5},
		{3, 2, 2, 0.6923077},
		{4, 2, 2, 0.8},
		{5, 2, 2, 0.862069},
		{10, 2, 2, 0.9615385},

		{4, 2, 3, 0.8888889},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogLogistic{scale: c.α, shape: c.β, location: 0}

			cdf := e.Distribution(c.x)
			inverse := e.Inverse(cdf)
			if math.Abs(inverse-c.x) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.x, inverse)
			}

		})
	}
}

func TestLogLogisticMean(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		α, β     float64
		expected float64
	}{
		{1, 2, 1.570795},
		{2, 2, 3.14159},
		{3, 3, 3.62759751692},
		{5, 4, 5.55360266602},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogLogistic{scale: c.α, shape: c.β, location: 0}

			res := e.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLogLogisticMeanNaN(t *testing.T) {

	cases := []struct {
		α, β float64
	}{
		{1, 1},
		{2, 1},
		{3, 1},
		{5, 1},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogLogistic{scale: c.α, shape: c.β, location: 0}

			res := e.Mean()
			if !math.IsNaN(res) {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, math.NaN(), res)
			}

		})
	}
}

func TestLogLogisticMedian(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		α, β     float64
		expected float64
	}{
		{1, 2, 1.570795},
		{2, 2, 3.14159},
		{3, 3, 3.62759751692},
		{5, 4, 5.55360266602},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogLogistic{scale: c.α, shape: c.β, location: 0}

			res := e.Median()
			if math.Abs(res-c.α) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.α, res)
			}

		})
	}
}

func TestLogLogisticMode(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		α, β     float64
		expected float64
	}{
		{1, 0.2, 0},
		{2, 0.9, 0},
		{3, 1, 0},
		{1, 2, 0.577350269189623},
		{1, 3, 0.793700525984102},
		{2, 3, 1.5874010519682},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogLogistic{scale: c.α, shape: c.β, location: 0}

			res := e.Mode()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLogLogisticVariance(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		α, β     float64
		expected float64
	}{
		{1, 3, 0.956236},
		{2, 4, 1.34838},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogLogistic{scale: c.α, shape: c.β, location: 0}

			res := e.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLogLogisticVarianceNaN(t *testing.T) {

	cases := []struct {
		α, β float64
	}{
		{1, 1},
		{2, 2},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogLogistic{scale: c.α, shape: c.β, location: 0}

			res := e.Variance()
			if !math.IsNaN(res) {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, math.NaN(), res)
			}

		})
	}
}

func TestLogLogisticSkewness(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		α, β     float64
		expected float64
	}{
		{1, 3, math.NaN()},
		{2, 4, 4.28478},
		{6.87623, 11.8232, 0.775272},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogLogistic{scale: c.α, shape: c.β, location: 0}

			res := e.Skewness()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLogLogisticExKurtosis(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		α, β     float64
		expected float64
	}{
		{11, 5, 26.5562},
		{6.87623, 11.8232, 2.74979},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogLogistic{scale: c.α, shape: c.β, location: 0}

			res := e.ExKurtosis()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
