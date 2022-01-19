package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with Mathematica PDF[BetaPrimeDistribution[p, q], x]
func TestBetaPrimeProbability(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		x, a, b  float64
		expected float64
	}{
		{0.99, 5.3, 4.1, 0.548823},
		{0.29, 2, 3, 0.974161},
		{0.99, 10, 50, 6.72493e-7},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BetaPrime{alpha: c.a, beta: c.b}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica CDF[BetaPrimeDistribution[p, q], x]
func TestBetaPrimeDistribution(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		x, a, b  float64
		expected float64
	}{
		{0.21, 1, 1, 0.173554},
		{0.78, 2, 3, 0.589591},
		{0.53, 10, 50, 0.999232},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BetaPrime{alpha: c.a, beta: c.b}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestBetaPrimeInverse(t *testing.T) {
	tol := 0.001
	cases := []struct {
		x, a, b  float64
		expected float64
	}{
		{0.173554, 1, 1, 0.21},
		{0.589591, 2, 3, 0.78},
		{0.999232, 10, 50, 0.53},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BetaPrime{alpha: c.a, beta: c.b}

			res := b.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Mean[BetaPrimeDistribution[p, q]]
func TestBetaPrimeMean(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{1, 2, 1},
		{20, 3, 10},
		{10, 50, 0.204082},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BetaPrime{alpha: c.a, beta: c.b}

			res := b.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Median[BetaPrimeDistribution[p, q]]
func TestBetaPrimeMedian(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{1, 2, 0.414214},
		{20, 3, 7.35394},
		{10, 50, 0.194673},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BetaPrime{alpha: c.a, beta: c.b}

			res := b.Median()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Variance[BetaPrimeDistribution[p, q]]
func TestBetaPrimeVariance(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{5.3, 7.85, 0.234593},
		{20, 3, 110},
		{10, 50, 0.00511939},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BetaPrime{alpha: c.a, beta: c.b}

			res := b.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Skewness[BetaPrimeDistribution[p, q]]
func TestBetaPrimeSkewness(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{5.3, 7.85, 2.16888},
		{10, 50, 0.837483},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BetaPrime{alpha: c.a, beta: c.b}

			res := b.Skewness()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
