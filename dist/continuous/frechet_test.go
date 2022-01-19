package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dfrechet(x, location = 0, scale = 1, shape) from package VGAM
func TestFrechetProbability(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		x, m, s, α float64 // location , scale , shape
		expected   float64
	}{
		{4, 2, 3, 1, 0.1673476},
		{6, 3, 2, 5, 0.1923984},
		{200, 50, 100, 10, 0.001136226},
		{70, -10, 80, 90, 0.4138644},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Frechet{shape: c.α, scale: c.s, location: c.m}

			res := e.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R pfrechet(x, location = 0, scale = 1, shape) from package VGAM
func TestFrechetDistribution(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		x, m, s, α float64 // location , scale , shape
		expected   float64
	}{
		{4, 2, 3, 1, 0.2231302},
		{6, 3, 2, 5, 0.8766151},
		{200, 50, 100, 10, 0.982808},
		{70, -10, 80, 90, 0.3678794},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Frechet{shape: c.α, scale: c.s, location: c.m}

			res := e.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestFrechetInverse(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		x, m, s, α float64 // location , scale , shape
		expected   float64
	}{
		{0.2231302, 2, 3, 1, 4},
		{0.8766151, 3, 2, 5, 6},
		{0.982808, 50, 100, 10, 200},
		{0.3678794, -10, 80, 90, 70},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Frechet{shape: c.α, scale: c.s, location: c.m}

			res := e.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestFrechetMean(t *testing.T) {
	tol := 0.001
	cases := []struct {
		α, s, m  float64 // location , scale , shape
		expected float64
	}{
		{10, 100, 50, 156.863},
		{90, 80, -10, 70.5229},
		{5.123764, 4.65443, -6.4578, -1.06403},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Frechet{shape: c.α, scale: c.s, location: c.m}

			res := e.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestFrechetMedian(t *testing.T) {
	tol := 0.001
	cases := []struct {
		α, s, m  float64 // location , scale , shape
		expected float64
	}{
		{10, 100, 50, 153.733},
		{90, 80, -10, 70.3265},
		{5.123764, 4.65443, -6.4578, -1.45823},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Frechet{shape: c.α, scale: c.s, location: c.m}

			res := e.Median()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestFrechetVariance(t *testing.T) {
	tol := 0.001
	cases := []struct {
		α, s, m  float64 // location , scale , shape
		expected float64
	}{
		{10, 100, 50, 222.624},
		{90, 80, -10, 1.33865},
		{5.123764, 4.65443, -6.4578, 2.6974},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Frechet{shape: c.α, scale: c.s, location: c.m}

			res := e.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestFrechetSkewness(t *testing.T) {
	tol := 0.001
	cases := []struct {
		α, s, m  float64 // location , scale , shape
		expected float64
	}{
		{10, 100, 50, 1.91034},
		{90, 80, -10, 1.20741},
		{5.123764, 4.65443, -6.4578, 3.4098},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Frechet{shape: c.α, scale: c.s, location: c.m}

			res := e.Skewness()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestFrechetExKurtosis(t *testing.T) {
	tol := 0.001
	cases := []struct {
		α, s, m  float64 // location , scale , shape
		expected float64
	}{
		{10, 100, 50, 7.97857},
		{90, 80, -10, 2.74111},
		{5.123764, 4.65443, -6.4578, 39.7567},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Frechet{shape: c.α, scale: c.s, location: c.m}

			res := e.ExKurtosis()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
