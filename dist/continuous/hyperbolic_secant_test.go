package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dsecant(x) from package VaRES
func TestHyperbolicSecantProbability(t *testing.T) {

	tol := 0.00000001

	cases := []struct {
		x        float64
		expected float64
	}{
		{1, 0.1992684},
		{-5, 0.0003882031},
		{-10.11, 1.267878e-07},
		{2.00000001, 0.04313337},
		{-8.5431, 1.485929e-06},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var hs HyperbolicSecant
			res := hs.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R psecant(x) from package VaRES
func TestHyperbolicSecantDistribution(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x        float64
		expected float64
	}{
		{1, 0.8695181},
		{-5, 0.0002471378},
		{-10.11, 8.071563e-08},
		{2.00000001, 0.9725063},
		{-8.5431, 9.459721e-07},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var hs HyperbolicSecant
			res := hs.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestHyperbolicSecantInverse(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x        float64
		expected float64
	}{
		{0.8695181, 1},
		{0.0002471378, -5},
		{8.071563e-08, -10.11},
		{0.9725063, 2.00000001},
		{9.459721e-07, -8.5431},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var hs HyperbolicSecant
			res := hs.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
