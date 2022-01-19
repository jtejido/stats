package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R drice(x, σ, v) from package VGAM
func TestRiceProbability(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		x, σ, v, expected float64
	}{
		{1, 1, 1, 0.4657596},
		{5, 4, 2, 0.1388959},
		{.5, 1, .05, 0.4407661},
		{1.3, 3.0098, 2.12134, 0.10435},
		{.998, 11.09653, 3.7867821, 0.00761752},
		{0.000009, .900, .5, 9.522188e-06},
		{100, 20, 10.5, 2.940822e-06},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Rice{c.v, c.σ, nil}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R price(x, σ, v) from package VGAM
func TestRiceDistribution(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		x, σ, v, expected float64
	}{
		{1, 1, 1, 0.2671202},
		{5, 4, 2, 0.4991223},
		{.5, 1, .05, 0.1173653},
		{1.3, 3.0098, 2.12134, 0.07026799},
		{.998, 11.09653, 3.7867821, 0.003808391},
		{0.000009, .900, .5, 4.284984e-11},
		{100, 20, 10.5, 0.9999872},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Rice{c.v, c.σ, nil}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestRiceInverse(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		x, σ, v, expected float64
	}{
		{0.2671202, 1, 1, 1},
		{0.4991223, 4, 2, 5},
		{0.1173653, 1, .05, .5},
		{0.07026799, 3.0098, 2.12134, 1.3},
		{0.003808391, 11.09653, 3.7867821, .998},
		{4.284984e-11, .900, .5, 0.000009},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Rice{c.v, c.σ, nil}

			res := b.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
