package continuous

import (
	"math"
	"strconv"
	"testing"
)

func TestLaplaceProbability(t *testing.T) {

	tol := 0.00000001

	cases := []struct {
		x, μ, b  float64
		expected float64
	}{
		{1, 0, 1, 0.1839397206},
		{1.1, 0, 1, 0.1664355418},
		{1.2, 0, 1, 0.150597106},
		{5, 0, 1, 0.0033689735},
		{1, 2, 1.4, 0.174836307},
		{1.1, 2, 1.4, 0.1877814373},
		{2.9, 2, 1.4, 0.1877814373},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Laplace{location: c.μ, scale: c.b}

			res := e.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLaplaceDistribution(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, μ, b  float64
		expected float64
	}{
		{1, 0, 1, 0.8160602794},
		{1.1, 0, 1, 0.8335644582},
		{1.2, 0, 1, 0.849402894},
		{5, 0, 1, 0.9966310265},
		{1, 2, 1.4, 0.2447708298},
		{1.1, 2, 1.4, 0.2628940122},
		{2.9, 2, 1.4, 0.7371059878},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Laplace{location: c.μ, scale: c.b}

			res := e.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLaplaceMean(t *testing.T) {
	i := 0
	for μ := -5; μ <= 5; μ++ {
		for b := 1; b <= 3; b++ {

			e := Laplace{location: float64(μ), scale: float64(b)}

			res := e.Mean()
			if float64(μ) != res {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, μ, res)
			}

			i++
		}
	}
}

func TestLaplaceMedian(t *testing.T) {
	i := 0
	for μ := -5; μ <= 5; μ++ {
		for b := 1; b <= 3; b++ {

			e := Laplace{location: float64(μ), scale: float64(b)}
			res := e.Median()
			if float64(μ) != res {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, μ, res)
			}
			i++
		}
	}
}

func TestLaplaceMode(t *testing.T) {
	i := 0
	for μ := -5; μ <= 5; μ++ {
		for b := 1; b <= 3; b++ {
			e := Laplace{location: float64(μ), scale: float64(b)}

			res := e.Mode()
			if float64(μ) != res {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, μ, res)
			}
			i++
		}
	}
}

func TestLaplaceVariance(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		μ, b     float64
		expected float64
	}{
		{1, 1, 2},
		{2, 1, 2},
		{3, 1, 2},
		{1, 2, 8},
		{2, 2, 8},
		{4, 3, 18},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Laplace{location: c.μ, scale: c.b}

			res := e.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R (rmutil) qlaplace(p, location, dispersion)
func TestLaplaceInverse(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		p, μ, b  float64
		expected float64
	}{
		{0, 0, 1, math.Inf(-1)},
		{0.1, 0, 1, -1.609438},
		{0.3, 0, 1, -0.5108256},
		{0.5, 0, 1, 0},
		{0.7, 0, 1, 0.5108256},
		{0.9, 0, 1, 1.609438},
		{1, 0, 1, math.Inf(1)},

		{0, 1, 1, math.Inf(-1)},
		{0.1, 1, 1, -0.6094379},
		{0.2, 1, 1, 0.08370927},
		{0.3, 1, 1, 0.4891744},
		{0.5, 1, 1, 1},
		{0.7, 1, 1, 1.510826},
		{0.9, 1, 1, 2.609438},
		{1, 1, 1, math.Inf(1)},

		{0, -1, 1, math.Inf(-1)},
		{0.1, -1, 1, -2.609438},
		{0.3, -1, 1, -1.510826},
		{0.5, -1, 1, -1},
		{0.7, -1, 1, -0.4891744},
		{0.9, -1, 1, 0.6094379},
		{1, -1, 1, math.Inf(1)},

		{0, 2, 4, math.Inf(-1)},
		{0.1, 2, 4, -4.437752},
		{0.3, 2, 4, -0.0433025},
		{0.5, 2, 4, 2},
		{0.7, 2, 4, 4.043302},
		{0.9, 2, 4, 8.437752},
		{1, 2, 4, math.Inf(1)},

		{0, 13, 9, math.Inf(-1)},
		{0.1, 13, 9, -1.484941},
		{0.3, 13, 9, 8.402569},
		{0.5, 13, 9, 13},
		{0.7, 13, 9, 17.59743},
		{0.9, 13, 9, 27.48494},
		{1, 13, 9, math.Inf(1)},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Laplace{location: c.μ, scale: c.b}

			res := e.Inverse(c.p)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
