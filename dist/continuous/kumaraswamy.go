package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// Kumaraswamy distribution
// https://en.wikipedia.org/wiki/Kumaraswamy_distribution
type Kumaraswamy struct {
	a, b float64
	src  rand.Source
}

func NewKumaraswamy(a, b float64) (*Kumaraswamy, error) {
	return NewKumaraswamyWithSource(a, b, nil)
}

func NewKumaraswamyWithSource(a, b float64, src rand.Source) (*Kumaraswamy, error) {
	if a <= 0 || b <= 0 {
		return nil, err.Invalid()
	}

	return &Kumaraswamy{a, b, nil}, nil
}

// a ∈ [0,∞)
// b ∈ [0,∞)
func (k *Kumaraswamy) Parameters() stats.Limits {
	return stats.Limits{
		"a": stats.Interval{0, math.Inf(1), false, true},
		"b": stats.Interval{0, math.Inf(1), false, true},
	}
}

// x ∈ (0,1)
func (k *Kumaraswamy) Support() stats.Interval {
	return stats.Interval{0, 1, false, false}
}

func (k *Kumaraswamy) Probability(x float64) float64 {
	if k.Support().IsWithinInterval(x) {
		return k.a * k.b * math.Pow(x, k.a-1) * math.Pow(1-math.Pow(x, k.a), k.b-1)
	}

	return 0
}

func (k *Kumaraswamy) Distribution(x float64) float64 {
	if k.Support().IsWithinInterval(x) {
		return 1 - math.Pow(1-math.Pow(x, k.a), k.b)
	} else if x >= 1 {
		return 1
	}

	return 0
}

func (k *Kumaraswamy) Mean() float64 {
	return (k.b * specfunc.Gamma(1+(1./k.a)) * specfunc.Gamma(k.b)) / specfunc.Gamma(1+(1./k.a)+k.b)
}

func (k *Kumaraswamy) Median() float64 {
	return math.Pow(1-math.Pow(2., -1/k.b), 1./k.a)
}

func (k *Kumaraswamy) Mode() float64 {
	if k.a >= 1 && k.b >= 1 {
		return math.Pow((k.a-1)/(k.a*k.b)-1, 1./k.a)
	}

	return math.NaN()
}

func (k *Kumaraswamy) Variance() float64 {
	m1 := k.rm(1)
	return -(m1 * m1) + k.rm(2)
}

func (k *Kumaraswamy) Skewness() float64 {
	m1 := k.rm(1)
	m2 := k.rm(2)
	m3 := k.rm(3)
	return (m1*(2*(m1*m1)-3*m2) + m3) / math.Pow(m2-(m1*m1), 3./2)
}

func (k *Kumaraswamy) ExKurtosis() float64 {
	m1 := k.rm(1)
	m2 := k.rm(2)
	m3 := k.rm(3)
	m4 := k.rm(4)
	return (-3*(m1*m1*m1*m1) + 6*(m1*m1)*m2 - 4*m1*m3 + m4 - 3*((m2-(m1*m1))*(m2-(m1*m1)))) / ((m2 - (m1 * m1)) * (m2 - (m1 * m1)))
}

// Raw Moment
func (k *Kumaraswamy) rm(n int) float64 {
	return (k.b * specfunc.Beta(1.0+float64(n)/k.a, k.b))
}

func (k *Kumaraswamy) Entropy() float64 {
	return (1 - (1. / k.b)) + (1-(1./k.a))*smath.Harmonic(k.b) - math.Log(k.a*k.b)
}

func (k *Kumaraswamy) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return 1
	}

	return math.Pow(1-math.Pow(1-p, 1/k.b), 1/k.a)
}

func (k *Kumaraswamy) Rand() float64 {
	var rnd float64
	if k.src != nil {
		rnd = rand.New(k.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return k.Inverse(rnd)
}
