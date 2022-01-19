package continuous

import (
	"math"
	"strconv"
	"testing"
)

func TestGB2Probability(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		x, a, b, p, q float64
		expected      float64
	}{
		{0.01, 2, 3, 3.2, 7.4, 1.070409e-11},
		{0.02, 2, 3, 1.43, 6.243, 0.0009695532},
		{0.03, 5.2, 1.33, .456, .243, 0.003866149},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := GB2{alpha: c.a, beta: c.b, p: c.p, q: c.q}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestGB2Distribution(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		x, a, b, p, q float64
		expected      float64
	}{
		{0.01, 2, 3, 3.2, 7.4, 1.672562e-14},
		{0.02, 2, 3, 1.43, 6.243, 6.781044e-06},
		{0.03, 5.2, 1.33, .456, .243, 4.891383e-05},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := GB2{alpha: c.a, beta: c.b, p: c.p, q: c.q}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
