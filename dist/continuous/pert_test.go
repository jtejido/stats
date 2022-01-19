package continuous

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPERTProbability(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x, min, max, mode float64
		expected          float64
	}{

		{1.5423, .4, 5.3, 3.4, 0.10583},
		{4.999, -5.2342, 5.3, 0.435, 0.00372473},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := PERT{c.min, c.max, c.mode, nil}

			res := b.Probability(c.x)
			run_test(t, res, c.expected, tol, fmt.Sprintf("PERTProbability(%v, %v, %v) at x = %v", c.min, c.max, c.mode, c.x))

		})
	}
}

func TestPERTMedian(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		min, max, mode float64
		expected       float64
	}{

		{.4, 5.3, 3.4, 3.2599},
		{-5.2342, 5.3, 0.435, 0.332612},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := PERT{c.min, c.max, c.mode, nil}

			res := b.Median()
			run_test(t, res, c.expected, tol, fmt.Sprintf("PERTMedian(%v, %v, %v)", c.min, c.max, c.mode))

		})
	}
}

func TestPERTMean(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		min, max, mode float64
		expected       float64
	}{

		{.4, 5.3, 3.4, 3.21667},
		{-5.2342, 5.3, 0.435, 0.300967},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := PERT{c.min, c.max, c.mode, nil}

			res := b.Mean()
			run_test(t, res, c.expected, tol, fmt.Sprintf("PERTMean(%v, %v, %v)", c.min, c.max, c.mode))

		})
	}
}

func TestPERTVariance(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		min, max, mode float64
		expected       float64
	}{

		{.4, 5.3, 3.4, .838294},
		{-5.2342, 5.3, 0.435, 3.95293},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := PERT{c.min, c.max, c.mode, nil}

			res := b.Variance()
			run_test(t, res, c.expected, tol, fmt.Sprintf("PERTVariance(%v, %v, %v)", c.min, c.max, c.mode))

		})
	}
}

func TestPERTSkewness(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		min, max, mode float64
		expected       float64
	}{

		{.4, 5.3, 3.4, -0.200237},
		{-5.2342, 5.3, 0.435, -0.0674145},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := PERT{c.min, c.max, c.mode, nil}

			res := b.Skewness()
			run_test(t, res, c.expected, tol, fmt.Sprintf("PERTSkewness(%v, %v, %v)", c.min, c.max, c.mode))

		})
	}
}

func TestPERTExKurtosis(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		min, max, mode float64
		expected       float64
	}{

		{.4, 5.3, 3.4, -0.613207},
		{-5.2342, 5.3, 0.435, -0.660607},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := PERT{c.min, c.max, c.mode, nil}

			res := b.ExKurtosis()
			run_test(t, res, c.expected, tol, fmt.Sprintf("PERTExKurtosis(%v, %v, %v)", c.min, c.max, c.mode))
		})
	}
}

func TestPERTMode(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		min, max, mode float64
		expected       float64
	}{

		{.4, 5.3, 3.4, 3.4},
		{-5.2342, 5.3, 0.435, 0.435},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := PERT{c.min, c.max, c.mode, nil}

			res := b.Mode()
			run_test(t, res, c.expected, tol, fmt.Sprintf("PERTMode(%v, %v, %v)", c.min, c.max, c.mode))

		})
	}
}
