package continuous

import (
	"fmt"
	"strconv"
	"testing"
)

// Wolfram Alpha
func TestModifiedPERTProbability(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x, min, max, mode, shape float64
		expected                 float64
	}{

		{1.5423, .4, 5.3, 3.4, 2, 0.162185},
		{4.999, -5.2342, 5.3, 0.435, 5, 0.0015384},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ModifiedPERT{c.min, c.max, c.mode, c.shape, nil}

			res := b.Probability(c.x)
			run_test(t, res, c.expected, tol, fmt.Sprintf("ModifiedPERTProbability(%v, %v,%v, %v) at x=%v", c.min, c.max, c.mode, c.shape, c.x))
		})
	}
}

// From R package mc2d ppert(min, mode , max , shape)
func TestModifiedPERTDistribution(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		x, min, max, mode, shape float64
		expected                 float64
	}{

		{1.5423, .4, 5.3, 3.4, 2, 0.08928083},
		{4.999, -5.2342, 5.3, 0.435, 5, 0.9998575},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ModifiedPERT{c.min, c.max, c.mode, c.shape, nil}

			res := b.Distribution(c.x)
			run_test(t, res, c.expected, tol, fmt.Sprintf("ModifiedPERTDistribution(%v, %v,%v, %v) at x=%v", c.min, c.max, c.mode, c.shape, c.x))

		})
	}
}

func TestModifiedPERTInverse(t *testing.T) {
	tol := 0.0001
	cases := []struct {
		x, min, max, mode, shape float64
		expected                 float64
	}{

		{0.08928083, .4, 5.3, 3.4, 2, 1.5423},
		{0.9998575, -5.2342, 5.3, 0.435, 5, 4.999},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ModifiedPERT{c.min, c.max, c.mode, c.shape, nil}

			res := b.Inverse(c.x)
			run_test(t, res, c.expected, tol, fmt.Sprintf("ModifiedPERTInverse(%v, %v,%v, %v) at x=%v", c.min, c.max, c.mode, c.shape, c.x))

		})
	}
}

func TestModifiedPERTMedian(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		min, max, mode, shape float64
		expected              float64
	}{

		{.4, 5.3, 3.4, 2, 3.1749},
		{-5.2342, 5.3, 0.435, 5, 0.348948},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ModifiedPERT{c.min, c.max, c.mode, c.shape, nil}

			res := b.Median()
			run_test(t, res, c.expected, tol, fmt.Sprintf("ModifiedPERTMedian(%v, %v,%v, %v) ", c.min, c.max, c.mode, c.shape))

		})
	}
}

func TestModifiedPERTMean(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		min, max, mode, shape float64
		expected              float64
	}{

		{.4, 5.3, 3.4, 2, 3.125},
		{-5.2342, 5.3, 0.435, 5, 0.320114},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ModifiedPERT{c.min, c.max, c.mode, c.shape, nil}

			res := b.Mean()
			run_test(t, res, c.expected, tol, fmt.Sprintf("ModifiedPERTMean(%v, %v,%v, %v) ", c.min, c.max, c.mode, c.shape))

		})
	}
}

func TestModifiedPERTVariance(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		min, max, mode, shape float64
		expected              float64
	}{

		{.4, 5.3, 3.4, 2, 1.18537},
		{-5.2342, 5.3, 0.435, 5, 3.45748},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ModifiedPERT{c.min, c.max, c.mode, c.shape, nil}

			res := b.Variance()
			run_test(t, res, c.expected, tol, fmt.Sprintf("ModifiedPERTVariance(%v, %v,%v, %v) ", c.min, c.max, c.mode, c.shape))

		})
	}
}

func TestModifiedPERTSkewness(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		min, max, mode, shape float64
		expected              float64
	}{

		{.4, 5.3, 3.4, 2, -0.168389},
		{-5.2342, 5.3, 0.435, 5, -0.0686505},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ModifiedPERT{c.min, c.max, c.mode, c.shape, nil}

			res := b.Skewness()
			run_test(t, res, c.expected, tol, fmt.Sprintf("ModifiedPERTSkewness(%v, %v,%v, %v) ", c.min, c.max, c.mode, c.shape))

		})
	}
}

func TestModifiedPERTExKurtosis(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		min, max, mode, shape float64
		expected              float64
	}{

		{.4, 5.3, 3.4, 2, -0.820687},
		{-5.2342, 5.3, 0.435, 5, -0.593638},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ModifiedPERT{c.min, c.max, c.mode, c.shape, nil}

			res := b.ExKurtosis()
			run_test(t, res, c.expected, tol, fmt.Sprintf("ModifiedPERTExKurtosis(%v, %v,%v, %v) ", c.min, c.max, c.mode, c.shape))

		})
	}
}

func TestModifiedPERTMode(t *testing.T) {
	tol := 0.00001
	cases := []struct {
		min, max, mode, shape float64
		expected              float64
	}{

		{.4, 5.3, 3.4, 2, 3.4},
		{-5.2342, 5.3, 0.435, 5, 0.435},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := ModifiedPERT{c.min, c.max, c.mode, c.shape, nil}

			res := b.Mode()
			run_test(t, res, c.expected, tol, fmt.Sprintf("ModifiedPERTMode(%v, %v,%v, %v) ", c.min, c.max, c.mode, c.shape))

		})
	}
}
