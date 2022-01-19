package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Test data created with calculator http://keisan.casio.com/exec/system/1180573217
// Additional data generated with R dgamma(x, shape = k, scale = θ)
func TestGammaProbability(t *testing.T) {

	tol := 0.00000001

	cases := []struct {
		x, k, θ  float64
		expected float64
	}{
		{1, 1, 1, 0.3678794411714423215955},
		{1, 2, 1, 0.3678794411714423215955},
		{1, 1, 2, 0.3032653298563167118019},
		{2, 2, 2, 0.1839397205857211607978},
		{2, 4, 1, 0.180447044315483589192},
		{4, 2, 5, 0.07189263425875545462882},
		{18, 2, 5, 0.01967308016205064377713},
		{75, 2, 5, 9.177069615054773651144e-7},
		{0.1, 0.1, 0.1, 0.386691694403023771966},
		{15, 0.1, 0.1, 8.2986014463775253874e-68},
		{4, 0.5, 6, 0.05912753695472959648351},

		{1, 4, 5, 0.0002183282},
		{2, 4, 5, 0.001430016},
		{3, 4, 5, 0.003951444},
		{5, 4, 5, 0.01226265},
		{15, 4, 5, 0.04480836},
		{115, 4, 5, 4.161876e-08},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gamma{shape: c.k, rate: 1 / c.θ}
			res := e.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Test data created with calculator http://keisan.casio.com/exec/system/1180573217
// Additional data generated with R pgamma(x, shape = k, scale = θ)
func TestGammaDistribution(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, k, θ  float64
		expected float64
	}{
		{1, 1, 1, 0.6321205588285576784045},
		{1, 2, 1, 0.264241117657115356809},
		{1, 1, 2, 0.3934693402873665763962},
		{2, 2, 2, 0.264241117657115356809},
		{2, 4, 1, 0.142876539501452951338},
		{4, 2, 5, 0.1912078645890011354258},
		{18, 2, 5, 0.8743108767424542203128},
		{75, 2, 5, 0.9999951055628719707874},
		{0.1, 0.1, 0.1, 0.975872656273672222617},
		{15, 0.1, 0.1, 1},
		{4, 0.5, 6, 0.7517869210100764165283},

		{1, 4, 5, 5.684024e-05},
		{2, 4, 5, 0.0007762514},
		{3, 4, 5, 0.003358069},
		{5, 4, 5, 0.01898816},
		{15, 4, 5, 0.3527681},
		{115, 4, 5, 0.9999998},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gamma{shape: c.k, rate: 1 / c.θ}

			res := e.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestGammaInverse(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		x, k, θ  float64
		expected float64
	}{
		{0.6321205588285576784045, 1, 1, 1},
		{0.264241117657115356809, 2, 1, 1},
		{0.3934693402873665763962, 1, 2, 1},
		{5.684024e-05, 4, 5, 1},
		{0.0007762514, 4, 5, 2},
		{0.003358069, 4, 5, 3},
		{0.01898816, 4, 5, 5},
		{0.3527681, 4, 5, 15},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gamma{shape: c.k, rate: 1 / c.θ}

			res := e.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestGammaMean(t *testing.T) {

	tol := 0.0001

	cases := []struct {
		k, θ     float64
		expected float64
	}{
		{1, 1, 1.0},
		{1, 2, 2.0},
		{2, 1, 2.0},
		{9, 0.5, 4.5},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gamma{shape: c.k, rate: 1 / c.θ}

			res := e.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestGammaMedian(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		k, θ     float64
		expected float64
	}{
		{1, 1, 0.6875},
		{1, 2, 1.375},
		{2, 1, 1.6774193548387},
		{9, 0.5, 4.33455882352943},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gamma{shape: c.k, rate: 1 / c.θ}

			res := e.Median()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestGammaMode(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		k, θ     float64
		expected float64
	}{
		{1, 1, 0},
		{1, 2, 0},
		{2, 1, 1},
		{2, 2, 2},
		{2, 3, 3},
		{3, 1, 2},
		{3, 2, 4},
		{3, 3, 6},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gamma{shape: c.k, rate: 1 / c.θ}

			res := e.Mode()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestGammaModeNaN(t *testing.T) {

	cases := []struct {
		k, θ float64
	}{
		{0.1, 1},
		{0.5, 3},
		{0.9, 6},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gamma{shape: c.k, rate: 1 / c.θ}

			res := e.Mode()
			if !math.IsNaN(res) {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, math.NaN(), res)
			}

		})
	}
}

func TestGammaVariance(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		k, θ     float64
		expected float64
	}{
		{1, 1, 1},
		{1, 2, 4},
		{2, 1, 2},
		{2, 2, 8},
		{2, 3, 18},
		{3, 1, 3},
		{3, 2, 12},
		{3, 3, 27},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gamma{shape: c.k, rate: 1 / c.θ}

			res := e.Variance()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica Skewness[GammaDistribution[k,θ]]
func TestGammaSkewness(t *testing.T) {

	tol := 0.000001

	cases := []struct {
		k, θ     float64
		expected float64
	}{
		{1, 1, 2},
		{1, 2, 2},
		{4.8776, 8.92743, 0.90558},
		{71.0984372, 45.354556, 0.237192},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gamma{shape: c.k, rate: 1 / c.θ}

			res := e.Skewness()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with Mathematica kurtosis excess[GammaDistribution[k,θ]]
func TestGammaExKurtosis(t *testing.T) {

	tol := 0.00001

	cases := []struct {
		k, θ     float64
		expected float64
	}{
		{1, 1, 6},
		{1, 2, 6},
		{4.8776, 8.92743, 1.23011},
		{71.0984372, 45.354556, 0.08439},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := Gamma{shape: c.k, rate: 1 / c.θ}

			res := e.ExKurtosis()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
