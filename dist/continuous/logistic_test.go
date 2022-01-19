package continuous

import (
	"math"
	"strconv"
	"testing"
)

func TestLogisticProbability(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, μ, s  float64
		expected float64
	}{
		{0, 0, 1, 0.25},
		{1, 0, 1, 0.1966119},
		{2, 0, 1, 0.1049936},
		{3, 0, 1, 0.04517666},
		{4, 0, 1, 0.01766271},
		{5, 0, 1, 0.006648057},
		{10, 0, 1, 4.539581e-05},

		{0, 1, 1, 0.1966119},
		{1, 1, 1, 0.25},
		{2, 1, 1, 0.1966119},
		{3, 1, 1, 0.1049936},
		{4, 1, 1, 0.04517666},
		{5, 1, 1, 0.01766271},
		{10, 1, 1, 0.0001233793},

		{-5, 0, 0.7, 0.001127488648},
		{-4.2, 0, 0.7, 0.003523584702},
		{-3.5, 0, 0.7, 0.009497223815},
		{-3.0, 0, 0.7, 0.01913226324},
		{-2.0, 0, 0.7, 0.07337619322},
		{-0.1, 0, 0.7, 0.3553268797},
		{0, 0, 0.7, 0.3571428571},
		{0.1, 0, 0.7, 0.3553268797},
		{3.5, 0, 0.7, 0.009497223815},
		{4.2, 0, 0.7, 0.003523584702},
		{5, 0, 0.7, 0.001127488648},

		{-5, 2, 1.5, 0.006152781498},
		{-3.7, 2, 1.5, 0.01426832061},
		{0, 2, 1.5, 0.1100606731},
		{3.7, 2, 1.5, 0.1228210582},
		{5, 2, 1.5, 0.06999572},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Logistic{location: c.μ, scale: c.s}

			res := e.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLogisticDistribution(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, μ, s  float64
		expected float64
	}{
		{0, 0, 1, 0.5},
		{1, 0, 1, 0.7310586},
		{2, 0, 1, 0.8807971},
		{3, 0, 1, 0.9525741},
		{4, 0, 1, 0.9820138},
		{5, 0, 1, 0.9933071},
		{10, 0, 1, 0.9999546},

		{0, 1, 1, 0.2689414},
		{1, 1, 1, 0.5},
		{2, 1, 1, 0.7310586},
		{3, 1, 1, 0.8807971},
		{4, 1, 1, 0.9525741},
		{5, 1, 1, 0.9820138},
		{10, 1, 1, 0.9998766},

		{-4.8, 0, 0.7, 0.001050809752},
		{-3.5, 0, 0.7, 0.006692850924},
		{-3.0, 0, 0.7, 0.01357691694},
		{-2.0, 0, 0.7, 0.05431326613},
		{-0.1, 0, 0.7, 0.4643463292},
		{0, 0, 0.7, 0.5},
		{0.1, 0, 0.7, 0.5356536708},
		{3.5, 0, 0.7, 0.9933071491},
		{4.2, 0, 0.7, 0.9975273768},
		{5, 0, 0.7, 0.9992101341},

		{-5, 2, 1.5, 0.009315959345},
		{-3.7, 2, 1.5, 0.02188127094},
		{0, 2, 1.5, 0.2086085273},
		{3.7, 2, 1.5, 0.7564535292},
		{5, 2, 1.5, 0.880797078},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Logistic{location: c.μ, scale: c.s}

			res := e.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLogisticInverseDistribution(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, μ, s  float64
		expected float64
	}{
		{0, 0, 1, 0.5},
		{1, 0, 1, 0.7310586},
		{2, 0, 1, 0.8807971},
		{3, 0, 1, 0.9525741},
		{4, 0, 1, 0.9820138},
		{5, 0, 1, 0.9933071},
		{10, 0, 1, 0.9999546},

		{0, 1, 1, 0.2689414},
		{1, 1, 1, 0.5},
		{2, 1, 1, 0.7310586},
		{3, 1, 1, 0.8807971},
		{4, 1, 1, 0.9525741},
		{5, 1, 1, 0.9820138},
		{10, 1, 1, 0.9998766},

		{-4.8, 0, 0.7, 0.001050809752},
		{-3.5, 0, 0.7, 0.006692850924},
		{-3.0, 0, 0.7, 0.01357691694},
		{-2.0, 0, 0.7, 0.05431326613},
		{-0.1, 0, 0.7, 0.4643463292},
		{0, 0, 0.7, 0.5},
		{0.1, 0, 0.7, 0.5356536708},
		{3.5, 0, 0.7, 0.9933071491},
		{4.2, 0, 0.7, 0.9975273768},
		{5, 0, 0.7, 0.9992101341},

		{-5, 2, 1.5, 0.009315959345},
		{-3.7, 2, 1.5, 0.02188127094},
		{0, 2, 1.5, 0.2086085273},
		{3.7, 2, 1.5, 0.7564535292},
		{5, 2, 1.5, 0.880797078},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Logistic{location: c.μ, scale: c.s}

			cdf := e.Distribution(c.x)
			inverse := e.Inverse(cdf)
			if math.Abs(inverse-c.x) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.x, inverse)
			}

		})
	}
}

