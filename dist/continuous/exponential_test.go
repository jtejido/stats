package continuous

import (
	"math"
	"strconv"
	"testing"
)

func TestExponentialProbability(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		λ, x     float64
		expected float64
	}{

		{0.1, 0, 0.1},
		{0.1, 0.1, 0.09900498},
		{0.1, 0.5, 0.09512294},
		{0.1, 1, 0.09048374},
		{0.1, 2, 0.08187308},
		{0.1, 3, 0.07408182},
		{0.1, 10, 0.03678794},
		{0.1, 50, 0.0006737947},

		{1, 0, 1},
		{1, 0.1, 0.9048374},
		{1, 0.5, 0.6065307},
		{1, 1, 0.3678794},
		{1, 2, 0.1353353},
		{1, 3, 0.04978707},
		{1, 4, 0.01831564},

		{2, 0, 2},
		{2, 0.1, 1.637462},
		{2, 0.5, 0.7357589},
		{2, 1, 0.2706706},
		{2, 2, 0.03663128},
		{2, 3, 0.004957504},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Exponential{rate: c.λ}

			res := e.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestExponentialDistribution(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		λ, x     float64
		expected float64
	}{

		{0.1, 0, 0},
		{0.1, 0.1, 0.009950166},
		{0.1, 0.5, 0.04877058},
		{0.1, 1, 0.09516258},
		{0.1, 2, 0.1812692},
		{0.1, 3, 0.2591818},
		{0.1, 10, 0.6321206},
		{0.1, 50, 0.9932621},

		{1, 0, 0},
		{1, 0.1, 0.09516258},
		{1, 0.5, 0.3934693},
		{1, 1, 0.6321206},
		{1, 2, 0.8646647},
		{1, 3, 0.9502129},
		{1, 4, 0.9816844},

		{2, 0, 0},
		{2, 0.1, 0.1812692},
		{2, 0.5, 0.6321206},
		{2, 1, 0.8646647},
		{2, 2, 0.9816844},
		{2, 3, 0.9975212},

		{1. / 3, 2, 0.4865829},
		{1. / 3, 4, 0.7364029},
		{1. / 5, 4, 0.550671},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Exponential{rate: c.λ}

			res := e.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestExponentialMean(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		λ, μ float64
	}{
		{1, 1},

		{2, 0.5},
		{3, 0.33333},
		{4, 0.25},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Exponential{rate: c.λ}

			res := e.Mean()
			if math.Abs(res-c.μ) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.μ, res)
			}

		})
	}
}

func TestExponentialMedian(t *testing.T) {

	tol := 0.0000001

	cases := []struct {
		λ, expectedMedian float64
	}{
		{1, 0.69314718055995},
		{2, 0.3465735902799},
		{3, 0.23104906018665},
		{4, 0.17328679513999},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Exponential{rate: c.λ}

			res := e.Median()
			if math.Abs(res-c.expectedMedian) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expectedMedian, res)
			}

		})
	}
}

func TestExponentialMode(t *testing.T) {
	cases := []struct {
		λ, expectedMedian float64
	}{
		{1, 0.69314718055995},
		{2, 0.3465735902799},
		{3, 0.23104906018665},
		{4, 0.17328679513999},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Exponential{rate: c.λ}

			res := e.Mode()
			if res > 0 {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, 0, res)
			}

		})
	}
}

func TestExponentialVariance(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		λ, expectedVariance float64
	}{
		{1, 1},
		{2, 0.25},
		{3, 0.111111111111111},
		{4, 0.0625},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Exponential{rate: c.λ}

			res := e.Variance()
			if math.Abs(res-c.expectedVariance) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expectedVariance, res)
			}

		})
	}
}

func TestExponentialInverse(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		λ, p, expected float64
	}{
		{0.1, 0, 0},
		{0.1, 0.1, 1.053605},
		{0.1, 0.3, 3.566749},
		{0.1, 0.5, 6.931472},
		{0.1, 0.7, 12.03973},
		{0.1, 0.9, 23.02585},
		{1, 0, 0},
		{1, 0.1, 0.1053605},
		{1, 0.3, 0.3566749},
		{1, 0.5, 0.6931472},
		{1, 0.7, 1.203973},
		{1, 0.9, 2.302585},
		{2, 0, 0},
		{2, 0.1, 0.05268026},
		{2, 0.3, 0.1783375},
		{2, 0.5, 0.3465736},
		{2, 0.7, 0.6019864},
		{2, 0.9, 1.151293},
		{1. / 3, 0, 0},
		{1. / 3, 0.1, 0.3160815},
		{1. / 3, 0.3, 1.070025},
		{1. / 3, 0.5, 2.079442},
		{1. / 3, 0.7, 3.611918},
		{1. / 3, 0.9, 6.907755},
		{4, 0, 0},
		{4, 0.1, 0.02634013},
		{4, 0.3, 0.08916874},
		{4, 0.5, 0.1732868},
		{4, 0.7, 0.3009932},
		{4, 0.9, 0.5756463},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Exponential{rate: c.λ}

			res := e.Inverse(c.p)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestExponentialInverseDistribution(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		λ, p, expected float64
	}{
		{0.1, 0, 0},
		{0.1, 0.1, 1.053605},
		{0.1, 0.3, 3.566749},
		{0.1, 0.5, 6.931472},
		{0.1, 0.7, 12.03973},
		{0.1, 0.9, 23.02585},
		{1, 0, 0},
		{1, 0.1, 0.1053605},
		{1, 0.3, 0.3566749},
		{1, 0.5, 0.6931472},
		{1, 0.7, 1.203973},
		{1, 0.9, 2.302585},
		{2, 0, 0},
		{2, 0.1, 0.05268026},
		{2, 0.3, 0.1783375},
		{2, 0.5, 0.3465736},
		{2, 0.7, 0.6019864},
		{2, 0.9, 1.151293},
		{1. / 3, 0, 0},
		{1. / 3, 0.1, 0.3160815},
		{1. / 3, 0.3, 1.070025},
		{1. / 3, 0.5, 2.079442},
		{1. / 3, 0.7, 3.611918},
		{1. / 3, 0.9, 6.907755},
		{4, 0, 0},
		{4, 0.1, 0.02634013},
		{4, 0.3, 0.08916874},
		{4, 0.5, 0.1732868},
		{4, 0.7, 0.3009932},
		{4, 0.9, 0.5756463},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Exponential{rate: c.λ}

			cdf := e.Distribution(c.p)
			res := e.Inverse(cdf)
			if math.Abs(res-c.p) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.p, res)
			}

		})
	}
}
