package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dnaka(x, Ω, m) from package VGAM
func TestNakagamiProbability(t *testing.T) {

	tol := 0.0000001

	cases := []struct {
		x, Ω, m  float64
		expected float64
	}{
		{1, 2, 1, 0.6065307},
		{4, 2, 3, 1.304686e-07},
		{1.2, .1, 5.9, 5.024276e-28},
		{4.9078671, 3.8162378163289, 10.0001, 4.366385e-16},
		{9.98789, 10.918273, 4.97675, 2.745714e-14},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Nakagami{shape: c.m, spread: c.Ω}

			res := e.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R pnaka(x, Ω, m) from package VGAM
func TestNakagamiDistribution(t *testing.T) {

	tol := 0.0000001

	cases := []struct {
		x, Ω, m  float64
		expected float64
	}{
		{1, 2, 1, 0.3934693},
		{4, 2, 3, 1},
		{1.2, .1, 5.9, 1},
		{2.9078671, 3.8162378163289, 10.0001, 0.9986344},
		{1.98789, 10.918273, 4.97675, 0.03750934},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Nakagami{shape: c.m, spread: c.Ω}

			res := e.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestNakagamiInverse(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		x, Ω, m  float64
		expected float64
	}{
		{0.3934693, 2, 1, 1},
		{1, 2, 3, math.Inf(1)},
		{1, .1, 5.9, math.Inf(1)},
		{0.9986344, 3.8162378163289, 10.0001, 2.9078671},
		{0.03750934, 10.918273, 4.97675, 1.98789},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Nakagami{shape: c.m, spread: c.Ω}

			res := e.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestNakagamiMean(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		m, Ω     float64
		expected float64
	}{
		{1, 2, 1.25331},
		{3, 2, 1.35675},
		{5.9, .1, 0.309606},
		{10.0001, 3.8162378163289, 1.92926},
		{4.97675, 10.918273, 3.22246},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Nakagami{shape: c.m, spread: c.Ω}

			res := e.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestNakagamiVariance(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		m, Ω     float64
		expected float64
	}{
		{1, 2, 0.429204},
		{3, 2, 0.159223},
		{5.9, .1, 0.00414395},
		{10.0001, 3.8162378163289, 0.0941837},
		{4.97675, 10.918273, 0.534049},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Nakagami{shape: c.m, spread: c.Ω}

			res := e.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestNakagamiSkewness(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		m, Ω     float64
		expected float64
	}{
		{1, 2, 0.631111},
		{3, 2, 0.317911},
		{5.9, .1, 0.216661},
		{10.0001, 3.8162378163289, 0.163039},
		{4.97675, 10.918273, 0.238046},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Nakagami{shape: c.m, spread: c.Ω}

			res := e.Skewness()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestNakagamiExKurtosis(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		m, Ω     float64
		expected float64
	}{
		{1, 2, 0.245089},
		{3, 2, 0.0251113},
		{5.9, .1, 0.00601941},
		{10.0001, 3.8162378163289, 0.00201003},
		{4.97675, 10.918273, 0.00860426},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Nakagami{shape: c.m, spread: c.Ω}

			res := e.ExKurtosis()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
