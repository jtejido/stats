package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dlevy(y,m,s) from package rmutil
func TestLevyProbability(t *testing.T) {

	tol := 0.0000001

	cases := []struct {
		x, μ, c  float64
		expected float64
	}{
		{1, 0, 1, 0.2419707},
		{1.1, 0, 1, 0.2194899},
		{1.2, 0, 1, 0.2000701},
		{5, 0, 1, 0.03228685},
		{7, 6, 5, 0.07322491},
		{4, 3.123123, 5.453, 0.05063548},
		{43535, 1.23424, 6.234, 1.096536e-07},
		{6, .23424, .653243, 0.02200698},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Levy{location: c.μ, scale: c.c}

			res := e.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R plevy(y,m,s) from package rmutil
func TestLevyDistribution(t *testing.T) {

	tol := 0.0000001

	cases := []struct {
		x, μ, c  float64
		expected float64
	}{
		{1, 0, 1, 0.3173105},
		{1.1, 0, 1, 00.3403557},
		{1.2, 0, 1, 0.3613104},
		{5, 0, 1, 0.6547208},
		{7, 6, 5, 0.02534732},
		{4, 3.123123, 5.453, 0.01264107},
		{43535, 1.23424, 6.234, 0.9904523},
		{6, .23424, .653243, 0.7364214},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Levy{location: c.μ, scale: c.c}

			res := e.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLevyMedian(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		μ, c     float64
		expected float64
	}{
		{0, 1, 2.19811},
		{3.123123, 5.453, 15.1094},
		{1.23424, 6.234, 14.9373},
		{.23424, .653243, 1.67014},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Levy{location: c.μ, scale: c.c}

			res := e.Median()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLevyInverse(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		x, μ, c  float64
		expected float64
	}{
		{0.3173105, 0, 1, 1},
		{00.3403557, 0, 1, 1.1},
		{0.3613104, 0, 1, 1.2},
		{0.6547208, 0, 1, 5},
		{0.02534732, 6, 5, 7},
		{0.01264107, 3.123123, 5.453, 4},
		{0.7364214, .23424, .653243, 6},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Levy{location: c.μ, scale: c.c}

			res := e.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
