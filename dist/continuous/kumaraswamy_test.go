package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dkumar(x, a, b) from package extraDistr
func TestKumaraswamyProbability(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, a, b  float64
		expected float64
	}{
		{.31, 3.342, 4.32, 0.8693421},
		{.0012, 6.1231, 7.123, 4.742286e-14},
		{.2101435, 1.1231, 5.123, 2.165181},
		{.9234, 5.123, 10.45645, 0.001250973},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Kumaraswamy{a: c.a, b: c.b}

			res := e.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R pkumar(x, a, b) from package extraDistr
func TestKumaraswamyDistribution(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, a, b  float64
		expected float64
	}{
		{.31, 3.342, 4.32, 0.08340847},
		{.2101435, 1.1231, 5.123, 0.6230961},
		{.9234, 5.123, 10.45645, 0.9999891},
		{.762461, 15.123, 50.45645, 0.5691524760806057},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Kumaraswamy{a: c.a, b: c.b}

			res := e.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestKumaraswamyInverse(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		x, a, b  float64
		expected float64
	}{
		{0.08340847, 3.342, 4.32, .31},
		{0.6230961, 1.1231, 5.123, .2101435},
		{0.9999891, 5.123, 10.45645, .9234},
		{0.5691524760806057, 15.123, 50.45645, .762461},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Kumaraswamy{a: c.a, b: c.b}

			res := e.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestKumaraswamyMean(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		a, b     float64
		expected float64
	}{
		{3.342, 4.32, 0.555296},
		{1.1231, 5.123, 0.192431},
		{5.123, 10.45645, 0.575189},
		{15.123, 50.45645, 0.744785},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Kumaraswamy{a: c.a, b: c.b}

			res := e.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestKumaraswamyMedian(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		a, b     float64
		expected float64
	}{
		{3.342, 4.32, 0.564852},
		{1.1231, 5.123, 0.158727},
		{5.123, 10.45645, 0.584997},
		{15.123, 50.45645, 0.752789},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Kumaraswamy{a: c.a, b: c.b}

			res := e.Median()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestKumaraswamyVariance(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		a, b     float64
		expected float64
	}{
		{3.342, 4.32, 0.0276431},
		{1.1231, 5.123, 0.0218351},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Kumaraswamy{a: c.a, b: c.b}

			res := e.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestKumaraswamySkewness(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		a, b     float64
		expected float64
	}{
		{3.342, 4.32, -0.247264},
		{1.1231, 5.123, 1.0061},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Kumaraswamy{a: c.a, b: c.b}

			res := e.Skewness()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestKumaraswamyExKurtosis(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		a, b     float64
		expected float64
	}{
		{3.342, 4.32, -0.443862},
		{1.1231, 5.123, 0.685971},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Kumaraswamy{a: c.a, b: c.b}

			res := e.ExKurtosis()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
