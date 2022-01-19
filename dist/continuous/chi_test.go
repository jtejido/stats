package continuous

import (
	"math"
	"strconv"
	"testing"
)

func TestChiProbability(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x        float64
		k        int
		expected float64
	}{
		{0, 1, 0},
		{1, 1, 0.483941},
		{2, 1, 0.107982},
		{3, 1, 0.0088637},
		{0, 2, 0},
		{1, 2, 0.606531},
		{2, 2, 0.270671},
		{3, 2, 0.033327},
		{0, 3, 0},
		{1, 3, 0.483941},
		{2, 3, 0.431928},
		{3, 3, 0.0797733},
		{1, 4, 0.303265},
		{1, 5, 0.161314},
		{2, 3, 0.431928},
		{2.5, 1, 0.0350566},
		{2.5, 10, 0.436474},
		{2.5, 15, 0.0966411},
		{6, 5, 5.24956e-6},
		{6, 20, 0.0499505},
		{17.66, 6, 4.06385e-63},
		{0.09, 6, 7.35129e-7},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := &Chi{dof: c.k}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestChiDistribution(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x        float64
		k        int
		expected float64
	}{
		{0, 1, 0},
		{1, 1, 0.682689},
		{2, 1, 0.9545},
		{3, 1, 0.9973},
		{0, 2, 0},
		{1, 2, 0.393469},
		{2, 2, 0.864665},
		{3, 2, 0.988891},
		{0, 3, 0},
		{1, 3, 0.198748},
		{2, 3, 0.738536},
		{3, 3, 0.970709},
		{1, 4, 0.090204},
		{1, 5, 0.0374342},
		{2, 3, 0.738536},
		{2.5, 1, 0.987581},
		{2.5, 10, 0.20616},
		{2.5, 15, 0.0247646},
		{6, 5, 0.999999},
		{6, 20, 0.984619},
		{17.66, 6, 1},
		{0.09, 6, 1.10381e-8},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Chi{dof: c.k}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestChiInverse(t *testing.T) {
	tol := 0.001
	cases := []struct {
		p        float64
		k        int
		expected float64
	}{
		{0, 1, 0},
		{0.682689, 1, 1},
		{0.9545, 1, 2},
		{0.9973, 1, 3},
		{0, 2, 0},
		{0.393469, 2, 1},
		{0.864665, 2, 2},
		{0.988891, 2, 3},
		{0, 3, 0},
		{0.198748, 3, 1},
		{0.738536, 3, 2},
		{0.970709, 3, 3},
		{0.090204, 4, 1},
		{0.0374342, 5, 1},
		{0.738536, 3, 2},
		{0.987581, 1, 2.5},
		{0.20616, 10, 2.5},
		{0.0247646, 15, 2.5},
		{0.984619, 20, 6},
		{1.10381e-8, 6, 0.09},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Chi{dof: c.k}

			res := b.Inverse(c.p)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
