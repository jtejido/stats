package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dgompertz(x, η, b) from package extraDistr
// extra data from https://www.saecanet.com/20100716/saecanet_calculation_page.php?pagenumber=664
func TestGompertzProbability(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		x, η, b  float64 // scale, shape
		expected float64
	}{
		{1, 1, 1, 0.4875893},
		{1, 1.61, 2.1, 0.054050881792843},
		{.5, 4.1231, 7.987765, 2.565336e-10},
		{.1231, 3.87665, 3.98778, 3.420375},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gompertz{shape: c.η, scale: c.b}

			res := e.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R pgompertz(x, η, b) from package extraDistr
// extra data from https://www.saecanet.com/20100716/saecanet_calculation_page.php?pagenumber=664
func TestGompertzDistribution(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		x, η, b  float64
		expected float64
	}{
		{1, 1, 1, 0.8206259},
		{1, 1.61, 2.1, 0.99588889569672},
		{.5, 4.1231, 7.987765, 1},
		{.1231, 3.87665, 3.98778, 0.4599627},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gompertz{shape: c.η, scale: c.b}

			res := e.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
