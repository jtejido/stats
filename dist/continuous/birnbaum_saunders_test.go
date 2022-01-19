package continuous

// import (
// 	"math"
// 	"strconv"
// 	"testing"
// )

// // Generated with R dfatigue(x, γ, β, μ) from package extraDistr
// func TestBirnbaumSaundersProbability(t *testing.T) {
// 	tol := 0.0001
// 	cases := []struct {
// 		x, γ, β, μ float64
// 		expected   float64
// 	}{
// 		{2.5, 3, 4, 1, 0.09393379},
// 		{1.098986, 1.876786, 4.986765, 1, 0.008069556},
// 		{7.334, 3.8678756, 2.9878796, 1, 0.01710508},
// 	}

// 	for i, c := range cases {
// 		t.Run(strconv.Itoa(i), func(t *testing.T) {
// 			b := BirnbaumSaunders{c.μ, c.γ, c.β, nil}

// 			res := b.Probability(c.x)

// 			if math.Abs(res-c.expected) > tol {
// 				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
// 			}

// 		})
// 	}
// }

// // Generated with R pfatigue(x, γ, β, μ) from package extraDistr
// func TestBirnbaumSaundersDistribution(t *testing.T) {
// 	tol := 0.0001
// 	cases := []struct {
// 		x, γ, β, μ float64
// 		expected   float64
// 	}{
// 		{2.5, 3, 4, 1, 0.3668504},
// 		{1.098986, 1.876786, 4.986765, 1, 0.0001049426},
// 		{7.334, 3.8678756, 2.9878796, 1, 0.5788141},
// 	}

// 	for i, c := range cases {
// 		t.Run(strconv.Itoa(i), func(t *testing.T) {
// 			b := BirnbaumSaunders{c.μ, c.γ, c.β, nil}

// 			res := b.Distribution(c.x)

// 			if math.Abs(res-c.expected) > tol {
// 				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
// 			}

// 		})
// 	}
// }

// func TestBirnbaumSaundersInverse(t *testing.T) {
// 	tol := 0.0001
// 	cases := []struct {
// 		x, γ, β, μ float64
// 		expected   float64
// 	}{
// 		{0.3668504, 3, 4, 1, 2.5},
// 		{0.0001049426, 1.876786, 4.986765, 1, 1.098986},
// 		{0.5788141, 3.8678756, 2.9878796, 1, 7.334},
// 	}

// 	for i, c := range cases {
// 		t.Run(strconv.Itoa(i), func(t *testing.T) {
// 			b := BirnbaumSaunders{c.μ, c.γ, c.β, nil}

// 			res := b.Inverse(c.x)

// 			if math.Abs(res-c.expected) > tol {
// 				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
// 			}

// 		})
// 	}
// }

// // Generated with Mathematica Variance[BirnbaumSaundersDistribution[a, b]]
// func TestBirnbaumSaundersVariance(t *testing.T) {
// 	tol := 0.0001
// 	cases := []struct {
// 		γ, β     float64
// 		expected float64
// 	}{
// 		{5.3, 7.85, 16.4615},
// 		{2, .5, 96},
// 		{4, .1, 33600.},
// 	}

// 	for i, c := range cases {
// 		t.Run(strconv.Itoa(i), func(t *testing.T) {
// 			b := BirnbaumSaunders{Location: 0, Shape: c.γ, Scale: c.β}

// 			res := b.Variance()
// 			if math.Abs(res-c.expected) > tol {
// 				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
// 			}

// 		})
// 	}
// }

// // Generated with Mathematica Mean[BirnbaumSaundersDistribution[a, b]]
// func TestBirnbaumSaundersMean(t *testing.T) {
// 	tol := 0.0001
// 	cases := []struct {
// 		γ, β     float64
// 		expected float64
// 	}{
// 		{5.3, 7.85, 1.91656},
// 		{2, .5, 6},
// 		{4, .1, 90},
// 	}

// 	for i, c := range cases {
// 		t.Run(strconv.Itoa(i), func(t *testing.T) {
// 			b := BirnbaumSaunders{Location: 0, Shape: c.γ, Scale: c.β}

// 			res := b.Mean()
// 			if math.Abs(res-c.expected) > tol {
// 				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
// 			}

// 		})
// 	}
// }
