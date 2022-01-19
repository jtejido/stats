package continuous

import (
	"math"
	"strconv"
	"testing"
)

// same test as Pareto with L as xm and H as +Inf
func TestParetoBoundedProbability(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x, α, a, b float64
		expected   float64
	}{

		{1, 1, 1, math.Inf(1), 1},
		{2, 1, 1, math.Inf(1), 0.25},
		{3, 1, 1, math.Inf(1), 0.1111111},
		{4, 1, 1, math.Inf(1), 0.0625},
		{5, 1, 1, math.Inf(1), 0.04},
		{10, 1, 1, math.Inf(1), 0.01},

		{1, 2, 1, math.Inf(1), 2},
		{2, 2, 1, math.Inf(1), 0.25},
		{3, 2, 1, math.Inf(1), 0.07407407},
		{4, 2, 1, math.Inf(1), 0.03125},
		{5, 2, 1, math.Inf(1), 0.016},
		{10, 2, 1, math.Inf(1), 0.002},

		{2, 1, 2, math.Inf(1), 0.5},
		{3, 1, 2, math.Inf(1), 0.2222222},
		{4, 1, 2, math.Inf(1), 0.125},
		{5, 1, 2, math.Inf(1), 0.08},
		{10, 1, 2, math.Inf(1), 0.02},

		{2, 2, 2, math.Inf(1), 1},
		{3, 2, 2, math.Inf(1), 0.2962963},
		{4, 2, 2, math.Inf(1), 0.125},
		{5, 2, 2, math.Inf(1), 0.064},
		{10, 2, 2, math.Inf(1), 0.008},

		{4, 8, 2, math.Inf(1), 0.0078125},
		{5, 8, 2, math.Inf(1), 0.001048576},
		{9, 4, 5, math.Inf(1), 0.04233772},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ParetoBounded{c.a, c.b, c.α, nil}

			res := b.Probability(c.x)
			run_test(t, res, c.expected, tol, "ParetoBoundedProbability")

		})
	}
}

// same test as Pareto with L as xm and H as +Inf
func TestParetoBoundedDistribution(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x, α, a, b float64
		expected   float64
	}{
		{1, 1, 1, math.Inf(1), 0},
		{2, 1, 1, math.Inf(1), 0.5},
		{3, 1, 1, math.Inf(1), 0.6666667},
		{4, 1, 1, math.Inf(1), 0.75},
		{5, 1, 1, math.Inf(1), 0.8},
		{10, 1, 1, math.Inf(1), 0.9},
		{1, 2, 1, math.Inf(1), 0},
		{2, 2, 1, math.Inf(1), 0.75},
		{3, 2, 1, math.Inf(1), 0.8888889},
		{4, 2, 1, math.Inf(1), 0.9375},
		{5, 2, 1, math.Inf(1), 0.96},
		{10, 2, 1, math.Inf(1), 0.99},
		{2, 1, 2, math.Inf(1), 0},
		{3, 1, 2, math.Inf(1), 0.3333333},
		{4, 1, 2, math.Inf(1), 0.5},
		{5, 1, 2, math.Inf(1), 0.6},
		{10, 1, 2, math.Inf(1), 0.8},
		{2, 2, 2, math.Inf(1), 0},
		{3, 2, 2, math.Inf(1), 0.5555556},
		{4, 2, 2, math.Inf(1), 0.75},
		{5, 2, 2, math.Inf(1), 0.84},
		{10, 2, 2, math.Inf(1), 0.96},
		{4, 8, 2, math.Inf(1), 0.9960938},
		{5, 8, 2, math.Inf(1), 0.9993446},
		{9, 4, 5, math.Inf(1), 0.9047401},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ParetoBounded{c.a, c.b, c.α, nil}

			res := b.Distribution(c.x)
			run_test(t, res, c.expected, tol, "ParetoBoundedDistribution")

		})
	}
}

func TestParetoBoundedMean(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		α, a, b  float64
		expected float64
	}{
		{2, 1, math.Inf(1), 2},
		{3, 1, math.Inf(1), 1.5},
		{3, 2, math.Inf(1), 3},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ParetoBounded{c.a, c.b, c.α, nil}

			res := b.Mean()
			run_test(t, res, c.expected, tol, "ParetoBoundedMean")

		})
	}
}

func TestParetoBoundedVariance(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		α, a, b  float64
		expected float64
	}{
		{3, 1, math.Inf(1), 0.75},
		{3, 2, math.Inf(1), 3},
		{4, 3, math.Inf(1), 2},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ParetoBounded{c.a, c.b, c.α, nil}

			res := b.Variance()
			run_test(t, res, c.expected, tol, "ParetoBoundedVariance")

		})
	}
}
