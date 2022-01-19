package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Benktander type II distribution (Benktander-Weibull Distribution)
// https://en.wikipedia.org/wiki/Benktander_type_II_distribution
type BenktanderType2 struct {
	baseContinuousWithSource
	a, b float64
}

func NewBenktanderType2(a, b float64) (*BenktanderType2, error) {
	return NewBenktanderType2WithSource(a, b, nil)
}

func NewBenktanderType2WithSource(a, b float64, src rand.Source) (*BenktanderType2, error) {
	if a <= 0 || b <= 0 {
		return nil, err.Invalid()
	}

	ret := new(BenktanderType2)
	ret.a = a
	ret.b = b
	ret.src = src

	return ret, nil
}

func (bsk *BenktanderType2) String() string {
	return "BenktanderType2: Parameters - " + bsk.Parameters().String() + ", Support(x) - " + bsk.Support().String()
}

// a ∈ (0,∞)
// b ∈ (0,1]
func (bsk *BenktanderType2) Parameters() stats.Limits {
	return stats.Limits{
		"a": stats.Interval{0, math.Inf(1), true, true},
		"b": stats.Interval{0, 1, true, false},
	}
}

// x ∈ [1,∞)
func (bsk *BenktanderType2) Support() stats.Interval {
	return stats.Interval{1, math.Inf(1), false, true}
}

func (bsk *BenktanderType2) Probability(x float64) float64 {
	if bsk.Support().IsWithinInterval(x) {
		return math.Exp((bsk.a/bsk.b)*(1-math.Pow(x, bsk.b))) * math.Pow(x, bsk.b-2.) * (bsk.a*math.Pow(x, bsk.b) - bsk.b + 1.)
	}

	return 0
}

func (bsk *BenktanderType2) Distribution(x float64) float64 {
	if bsk.Support().IsWithinInterval(x) {
		e := (bsk.a / bsk.b) * (1 - math.Pow(x, bsk.b))
		return 1 - math.Pow(x, bsk.b-1)*math.Exp(e)
	}

	return 0
}

func (bsk *BenktanderType2) Mean() float64 {
	return 1. + (1 / bsk.a)
}

func (bsk *BenktanderType2) Mode() float64 {
	return 1.
}

func (bsk *BenktanderType2) Median() float64 {
	if bsk.b == 1 {
		return (math.Log(2.) / bsk.a) + 1
	}

	num := math.Pow(2, bsk.b/(1-bsk.b)) * bsk.a * math.Exp(bsk.a/(1-bsk.b))
	denom := 1 - bsk.b
	return math.Pow(((1-bsk.b)/bsk.a)*specfunc.Lambert_W0(num/denom), 1/bsk.b)
}

func (bsk *BenktanderType2) Inverse(p float64) float64 {
	if 0 < p && p < 1 && bsk.b == 1 {
		return 1 - (math.Log(1-p) / bsk.a)
	}

	if bsk.b != 1 && p > 0 || p < 1 {
		a := bsk.a * math.Exp(bsk.a/(1-bsk.b)) * (1 - p) * (-bsk.b / (1 - bsk.b))
		b := 1 - bsk.b
		num := (1 - bsk.b) * specfunc.Lambert_W0(a/b)
		denom := bsk.a
		return math.Pow(num/denom, 1/bsk.b)
	}

	if p <= 0 {
		return 1
	}

	return math.Inf(1)
}

func (bsk *BenktanderType2) Variance() float64 {
	m1 := bsk.rm(1)
	m2 := bsk.rm(2)
	return -(m1 * m1) + m2
}

func (bsk *BenktanderType2) Skewness() float64 {
	m1 := bsk.rm(1)
	m2 := bsk.rm(2)
	m3 := bsk.rm(3)
	return (m1*(2*(m1*m1)-3*m2) + m3) / math.Pow(m2-(m1*m1), 3./2)
}

func (bsk *BenktanderType2) ExKurtosis() float64 {
	m1 := bsk.rm(1)
	m2 := bsk.rm(2)
	m3 := bsk.rm(3)
	m4 := bsk.rm(4)
	return (-3*(m1*m1*m1*m1) + 6*(m1*m1)*m2 - 4*m1*m3 + m4 - 3*((m2-(m1*m1))*(m2-(m1*m1)))) / ((m2 - (m1 * m1)) * (m2 - (m1 * m1)))
}

func (bsk *BenktanderType2) rm(r float64) float64 {
	return math.Exp(bsk.a/bsk.b) * (((bsk.a * specfunc.Expint_En(-1+int((1-r)/bsk.b), bsk.a/bsk.b)) / bsk.b) - (1-(1/bsk.b))*specfunc.Expint_En(int((1-r)/bsk.b), bsk.a/bsk.b))
}

func (bsk *BenktanderType2) Rand() float64 {
	var rnd float64
	if bsk.src != nil {
		rnd = rand.New(bsk.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return bsk.Inverse(rnd)
}