func TestLogisticMean(t *testing.T) {
	i := 0
	for μ := -5; μ <= 5; μ++ {
		for s := 1; s <= 3; s++ {

			e := Logistic{location: float64(μ), scale: float64(s)}

			res := e.Mean()
			if float64(μ) != res {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, μ, res)
			}
			i++
		}
	}
}

func TestLogisticMedian(t *testing.T) {
	i := 0
	for μ := -5; μ <= 5; μ++ {
		for s := 1; s <= 3; s++ {
			e := Logistic{location: float64(μ), scale: float64(s)}

			res := e.Median()
			if float64(μ) != res {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, μ, res)
			}
			i++
		}
	}
}

func TestLogisticMode(t *testing.T) {
	i := 0
	for μ := -5; μ <= 5; μ++ {
		for s := 1; s <= 3; s++ {
			e := Logistic{location: float64(μ), scale: float64(s)}

			res := e.Mode()
			if float64(μ) != res {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, μ, res)
			}
			i++
		}
	}
}

func TestLogisticVariance(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		μ, s     float64
		expected float64
	}{
		{0, 1, 3.28986813369645},
		{0, 2, 13.15947253478581},
		{0, 3, 29.60881320326808},
		{5, 4, 52.63789013914325},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Logistic{location: c.μ, scale: c.s}

			res := e.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R (stats) qlogis(p, location, scale)
func TestLogisticInverse(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		p, μ, s  float64
		expected float64
	}{
		{0, -1, 1, math.Inf(-1)},
		{0.1, -1, 1, -3.197225},
		{0.3, -1, 1, -1.847298},
		{0.5, -1, 1, -1},
		{0.7, -1, 1, -0.1527021},
		{0.9, -1, 1, 1.197225},
		{1, -1, 1, math.Inf(1)},

		{0, 0, 1, math.Inf(-1)},
		{0.1, 0, 1, -2.197225},
		{0.3, 0, 1, -0.8472979},
		{0.5, 0, 1, 0},
		{0.7, 0, 1, 0.8472979},
		{0.9, 0, 1, 2.197225},
		{1, 0, 1, math.Inf(1)},

		{0, 1, 1, math.Inf(-1)},
		{0.1, 1, 1, -1.197225},
		{0.3, 1, 1, 0.1527021},
		{0.5, 1, 1, 1},
		{0.7, 1, 1, 1.847298},
		{0.9, 1, 1, 3.197225},
		{1, 1, 1, math.Inf(1)},

		{0, 2, 5, math.Inf(-1)},
		{0.1, 2, 5, -8.986123},
		{0.3, 2, 5, -2.236489},
		{0.5, 2, 5, 2},
		{0.7, 2, 5, 6.236489},
		{0.9, 2, 5, 12.98612},
		{1, 2, 5, math.Inf(1)},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Logistic{location: c.μ, scale: c.s}

			res := e.Inverse(c.p)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
