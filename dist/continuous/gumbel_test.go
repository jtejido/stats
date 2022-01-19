package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dgumbel(x, μ, β) from package VGAM
func TestGumbelProbability(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, μ, β  float64
		expected float64
	}{
		{1, 2, 3, 0.1152224},
		{15.97887, 4.098786, 10.8786, 0.02205133},
		{100.987785, 110.987674, 115.978785785, 0.003159841},
		{-50.86876, -75.98789, 1.9876876, 1.634231e-06},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gumbel{location: c.μ, scale: c.β}

			res := e.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R pgumbel(x, μ, β) from package VGAM
func TestGumbelDistribution(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, μ, β  float64
		expected float64
	}{
		{1, 2, 3, 0.2476813},
		{15.97887, 4.098786, 10.8786, 0.7149629},
		{100.987785, 110.987674, 115.978785785, 0.3362004},
		{-50.86876, -75.98789, 1.9876876, 0.9999968},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gumbel{location: c.μ, scale: c.β}

			res := e.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestGumbelMean(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		μ, β     float64
		expected float64
	}{
		{2, 3, 3.73165},
		{4.098786, 10.8786, 10.3781},
		{-75.98789, 1.9876876, -74.8406},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gumbel{location: c.μ, scale: c.β}

			res := e.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestGumbelMode(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		μ, β     float64
		expected float64
	}{
		{2, 3, 2},
		{4.098786, 10.8786, 4.09879},
		{-75.98789, 1.9876876, -75.9879},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gumbel{location: c.μ, scale: c.β}

			res := e.Mode()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestGumbelVariance(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		μ, β     float64
		expected float64
	}{
		{2, 3, 14.8044},
		{4.098786, 10.8786, 194.668},
		{-75.98789, 1.9876876, 6.49897},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gumbel{location: c.μ, scale: c.β}

			res := e.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestGumbelSkewness(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		μ, β     float64
		expected float64
	}{
		{2, 3, 1.13955},
		{4.098786, 10.8786, 1.13955},
		{-75.98789, 1.9876876, 1.13955},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gumbel{location: c.μ, scale: c.β}

			res := e.Skewness()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
