package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dbenini(x, σ, β) from package VGAM
// Take note that this is 2-parameter where α=0
func TestBeniniProbability(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		x, α, β, σ float64
		expected   float64
	}{
		{4, 0, 1, 2, 0.2143569},
		{5, 0, 1, 4, 0.08492186},
		{4, 0, 10, 3, 0.6287188},
		{7, 0, 21, 5, 0.1873147},
		{4, 0, 5, 3, 0.4754881},
		{8.246562, 4.6623, 10.6573, 7.4347, 0.458359},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Benini{alpha: c.α, beta: c.β, sigma: c.σ}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R pbenini(x, σ, β) from package VGAM
// Take note that this is 2-parameter where α=0
func TestBeniniDistribution(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		x, α, β, σ float64
		expected   float64
	}{
		{4, 0, 1, 2, 0.3814969},
		{5, 0, 1, 4, 0.04857369},
		{4, 0, 10, 3, 0.5629072},
		{7, 0, 21, 5, 0.9072164},
		{4, 0, 5, 3, 0.3388701},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Benini{alpha: c.α, beta: c.β, sigma: c.σ}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestBeniniInverse(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x, α, β, σ float64
		expected   float64
	}{
		{0.3814969, 0, 1, 2, 4},
		{0.04857369, 0, 1, 4, 5},
		{0.5629072, 0, 10, 3, 4},
		{0.9072164, 0, 21, 5, 7},
		{0.3388701, 0, 5, 3, 4},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Benini{alpha: c.α, beta: c.β, sigma: c.σ}

			res := b.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Variance[BeniniDistribution[a, b]]
func TestBeniniVariance(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		α, β, σ  float64
		expected float64
	}{
		{0, 1, 2, 9.69602},
		{0, 1, 4, 38.7841},
		{0, 10, 3, 0.383294},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Benini{alpha: c.α, beta: c.β, sigma: c.σ}

			res := b.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Mean[BeniniDistribution[a, b]]
func TestBeniniMean(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		α, β, σ  float64
		expected float64
	}{
		{0, 1, 2, 5.46047},
		{0, 1, 4, 10.9209},
		{0, 10, 3, 4.01456},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Benini{alpha: c.α, beta: c.β, sigma: c.σ}

			res := b.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Median[BeniniDistribution[a, b]]
func TestBeniniMedian(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		α, β, σ  float64
		expected float64
	}{
		{0, 1, 2, 4.59837},
		{0, 1, 4, 9.19674},
		{0, 10, 3, 3.90356},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Benini{alpha: c.α, beta: c.β, sigma: c.σ}

			res := b.Median()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
