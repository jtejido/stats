package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dsgomp(x, b, η) from package extraDistr
func TestShiftedGompertzProbability(t *testing.T) {
	tol := 0.0000001

	cases := []struct {
		x, b, η, expected float64
	}{
		{1, 2, 3, 0.648175},
		{4.1123214, 9.76564, 4.9866, 2.117646e-16},
		{.87667, 4.9831978, 7.896876, 0.5024799},
		{.00008, 20.89756, 17.09886, 8.280639e-07},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ShiftedGompertz{c.b, c.η, nil}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R psgomp(x, b, η) from package extraDistr
func TestShiftedGompertzDistribution(t *testing.T) {
	tol := 0.0000001

	cases := []struct {
		x, b, η, expected float64
	}{
		{1, 2, 3, 0.5761315},
		{4.1123214, 9.76564, 4.9866, 1},
		{.87667, 4.9831978, 7.896876, 0.8933357},
		{.00008, 20.89756, 17.09886, 6.445942e-11},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ShiftedGompertz{c.b, c.η, nil}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestShiftedGompertzInverse(t *testing.T) {
	tol := 0.0000001

	cases := []struct {
		x, b, η, expected float64
	}{
		{0.5761315, 2, 3, 1},
		{0.8933357, 4.9831978, 7.896876, .87667},
		{6.445942e-11, 20.89756, 17.09886, .00008},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ShiftedGompertz{c.b, c.η, nil}

			res := b.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
