package continuous

import (
	"math"
	"strconv"
	"testing"
)

func TestMaxwellBoltzmannProbability(t *testing.T) {
	tol := 0.0000001

	cases := []struct {
		x, α, expected float64
	}{
		{.875454365, 7.187531, 0.00163473},
		{1.187531, 8.2341, 0.00199464},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := MaxwellBoltzmann{scale: c.α}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestMaxwellBoltzmannDistribution(t *testing.T) {
	tol := 0.0000001

	cases := []struct {
		x, α, expected float64
	}{
		{.875454365, 7.187531, 0.000478463},
		{1.187531, 8.2341, 0.000792861},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := MaxwellBoltzmann{scale: c.α}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestMaxwellBoltzmannInverse(t *testing.T) {
	tol := 0.000001

	cases := []struct {
		x, α, expected float64
	}{
		{0.000478463, 7.187531, .875454365},
		{0.000792861, 8.2341, 1.187531},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := MaxwellBoltzmann{scale: c.α}

			res := b.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
