package continuous

import (
	"math"
	"strconv"
	"testing"
)

// Generated with R dlomax(x, λ, α) from package VGAM
func TestParetoType2Probability(t *testing.T) {
	tol := 0.0000001

	cases := []struct {
		x, λ, α, expected float64
	}{
		{1, 2, 3, 0.2962963},
		{.87675, 4.9876675475, 10.987876, 0.3161843},
		{10.87676, 25.98789, 11.75765678217, 0.005229241},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ParetoType2{xmin: c.λ, shape: c.α}

			res := b.Probability(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Generated with R plomax(x, λ, α) from package VGAM
func TestParetoType2Distribution(t *testing.T) {
	tol := 0.0000001

	cases := []struct {
		x, λ, α, expected float64
	}{
		{1, 2, 3, 0.7037037},
		{.87675, 4.9876675475, 10.987876, 0.831247},
		{10.87676, 25.98789, 11.75765678217, 0.9836044},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ParetoType2{xmin: c.λ, shape: c.α}

			res := b.Distribution(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

func TestParetoType2Inverse(t *testing.T) {
	tol := 0.0000001

	cases := []struct {
		x, λ, α, expected float64
	}{
		{0.7037037, 2, 3, 1},
		{0.831247, 4.9876675475, 10.987876, .87675},
		{0.9836044, 25.98789, 11.75765678217, 10.87676},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ParetoType2{xmin: c.λ, shape: c.α}

			res := b.Inverse(c.x)
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Wolfram Alpha Mean[ParetoDistribution[k,α,μ]] with μ(Location) = 0
func TestParetoType2Mean(t *testing.T) {
	tol := 0.00001

	cases := []struct {
		λ, α, expected float64
	}{
		{2, 3, 1},
		{4.9876675475, 10.987876, 0.499372},
		{25.98789, 11.75765678217, 2.41576},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ParetoType2{xmin: c.λ, shape: c.α}

			res := b.Mean()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Wolfram Alpha Median[ParetoDistribution[k,α,μ]] with μ(Location) = 0
func TestParetoType2Median(t *testing.T) {
	tol := 0.00001

	cases := []struct {
		λ, α, expected float64
	}{
		{2, 3, 0.519842},
		{4.9876675475, 10.987876, 0.324773},
		{25.98789, 11.75765678217, 1.57812},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ParetoType2{xmin: c.λ, shape: c.α}

			res := b.Median()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Wolfram Alpha Skewness[ParetoDistribution[k,α,μ]] with μ(Location) = 0
func TestParetoType2Skewness(t *testing.T) {
	tol := 0.00001

	cases := []struct {
		λ, α, expected float64
	}{
		{2, 3, math.NaN()},
		{4.9876675475, 10.987876, 2.71464},
		{25.98789, 11.75765678217, 2.654149420687},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ParetoType2{xmin: c.λ, shape: c.α}

			res := b.Skewness()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}

// Wolfram Alpha excess kurtosis[ParetoDistribution[k,α,μ]] with μ(Location) = 0
func TestParetoType2ExKurtosis(t *testing.T) {
	tol := 0.0001

	cases := []struct {
		λ, α, expected float64
	}{
		{2, 3, math.NaN()},
		{4.9876675475, 10.987876, 13.4944},
		{25.98789, 11.75765678217, 12.70230080213},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ParetoType2{xmin: c.λ, shape: c.α}

			res := b.ExKurtosis()
			if math.Abs(res-c.expected) > tol {
				t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, res)
			}

		})
	}
}
