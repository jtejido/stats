package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dpareto(x, b a) from package EnvStats
func TestParetoProbability(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x, a, b  float64
		expected float64
	}{

		{1, 1, 1, 1},
		{2, 1, 1, 0.25},
		{3, 1, 1, 0.1111111},
		{4, 1, 1, 0.0625},
		{5, 1, 1, 0.04},
		{10, 1, 1, 0.01},

		{1, 2, 1, 2},
		{2, 2, 1, 0.25},
		{3, 2, 1, 0.07407407},
		{4, 2, 1, 0.03125},
		{5, 2, 1, 0.016},
		{10, 2, 1, 0.002},

		{2, 1, 2, 0.5},
		{3, 1, 2, 0.2222222},
		{4, 1, 2, 0.125},
		{5, 1, 2, 0.08},
		{10, 1, 2, 0.02},

		{2, 2, 2, 1},
		{3, 2, 2, 0.2962963},
		{4, 2, 2, 0.125},
		{5, 2, 2, 0.064},
		{10, 2, 2, 0.008},

		{4, 8, 2, 0.0078125},
		{5, 8, 2, 0.001048576},
		{9, 4, 5, 0.04233772},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Pareto{c.a, c.b, nil, nil}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R ppareto(x, b a) from package EnvStats
func TestParetoDistribution(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x, a, b  float64
		expected float64
	}{

		{1, 1, 1, 0},
		{2, 1, 1, 0.5},
		{3, 1, 1, 0.6666667},
		{4, 1, 1, 0.75},
		{5, 1, 1, 0.8},
		{10, 1, 1, 0.9},

		{1, 2, 1, 0},
		{2, 2, 1, 0.75},
		{3, 2, 1, 0.8888889},
		{4, 2, 1, 0.9375},
		{5, 2, 1, 0.96},
		{10, 2, 1, 0.99},

		{2, 1, 2, 0},
		{3, 1, 2, 0.3333333},
		{4, 1, 2, 0.5},
		{5, 1, 2, 0.6},
		{10, 1, 2, 0.8},

		{2, 2, 2, 0},
		{3, 2, 2, 0.5555556},
		{4, 2, 2, 0.75},
		{5, 2, 2, 0.84},
		{10, 2, 2, 0.96},

		{4, 8, 2, 0.9960938},
		{5, 8, 2, 0.9993446},
		{9, 4, 5, 0.9047401},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Pareto{c.a, c.b, nil, nil}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with https://solvemymath.com/online_math_calculator/statistics/continuous_distributions/pareto/quantile_pareto.php
func TestParetoInverse(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		p, a, b  float64
		expected float64
	}{
		{0.1, 1, 1, 1.1111111},
		{0.2, 1, 1, 1.25},
		{0.3, 1, 1, 1.42857},
		{0.4, 1, 1, 1.666666667},
		{0.5, 1, 1, 2},
		{0.6, 1, 1, 2.5},
		{0.7, 1, 1, 3.3333333},
		{0.8, 1, 1, 5},
		{0.9, 1, 1, 10},
		{0.1, 2, 2, 2.108185},
		{0.2, 2, 2, 2.2360679},
		{0.3, 2, 2, 2.390457218},
		{0.4, 2, 2, 2.5819888974},
		{0.5, 2, 2, 2.8284271247},
		{0.6, 2, 2, 3.1622776601},
		{0.7, 2, 2, 3.6514837167},
		{0.8, 2, 2, 4.4721359549},
		{0.9, 2, 2, 6.3245553203},
		{0.1, 4, 6, 6.1601405764},
		{0.2, 4, 6, 6.3442275806},
		{0.5, 4, 6, 7.1352426900},
		{0.9, 4, 6, 10.669676460},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Pareto{c.a, c.b, nil, nil}

			res := b.Inverse(c.p)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestParetoInverseDistribution(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x, a, b  float64
		expected float64
	}{
		{2, 1, 1, 0.5},
		{3, 1, 1, 0.6666667},
		{4, 1, 1, 0.75},
		{5, 1, 1, 0.8},
		{10, 1, 1, 0.9},

		{2, 2, 1, 0.75},
		{3, 2, 1, 0.8888889},
		{4, 2, 1, 0.9375},
		{5, 2, 1, 0.96},
		{10, 2, 1, 0.99},

		{3, 1, 2, 0.3333333},
		{4, 1, 2, 0.5},
		{5, 1, 2, 0.6},
		{10, 1, 2, 0.8},

		{3, 2, 2, 0.5555556},
		{4, 2, 2, 0.75},
		{5, 2, 2, 0.84},
		{10, 2, 2, 0.96},

		{4, 8, 2, 0.9960938},
		{5, 8, 2, 0.9993446},
		{9, 4, 5, 0.9047401},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Pareto{c.a, c.b, nil, nil}

			res := b.Distribution(c.x)
			inverse := b.Inverse(res)
			if math.Abs(c.x-inverse) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.x, inverse)
			}

		})
	}
}

func TestParetoMean(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{1, 2, math.Inf(1)},
		{0.4, 2, math.Inf(1)},
		{0.001, 2, math.Inf(1)},
		{2, 1, 2},
		{3, 1, 1.5},
		{3, 2, 3},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Pareto{c.a, c.b, nil, nil}

			res := b.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestParetoMedian(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{1, 1, 2},
		{1, 2, 1.414213562373095},
		{2, 1, 4},
		{4, 5, 4.59479341998814},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Pareto{c.a, c.b, nil, nil}

			res := b.Median()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestParetoMode(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{1, 1, 1},
		{2, 2, 2},
		{2, 1, 2},
		{4, 5, 4},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Pareto{c.a, c.b, nil, nil}

			res := b.Mode()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestParetoVariance(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{1, 1, math.Inf(1)},
		{2, 2, math.Inf(1)},
		{2, 1, math.Inf(1)},
		{3, 1, 0.75},
		{3, 2, 3},
		{4, 3, 2},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Pareto{c.a, c.b, nil, nil}

			res := b.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
