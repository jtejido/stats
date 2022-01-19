package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R (stats) dlnorm(q, meanlog sdlog)
func TestLogNormalProbability(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, μ, σ  float64
		expected float64
	}{
		{4.3, 6, 2, 0.003522012},
		{4.3, 6, 1, 3.082892e-06},
		{4.3, 1, 1, 0.08351597},
		{1, 6, 2, 0.002215924},
		{2, 6, 2, 0.002951125},
		{2, 3, 2, 0.0512813},

		{0.1, -2, 1, 3.810909},
		{1, -2, 1, 0.05399097},
		{2, -2, 1, 0.005307647},
		{5, -2, 1, 0.0001182869},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogNormal{location: c.μ, scale: c.σ}

			res := e.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R (stats) plnorm(q, meanlog sdlog)
func TestLogNormalDistribution(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, μ, σ  float64
		expected float64
	}{
		{4.3, 6, 2, 0.0115828},
		{4.3, 6, 1, 2.794294e-06},
		{4.3, 1, 1, 0.6767447},
		{1, 6, 2, 0.001349898},
		{2, 6, 2, 0.003983957},
		{2, 3, 2, 0.1243677},

		{0.1, -2, 1, 0.381103},
		{1, -2, 1, 0.9772499},
		{2, -2, 1, 0.9964609},
		{5, -2, 1, 0.9998466},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogNormal{location: c.μ, scale: c.σ}

			res := e.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLogNormalMean(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		μ, σ     float64
		expected float64
	}{
		{1, 1, 4.48168907034},
		{2, 2, 54.5981500331},
		{1.3, 1.6, 13.1971381597},
		{2.6, 3.16, 1983.86055382},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogNormal{location: c.μ, scale: c.σ}

			res := e.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLogNormalMedian(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		μ, σ     float64
		expected float64
	}{
		{1, 1, 2.718281828459045},
		{2, 2, 7.38905609893065},
		{1.3, 1.6, 3.669296667619244},
		{2.6, 3.16, 13.46373803500169},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogNormal{location: c.μ, scale: c.σ}

			res := e.Median()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLogNormalMode(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		μ, σ     float64
		expected float64
	}{
		{1, 1, 1},
		{1, 2, 0.049787068367864},
		{2, 1, 2.718281828459045},
		{2, 2, 0.135335283236613},
		{1.3, 1.6, 0.28365402649977},
		{2.6, 3.16, 0.000620118480873},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogNormal{location: c.μ, scale: c.σ}

			res := e.Mode()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLogNormalVariance(t *testing.T) {

	tol := 0.001

	cases := []struct {
		μ, σ     float64
		expected float64
	}{
		{1, 1, 34.51261310995665},
		{1, 2, 21623.03700131397116},
		{2, 1, 255.01563439015922},
		{2, 2, 159773.83343196209715},
		{1.3, 1.6, 2078.79512496361378},
		{2.6, 3.16, 85446299583.51734035309427},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogNormal{location: c.μ, scale: c.σ}

			res := e.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R (stats) qlnorm(p, meanlog, sdlog)
func TestLogNormalInverse(t *testing.T) {

	tol := 0.001

	cases := []struct {
		p, μ, σ  float64
		expected float64
	}{
		{0, -1, 1, 0},
		{0.1, -1, 1, 0.1021256},
		{0.2, -1, 1, 0.1585602},
		{0.3, -1, 1, 0.2177516},
		{0.5, -1, 1, 0.3678794},
		{0.7, -1, 1, 0.6215124},
		{0.9, -1, 1, 1.325184},
		{1, -1, 1, math.Inf(1)},

		{0, 1, 1, 0},
		{0.1, 1, 1, 0.754612},
		{0.3, 1, 1, 1.608978},
		{0.5, 1, 1, 2.718282},
		{0.7, 1, 1, 4.59239},
		{0.9, 1, 1, 9.791861},
		{1, 1, 1, math.Inf(1)},

		{0, 2, 3, 0},
		{0.1, 2, 3, 0.1580799},
		{0.3, 2, 3, 1.532344},
		{0.5, 2, 3, 7.389056},
		{0.7, 2, 3, 35.63048},
		{0.9, 2, 3, 345.3833},
		{1, 2, 3, math.Inf(1)},

		{0, 5, 2, 0},
		{0.1, 5, 2, 11.43749},
		{0.3, 5, 2, 51.99767},
		{0.5, 5, 2, 148.4132},
		{0.7, 5, 2, 423.6048},
		{0.8, 5, 2, 798.9053},
		{1, 5, 2, math.Inf(1)},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogNormal{location: c.μ, scale: c.σ}

			res := e.Inverse(c.p)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestLogNormalInverseDistribution(t *testing.T) {

	tol := 0.001

	cases := []struct {
		x, μ, σ  float64
		expected float64
	}{
		{4.3, 6, 2, 0.0115828},
		{4.3, 6, 1, 2.794294e-06},
		{4.3, 1, 1, 0.6767447},
		{1, 6, 2, 0.001349898},
		{2, 6, 2, 0.003983957},
		{2, 3, 2, 0.1243677},

		{0.1, -2, 1, 0.381103},
		{1, -2, 1, 0.9772499},
		{2, -2, 1, 0.9964609},
		{5, -2, 1, 0.9998466},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := LogNormal{location: c.μ, scale: c.σ}

			cdf := e.Distribution(c.x)
			inverse := e.Inverse(cdf)
			if math.Abs(inverse-c.x) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.x, inverse)
			}

		})
	}
}
