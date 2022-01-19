package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
)

// Benktander type I distribution (Benktander-Gibrat Distribution)
// https://en.wikipedia.org/wiki/Benktander_type_I_distribution
type BenktanderType1 struct {
	a, b float64
}

func NewBenktanderType1(a, b float64) (*BenktanderType1, error) {
	if a <= 0 || b <= 0 {
		return nil, err.Invalid()
	}

	return &BenktanderType1{a, b}, nil
}

func (bfk *BenktanderType1) String() string {
	return "BenktanderType1: Parameters - " + bfk.Parameters().String() + ", Support(x) - " + bfk.Support().String()
}

// a ∈ (0,∞)
// b ∈ (0,∞)
func (bfk *BenktanderType1) Parameters() stats.Limits {
	return stats.Limits{
		"a": stats.Interval{0, math.Inf(1), true, true},
		"b": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [1,∞)
func (bfk *BenktanderType1) Support() stats.Interval {
	return stats.Interval{1, math.Inf(1), false, true}
}

func (bfk *BenktanderType1) Probability(x float64) float64 {
	if bfk.Support().IsWithinInterval(x) {
		a := 1 + ((2. * bfk.b * math.Log(x)) / bfk.a)
		b := 1 + bfk.a + 2.*bfk.b*math.Log(x)

		return ((a * b) - ((2. * bfk.b) / bfk.a)) * math.Pow(x, -(2.+bfk.a+bfk.b*math.Log(x)))
	}

	return 0
}

func (bfk *BenktanderType1) Distribution(x float64) float64 {
	if bfk.Support().IsWithinInterval(x) {
		return 1. - (1.+((2.*bfk.b)/bfk.a)*math.Log(x))*math.Pow(x, -(bfk.a+1.+bfk.b*math.Log(x)))
	}

	return 0
}

func (bfk *BenktanderType1) Mean() float64 {
	return 1. + (1 / bfk.a)
}

func (bfk *BenktanderType1) Variance() float64 {
	m1 := bfk.rm(1)
	m2 := bfk.rm(2)
	return -(m1 * m1) + m2
}

func (bfk *BenktanderType1) Skewness() float64 {
	m1 := bfk.rm(1)
	m2 := bfk.rm(2)
	m3 := bfk.rm(3)
	return (m1*(2*(m1*m1)-3*m2) + m3) / math.Pow(m2-(m1*m1), 3./2)
}

func (bfk *BenktanderType1) ExKurtosis() float64 {
	m1 := bfk.rm(1)
	m2 := bfk.rm(2)
	m3 := bfk.rm(3)
	m4 := bfk.rm(4)
	return (-3*(m1*m1*m1*m1) + 6*(m1*m1)*m2 - 4*m1*m3 + m4 - 3*((m2-(m1*m1))*(m2-(m1*m1)))) / ((m2 - (m1 * m1)) * (m2 - (m1 * m1)))
}

func (bfk *BenktanderType1) rm(r float64) float64 {
	sqrtb := math.Sqrt(bfk.b)
	sqrtpi := math.Sqrt(math.Pi)
	n1 := math.Pow(1+bfk.a-r, 2)
	d1 := 4 * bfk.b
	en1d1 := math.Exp(n1 / d1)
	n2 := 1 + bfk.a - r
	d2 := 2 * sqrtb
	erfcn2d2 := math.Erfc(n2 / d2)
	num := 2*bfk.a*sqrtb + 2*sqrtb*r - en1d1*sqrtpi*r*erfcn2d2 + en1d1*sqrtpi*(r*r)*erfcn2d2
	denom := 2 * bfk.a * sqrtb
	return num / denom
}
