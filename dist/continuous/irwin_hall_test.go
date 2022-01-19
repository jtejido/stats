package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dirwin.hall(x, n) from package unifed
func TestIrwinHallProbability(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x        float64
		n        uint
		expected float64
	}{
		{1, 5, 0.04166667},
		{.4231, 4, 0.01262344},
		{.023234, 2, 0.023234},
		{13, 26, 0.2694598},
		{5.2312, 6, 0.002238132},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := IrwinHall{n: c.n}

			res := e.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestIrwinHallDistribution(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x        float64
		n        uint
		expected float64
	}{
		{1, 5, 0.00833333},
		{.4231, 4, 0.00133524},
		{.023234, 2, 0.000269909},
		{13, 26, 0.5},
		{5.2312, 6, 0.999713},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := IrwinHall{n: c.n}

			res := e.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
