package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dunif(x, min, max)
func TestUniformProbability(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		a, b, x  float64
		expected float64
	}{
		{0, 1, -1, 0},
		{0, 1, 0, 1},
		{0, 1, 0.5, 1},
		{0, 1, 1, 1},
		{0, 1, 2, 0},

		{0, 2, -1, 0},
		{0, 2, 0, 0.5},
		{0, 2, 0.5, 0.5},
		{0, 2, 1, 0.5},
		{0, 2, 2, 0.5},
		{0, 2, 3, 0},

		{1, 4, 2, 0.3333333},
		{1, 4, 3.4, 0.3333333},
		{1, 5.4, 3, 0.2272727},
		{1, 5.4, 0.3, 0},
		{1, 5.4, 6, 0},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Uniform{c.a, c.b, nil}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestUniformDistribution(t *testing.T) {
	tol := 0.001
	cases := []struct {
		a, b, x  float64
		expected float64
	}{
		{0, 1, -1, 0},
		{0, 1, 0, 0},
		{0, 1, 0.5, 0.5},
		{0, 1, 1, 1},
		{0, 1, 2, 1},

		{0, 2, -1, 0},
		{0, 2, 0, 0},
		{0, 2, 0.5, 0.25},
		{0, 2, 1, 0.5},
		{0, 2, 2, 1},
		{0, 2, 3, 1},

		{1, 4, 2, 0.3333333},
		{1, 4, 3.4, 0.8},
		{1, 5.4, 3, 0.4545455},
		{1, 5.4, 0.3, 0},
		{1, 5.4, 6, 1},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Uniform{c.a, c.b, nil}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestUniformMean(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{0, 1, 0.5},
		{0, 2, 1},
		{1, 2, 1.5},
		{2, 3, 5. / 2},
		{2, 4, 3},
		{5, 11, 8},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Uniform{c.a, c.b, nil}
			res := b.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestUniformMedian(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{0, 1, 0.5},
		{0, 2, 1},
		{1, 2, 1.5},
		{2, 3, 5. / 2},
		{2, 4, 3},
		{5, 11, 8},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Uniform{c.a, c.b, nil}

			res := b.Median()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestUniformMode(t *testing.T) {

	cases := []struct {
		a, b     float64
		expected float64
	}{
		{0, 1, 0.5},
		{0, 2, 1},
		{1, 2, 1.5},
		{2, 3, 5. / 2},
		{2, 4, 3},
		{5, 11, 8},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Uniform{c.a, c.b, nil}

			res := b.Mode()
			if res < c.a || res > c.b {
				t.Errorf("Mismatch. Case %d, should be greater than or equal to a or less than or equal to b.", i)
			}

		})
	}
}

func TestUniformVariance(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{0, 1, 1. / 12},
		{0, 2, 4. / 12},
		{1, 2, 1. / 12},
		{2, 3, 1. / 12},
		{2, 4, 4. / 12},
		{5, 11, 3},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Uniform{c.a, c.b, nil}

			res := b.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
