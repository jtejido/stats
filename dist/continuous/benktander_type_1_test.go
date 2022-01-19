package continuous

import (
	"fmt"
	"strconv"
	"testing"
)

// Generated with Mathematica PDF[BenktanderGibratDistribution[a, b], x]
func TestBenktanderType1Probability(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		x, a, b  float64
		expected float64
	}{
		{1.23424, 1, 1, 0.732651},
		{4.2346, 2, .5, 0.00784705},
		{7.23243, 4, .1, 0.0000277728},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BenktanderType1{a: c.a, b: c.b}

			res := b.Probability(c.x)
			run_test(t, res, c.expected, tol, fmt.Sprintf("BenktanderType1Probability(%v, %v)", c.a, c.b))

		})
	}
}

// Generated with Mathematica CDF[BenktanderGibratDistribution[a, b], x]
func TestBenktanderType1Distribution(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		x, a, b  float64
		expected float64
	}{
		{1.23424, 1, 1, 0.107657},
		{4.2346, 2, .5, 0.991999},
		{7.23243, 4, .1, 0.999962},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BenktanderType1{a: c.a, b: c.b}

			res := b.Distribution(c.x)
			run_test(t, res, c.expected, tol, fmt.Sprintf("BenktanderType1Distribution(%v, %v)", c.a, c.b))

		})
	}
}

// Generated with Mathematica Variance[BenktanderGibratDistribution[a, b]]
func TestBenktanderType1Variance(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{2, .5, 0.40568},
		{4, .1, 0.100686},
		{10, .7, 0.0118565},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BenktanderType1{a: c.a, b: c.b}

			res := b.Variance()
			run_test(t, res, c.expected, tol, fmt.Sprintf("BenktanderType1Variance(%v, %v)", c.a, c.b))
		})
	}
}

// Generated with Mathematica Mean[BenktanderGibratDistribution[a, b]]
func TestBenktanderType1Mean(t *testing.T) {
	tol := 0.000001
	cases := []struct {
		a, b     float64
		expected float64
	}{
		{2, .5, 1.5},
		{4, .1, 1.25},
		{10, .7, 1.1},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := BenktanderType1{a: c.a, b: c.b}

			res := b.Mean()
			run_test(t, res, c.expected, tol, fmt.Sprintf("BenktanderType1Mean(%v, %v)", c.a, c.b))

		})
	}
}
