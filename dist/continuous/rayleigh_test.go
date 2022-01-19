package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R drayleigh(x, σ) from package VGAM
func TestRayleighProbability(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		x, σ, expected float64
	}{
		{10, 10, 0.06065307},
		{.4, 5, 0.01594888},
		{1.8, 2.4, 0.2358874},
		{7.8, .4, 1.311403e-81},
		{1.6721, 4.1234, 0.09058225},
		{4.12312331, 9.18237189237, 0.04421141},
		{.54241231, 6.12312345, 0.01441052},
		{1.645434342, 2.54657, 0.2059263},
		{75, 2.54657, 5.166047e-188},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Rayleigh{c.σ, nil}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R prayleigh(x, σ) from package VGAM
func TestRayleighDistribution(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		x, σ, expected float64
	}{
		{10, 10, 0.3934693},
		{.4, 5, 0.003194885},
		{1.8, 2.4, 0.2451604},
		{7.8, .4, 1},
		{1.6721, 4.1234, 0.07893176},
		{4.12312331, 9.18237189237, 0.09589715},
		{.54241231, 6.12312345, 0.003915898},
		{1.645434342, 2.54657, 0.1883993},
		{75, 2.54657, 1},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Rayleigh{c.σ, nil}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestRayleighInverse(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		x, σ, expected float64
	}{
		{0.003194885, 5, .4},
		{0.2451604, 2.4, 1.8},
		{0.07893176, 4.1234, 1.6721},
		{0.09589715, 9.18237189237, 4.12312331},
		{0.003915898, 6.12312345, .54241231},
		{0.1883993, 2.54657, 1.645434342},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Rayleigh{c.σ, nil}

			res := b.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
