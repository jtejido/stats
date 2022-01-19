package continuous

import (
	"math"
	"strconv"
	"testing"
)

func TestCauchyProbability(t *testing.T) {
	tol := 0.000000001
	cases := []struct {
		x, x0, gamma float64
		expected     float64
	}{
		{1, 0, 1, 0.1591549430918953357689},
		{1, 0, 2, 0.1273239544735162686151},
		{1, 0, 3, 0.09549296585513720146133},
		{1, 0, 4, 0.07489644380795074624418},
		{1, 0, 5, 0.06121343965072897529573},
		{1, 0, 6, 0.05161781938115524403315},
		{0, 1, 1, 0.1591549430918953357689},
		{0, 1, 2, 0.1273239544735162686151},
		{0, 1, 3, 0.09549296585513720146133},
		{0, 1, 4, 0.07489644380795074624418},
		{0, 1, 5, 0.06121343965072897529573},
		{0, 1, 6, 0.05161781938115524403315},
		{1, 1, 1, 0.3183098861837906715378},
		{2, 3, 4, 0.07489644380795074624418},
		{4, 3, 2, 0.1273239544735162686151},
		{5, 5, 5, 0.06366197723675813430755},
		{-20., 7.3, 4.3, 0.001792050735277566691472},
		{-3., 7.3, 4.3, 0.01098677565090945486926},
		{-2., 7.3, 4.3, 0.01303803115441322049545},
		{-1., 7.3, 4.3, 0.01566413951236323973006},
		{0, 7.3, 4.3, 0.01906843843118277915314},
		{1, 7.3, 4.3, 0.02352582520780852333469},
		{2, 7.3, 4.3, 0.0293845536837762964279},
		{3, 7.3, 4.3, 0.03701277746323147343462},
		{20, 7.3, 4.3, 0.007613374739071642494229},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Cauchy{location: c.x0, scale: c.gamma}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestCauchyDistribution(t *testing.T) {
	tol := 0.000000001
	cases := []struct {
		x, x0, gamma float64
		expected     float64
	}{
		{1, 0, 1, 0.75},
		{1, 0, 2, 0.6475836176504332741754},
		{1, 0, 3, 0.6024163823495667258246},
		{1, 0, 4, 0.5779791303773693254605},
		{1, 0, 5, 0.5628329581890011838138},
		{1, 0, 6, 0.5525684567112534299508},
		{0, 1, 1, 0.25},
		{0, 1, 2, 0.3524163823495667258246},
		{0, 1, 3, 0.3975836176504332741754},
		{0, 1, 4, 0.4220208696226306745395},
		{0, 1, 5, 0.4371670418109988161863},
		{0, 1, 6, 0.4474315432887465700492},
		{1, 1, 1, 0.5},
		{2, 3, 4, 0.4220208696226306745395},
		{4, 3, 2, 0.6475836176504332741754},
		{5, 5, 5, 0.5},
		{-20., 7.3, 4.3, 0.04972817023155424541129},
		{-3., 7.3, 4.3, 0.1258852891111436766445},
		{-2., 7.3, 4.3, 0.1378566499474175095298},
		{-1., 7.3, 4.3, 0.1521523453170898354801},
		{0, 7.3, 4.3, 0.1694435179635968563959},
		{1, 7.3, 4.3, 0.1906393755555404651458},
		{2, 7.3, 4.3, 0.2169618719223694455636},
		{3, 7.3, 4.3, 0.25},
		{20, 7.3, 4.3, 0.8960821670587991836005},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Cauchy{location: c.x0, scale: c.gamma}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestCauchyInverseDistribution(t *testing.T) {
	tol := 0.000000001
	cases := []struct {
		x, x0, gamma float64
	}{
		{1, 0, 1},
		{1, 0, 2},
		{1, 0, 3},
		{1, 0, 4},
		{1, 0, 5},
		{1, 0, 6},
		{0, 1, 1},
		{0, 1, 2},
		{0, 1, 3},
		{0, 1, 4},
		{0, 1, 5},
		{0, 1, 6},
		{1, 1, 1},
		{2, 3, 4},
		{4, 3, 2},
		{5, 5, 5},
		{-20., 7.3, 4.3},
		{-3., 7.3, 4.3},
		{-2., 7.3, 4.3},
		{-1., 7.3, 4.3},
		{0, 7.3, 4.3},
		{1, 7.3, 4.3},
		{2, 7.3, 4.3},
		{3, 7.3, 4.3},
		{20, 7.3, 4.3},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Cauchy{location: c.x0, scale: c.gamma}

			cdf := b.Distribution(c.x)
			inv := b.Inverse(cdf)

			if math.Abs(c.x-inv) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.x, inv)
			}

		})
	}
}

func TestCauchyInverse(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		p, x0, gamma float64
		expected     float64
	}{
		{0.1, 1, 1, -2.077684},
		{0.3, 1, 1, 0.2734575},
		{0.5, 1, 1, 1},
		{0.7, 1, 1, 1.726543},
		{0.9, 1, 1, 4.077684},

		{0.1, 2, 3, -7.233051},
		{0.3, 2, 3, -0.1796276},
		{0.5, 2, 3, 2},
		{0.7, 2, 3, 4.179628},
		{0.9, 2, 3, 11.23305},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Cauchy{location: c.x0, scale: c.gamma}

			res := b.Inverse(c.p)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestCauchyMedian(t *testing.T) {
	cases := []struct {
		x0, gamma float64
	}{
		{-1, 0.1},
		{-1, 1},
		{-1, 2},
		{0, 0.1},
		{0, 1},
		{0, 2},
		{1, 0.1},
		{1, 1},
		{1, 2},
		{2, 3},
		{5, 3},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Cauchy{location: c.x0, scale: c.gamma}

			res := b.Median()
			if c.x0 != res {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.x0, res)
			}

		})
	}
}

func TestCauchyMode(t *testing.T) {
	cases := []struct {
		x0, gamma float64
	}{
		{-1, 0.1},
		{-1, 1},
		{-1, 2},
		{0, 0.1},
		{0, 1},
		{0, 2},
		{1, 0.1},
		{1, 1},
		{1, 2},
		{2, 3},
		{5, 3},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Cauchy{location: c.x0, scale: c.gamma}

			res := b.Mode()
			if c.x0 != res {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.x0, res)
			}

		})
	}
}
