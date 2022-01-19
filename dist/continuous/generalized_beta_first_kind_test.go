package continuous

import (
	"math"
	"strconv"
	"testing"
)

func TestGB1Distribution(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		x, a, b, p, q float64
		expected      float64
	}{
		{1.6307179347707614, 1, 2, 3, 4, 0.98732},
		{1.7512044519651069, 3, 2, 5.2, .4, 0.034524},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := GB1{alpha: c.a, beta: c.b, p: c.p, q: c.q}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestGB1Inverse(t *testing.T) {
	tol := 0.0000001
	cases := []struct {
		x, a, b, p, q float64
		expected      float64
	}{
		{0.98732, 1, 2, 3, 4, 1.6307179347707614},
		{0.034524, 3, 2, 5.2, .4, 1.7512044519651069},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := GB1{alpha: c.a, beta: c.b, p: c.p, q: c.q}

			res := b.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
