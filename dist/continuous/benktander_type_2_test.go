package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with Mathematica PDF[BenktanderWeibullDistribution[a, b], x]
func TestBenktanderType2Probability(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		x, a, b  float64
		expected float64
	}{
		{1.23424, 1, 1, 0.791172},
		{4.2346, 2, .5, 0.0076984},
		{7.23243, 4, .1, 0.0000212899},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BenktanderType2{a: c.a, b: c.b}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica CDF[BenktanderWeibullDistribution[a, b], x]
func TestBenktanderType2Distribution(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		x, a, b  float64
		expected float64
	}{
		{1.23424, 1, 1, 0.208828},
		{4.2346, 2, .5, 0.992937},
		{7.23243, 4, .1, 0.999973},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BenktanderType2{a: c.a, b: c.b}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Variance[BenktanderWeibullDistribution[a, b]]
func TestBenktanderType2Variance(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{1, 1, 1},
		{2, .5, 0.375},
		{4, .1, 0.0974025},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BenktanderType2{a: c.a, b: c.b}

			res := b.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Mean[BenktanderWeibullDistribution[a, b]]
func TestBenktanderType2Mean(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{2, .5, 1.5},
		{4, .1, 1.25},
		{10, .7, 1.1},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BenktanderType2{a: c.a, b: c.b}

			res := b.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Median[BenktanderWeibullDistribution[a, b]]
func TestBenktanderType2Median(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{2, .5, 1.30059},
		{4, .1, 1.15102},
		{10, .7, 1.06801},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BenktanderType2{a: c.a, b: c.b}

			res := b.Median()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestBenktanderType2ExKurtosis(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{2, .5, 20.8333},
		{4, .1, 37.577},
		{10, .7, 7.35857},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BenktanderType2{a: c.a, b: c.b}

			res := b.ExKurtosis()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestBenktanderType2Skewness(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{2, .5, -3.33403},
		{4, .1, -3.95765},
		{10, .7, -2.16493},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BenktanderType2{a: c.a, b: c.b}

			res := b.Skewness()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
