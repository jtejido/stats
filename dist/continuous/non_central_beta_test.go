package continuous

import (
	"strconv"
	"testing"
)

func TestNonCentralBetaProbability(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		x, alpha, beta, lambda float64
		expected               float64
	}{
		{.6, 1, 2, 3, 1.40715},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := NonCentralBeta{alpha: c.alpha, beta: c.beta, lambda: c.lambda}

			res := b.Probability(c.x)
			run_test(t, res, c.expected, tol, "NonCentralBetaProbability")
		})
	}
}

func TestNonCentralBetaDistribution(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x, alpha, beta, lambda float64
		expected               float64
	}{
		{.6, 1, 2, 3, 0.579545},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := NonCentralBeta{alpha: c.alpha, beta: c.beta, lambda: c.lambda}
			res := b.Distribution(c.x)
			run_test(t, res, c.expected, tol, "NonCentralBetaDistribution")
		})
	}
}
