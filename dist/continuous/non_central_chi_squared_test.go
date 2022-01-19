package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dchisq(x, k, λ) from package stats
func TestNonCentralChiSquaredProbability(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		x        float64
		k        int
		λ        float64
		expected float64
	}{
		{1, 1, 1, 0.2264666},
		{.345344231, 3, 5, 0.02127154},
		{1.423412, 8, 4.24324, 0.00254359},
		{6.3423412, 9, 5, 0.03962636},
		{13.12313, 10, 6.12313, 0.0634349},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := NonCentralChiSquared{c.k, c.λ, nil}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R pchisq(x, k, λ) from package stats
func TestNonCentralChiSquaredDistribution(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x        float64
		k        int
		λ        float64
		expected float64
	}{
		{1, 1, 1, 0.4772499},
		{.345344231, 3, 5, 0.004717978},
		{1.423412, 8, 4.24324, 0.0009735248},
		{6.3423412, 9, 5, 0.07947634},
		{13.12313, 10, 6.12313, 0.3648576},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := NonCentralChiSquared{c.k, c.λ, nil}
			res := b.Distribution(c.x)
			run_test(t, res, c.expected, tol, "NonCentralChiSquaredDistribution")
		})
	}
}

func TestNonCentralChiSquaredInverse(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x        float64
		k        int
		λ        float64
		expected float64
	}{
		{0.4772499, 1, 1, 1},
		{0.004717978, 3, 5, .345344231},
		{0.0009735248, 8, 4.24324, 1.423412},
		{0.07947634, 9, 5, 6.3423412},
		{0.3648576, 10, 6.12313, 13.12313},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := NonCentralChiSquared{c.k, c.λ, nil}
			res := b.Inverse(c.x)
			run_test(t, res, c.expected, tol, "NonCentralChiSquaredDistribution")
		})
	}
}

func TestNonCentralChiSquaredMean(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		k        int
		λ        float64
		expected float64
	}{
		{1, 1, 2},
		{3, 5, 8},
		{8, 4.24324, 12.2432},
		{9, 5, 14},
		{10, 6.12313, 16.1231},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := NonCentralChiSquared{c.k, c.λ, nil}
			res := b.Mean()
			run_test(t, res, c.expected, tol, "NonCentralChiSquaredMean")
		})
	}
}

func TestNonCentralChiSquaredVariance(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		k        int
		λ        float64
		expected float64
	}{
		{1, 1, 6},
		{3, 5, 26},
		{8, 4.24324, 32.973},
		{9, 5, 38},
		{10, 6.12313, 44.4925},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := NonCentralChiSquared{c.k, c.λ, nil}
			res := b.Variance()
			run_test(t, res, c.expected, tol, "NonCentralChiSquaredVariance")
		})
	}
}

func TestNonCentralChiSquaredSkewness(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		k        int
		λ        float64
		expected float64
	}{
		{1, 1, 2.17732},
		{3, 5, 1.08618},
		{8, 4.24324, 0.875884},
		{9, 5, 0.819645},
		{10, 6.12313, 0.764732},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := NonCentralChiSquared{c.k, c.λ, nil}
			res := b.Skewness()
			run_test(t, res, c.expected, tol, "NonCentralChiSquaredSkewness")
		})
	}
}

func TestNonCentralChiSquaredExKurtosis(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		k        int
		λ        float64
		expected float64
	}{
		{1, 1, 6.66667},
		{3, 5, 1.63314},
		{8, 4.24324, 1.10254},
		{9, 5, 0.963989},
		{10, 6.12313, 0.836358},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := NonCentralChiSquared{c.k, c.λ, nil}
			res := b.ExKurtosis()
			run_test(t, res, c.expected, tol, "NonCentralChiSquaredExKurtosis")
		})
	}
}
