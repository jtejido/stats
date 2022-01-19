package continuous

import (
	"math"
	"strconv"
	"testing"
)

// https://www.saecanet.com/20100716/saecanet_calculation_page.php?pagenumber=628
func TestErlangProbability(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		x        float64
		k        int
		λ        float64
		expected float64
	}{
		{1, 1, 1, 0.36788},
		{3, 10, 1, 0.0027},
		{5, 11, 1, 0.01813},
		{5, 11, .5, 0.0001},
		{3, 1, .25, 0.1181},
		{3, 1, 2, 0.005},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Erlang{shape: c.k, rate: c.λ}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// https://www.saecanet.com/20100716/saecanet_calculation_page.php?pagenumber=628
func TestErlangDistribution(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		x        float64
		k        int
		λ        float64
		expected float64
	}{
		{1, 1, 1, 0.6321},
		{3, 10, 1, 0.0011},
		{5, 11, 1, 0.0137},
		{5, 11, .5, 0.0001},
		{3, 1, .25, 0.5276},
		{3, 1, 2, 0.9975},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Erlang{shape: c.k, rate: c.λ}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestErlangInverse(t *testing.T) {
	tol := 0.01
	cases := []struct {
		x        float64
		k        int
		λ        float64
		expected float64
	}{
		{0.6321, 1, 1, 1},
		{0.0011, 10, 1, 3},
		{0.0137, 11, 1, 5},
		{0.5276, 1, .25, 3},
		{0.9975, 1, 2, 3},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Erlang{shape: c.k, rate: c.λ}

			res := b.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// https://www.saecanet.com/20100716/saecanet_calculation_page.php?pagenumber=628
func TestErlangMean(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		k        int
		λ        float64
		expected float64
	}{
		{1, 1, 1},
		{10, 1, 10},
		{11, 1, 11},
		{11, .5, 22},
		{1, .25, 4},
		{1, 2, 0.5},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Erlang{shape: c.k, rate: c.λ}

			res := b.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Skewness[ErlangDistribution[k,λ]]
func TestErlangSkewness(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		k        int
		λ        float64
		expected float64
	}{
		{1, .25, 2},
		{8, 267.439, 0.707107},
		{141, 8.765131987, 0.16843},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Erlang{shape: c.k, rate: c.λ}
			res := b.Skewness()

			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Mode[ErlangDistribution[k,λ]]
func TestErlangMode(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		k        int
		λ        float64
		expected float64
	}{
		{1, .25, 0},
		{8, 267.439, 0.0261742},
		{141, 8.765131987, 15.9724},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Erlang{shape: c.k, rate: c.λ}
			res := b.Mode()

			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Variance[ErlangDistribution[k,λ]]
func TestErlangVariance(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		k        int
		λ        float64
		expected float64
	}{
		{1, .25, 16},
		{8, 267.439, 0.000111851},
		{141, 8.765131987, 1.83528},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Erlang{shape: c.k, rate: c.λ}
			res := b.Variance()

			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
