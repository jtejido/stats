package continuous

import (
	"math"
	"strconv"
	"testing"
)

func TestDagumProbability(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		x, p, a, scale float64 // p,a,b
		expected       float64
	}{
		{.1, 1, 2, 3, 0.0221729},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Dagum{p: c.p, a: c.a, scale: c.scale}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestDagumDistribution(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		x, p, a, scale float64 // p,a,b
		expected       float64
	}{
		{.1, 1, 2, 3, 0.00110988},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Dagum{p: c.p, a: c.a, scale: c.scale}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestDagumInverse(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		x, p, a, scale float64 // p,a,b
		expected       float64
	}{
		{0.00110988, 1, 2, 3, .1},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Dagum{p: c.p, a: c.a, scale: c.scale}

			res := b.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestDagumMean(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		p, a, scale float64
		expected    float64
	}{
		{2, 10, 5, 5.59152},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Dagum{p: c.p, a: c.a, scale: c.scale}

			res := b.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestDagumVariance(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		p, a, scale float64
		expected    float64
	}{
		{2, 10, 5, 0.803639},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Dagum{p: c.p, a: c.a, scale: c.scale}

			res := b.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestDagumSkewness(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		p, a, scale float64
		expected    float64
	}{
		{2, 10, 5, 1.39721},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Dagum{p: c.p, a: c.a, scale: c.scale}

			res := b.Skewness()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestDagumExKurtosis(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		p, a, scale float64
		expected    float64
	}{
		{2, 10, 5, 5.30746},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Dagum{p: c.p, a: c.a, scale: c.scale}

			res := b.ExKurtosis()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
