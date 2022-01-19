package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dburr7(x, k, c) from package VaRES
func TestBurrProbability(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		x, k, c  float64
		expected float64
	}{
		{1, 2, 3, 0.75},
		{5, 1.4, 2.5, 0.002400062},
		{1.3, 5.9876823745, 8.765131987, 2.167489e-05},
		{7.096324, 4.123154, 1.73982374, 6.720252e-07},
		{.987, 2.987785, 4.98745646, 1.013249},
		{1, .298743, 1.987765, 0.2413805},
		{10.9876786, 2.98786, .9867654654, 0.0001600443},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Burr{c: c.c, k: c.k, scale: 1}

			res := b.Probability(c.x)

			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R pburr7(x, k, c) from package VaRES
func TestBurrDistribution(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		x, k, c  float64
		expected float64
	}{
		{1, 2, 3, 0.75},
		{5, 1.4, 2.5, 0.99651},
		{1.3, 5.9876823745, 8.765131987, 0.9999994},
		{7.096324, 4.123154, 1.73982374, 0.9999993},
		{.987, 2.987785, 4.98745646, 0.8612485},
		{1, .298743, 1.987765, 0.1870396},
		{10.9876786, 2.98786, .9867654654, 0.9993475},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Burr{c: c.c, k: c.k, scale: 1}

			res := b.Distribution(c.x)

			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestBurrInverse(t *testing.T) {
	tol := 0.001
	cases := []struct {
		x, k, c  float64
		expected float64
	}{
		{0.75, 2, 3, 1},
		{0.99651, 1.4, 2.5, 5},
		{0.9999994, 5.9876823745, 8.765131987, 1.3},
		// {0.9999993, 4.123154, 1.73982374, 7.096324},
		{0.8612485, 2.987785, 4.98745646, .987},
		{0.1870396, .298743, 1.987765, 1},
		{0.9993475, 2.98786, .9867654654, 10.9876786},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Burr{c: c.c, k: c.k, scale: 1}

			res := b.Inverse(c.x)

			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Mean[SinghMaddalaDistribution[q,a,1]]
func TestBurrMean(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		k, c     float64
		expected float64
	}{
		{2, 3, 0.806133},
		{1.4, 2.5, 1},
		{5.9876823745, 8.765131987, 0.779655},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Burr{c: c.c, k: c.k, scale: 1}

			res := b.Mean()

			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Variance[SinghMaddalaDistribution[q,a,1]]
func TestBurrVariance(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		k, c     float64
		expected float64
	}{
		{2, 3, 0.156283},
		{1.4, 2.5, 0.563244},
		{5.9876823745, 8.765131987, 0.0127741},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Burr{c: c.c, k: c.k, scale: 1}

			res := b.Variance()

			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Median[SinghMaddalaDistribution[q,a,1]]
func TestBurrMedian(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		k, c     float64
		expected float64
	}{
		{2, 3, 0.745432},
		{1.4, 2.5, 0.836862},
		{5.9876823745, 8.765131987, 0.787153},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Burr{c: c.c, k: c.k, scale: 1}

			res := b.Median()

			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Skewness[SinghMaddalaDistribution[q,a,1]]
func TestBurrSkewness(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		k, c     float64
		expected float64
	}{
		{2, 3, 1.58913},
		{1.4, 2.5, 7.12346},
		{5.9876823745, 8.765131987, -0.370531},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Burr{c: c.c, k: c.k, scale: 1}

			res := b.Skewness()

			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica kurtosis excess[SinghMaddalaDistribution[q,a,1]]
func TestBurrExKurtosis(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		k, c     float64
		expected float64
	}{
		{2, 3, 7.80945},
		{1.4, 2.5, math.NaN()},
		{5.9876823745, 8.765131987, 0.310003},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := Burr{c: c.c, k: c.k, scale: 1}

			res := b.ExKurtosis()

			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
