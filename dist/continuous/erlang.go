package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// Erlang distribution
// https://en.wikipedia.org/wiki/Erlang_distribution
type Erlang struct {
	baseContinuousWithSource
	shape int
	rate  float64 // k, λ
}

func NewErlang(shape int, rate float64) (*Erlang, error) {
	return NewErlangWithSource(shape, rate, nil)
}

func NewErlangWithSource(shape int, rate float64, src rand.Source) (*Erlang, error) {
	if shape <= 0 || rate <= 0 {
		return nil, err.Invalid()
	}

	r := new(Erlang)
	r.shape = shape
	r.rate = rate
	r.src = src

	return r, nil
}

func (e *Erlang) String() string {
	return "Erlang: Parameters - " + e.Parameters().String() + ", Support(x) - " + e.Support().String()
}

// k ∈ (0,∞)
// λ ∈ (0,∞)
func (e *Erlang) Parameters() stats.Limits {
	return stats.Limits{
		"k": stats.Interval{0, math.Inf(1), true, true},
		"λ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [0,∞)
func (e *Erlang) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), false, true}
}

func (e *Erlang) Probability(x float64) float64 {
	if e.Support().IsWithinInterval(x) {
		return (math.Pow(e.rate, float64(e.shape)) * math.Pow(x, float64(e.shape-1)) * math.Exp(-e.rate*x)) / specfunc.Fact(uint(e.shape-1))
	}

	return 0
}

func (e *Erlang) Distribution(x float64) float64 {
	if e.Support().IsWithinInterval(x) {
		return specfunc.Gamma_inc_P(float64(e.shape), e.rate*x)
	}

	return 0
}

func (e *Erlang) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return smath.InverseRegularizedLowerIncompleteGamma(float64(e.shape), p) / e.rate
}

func (e *Erlang) Entropy() float64 {
	return (1.0-float64(e.shape))*specfunc.Psi(float64(e.shape)) + math.Log(specfunc.Gamma(float64(e.shape))/e.rate) + float64(e.shape)
}

func (e *Erlang) ExKurtosis() float64 {
	return 6. / float64(e.shape)
}

func (e *Erlang) Skewness() float64 {
	return 2. / math.Sqrt(float64(e.shape))
}

func (e *Erlang) Mean() float64 {
	return float64(e.shape) / e.rate
}

func (e *Erlang) Mode() float64 {
	return (1. / e.rate) * (float64(e.shape) - 1.)
}

func (e *Erlang) Variance() float64 {
	return float64(e.shape) / (e.rate * e.rate)
}

func (e *Erlang) Median() float64 {
	return e.Inverse(.5)
}

func (e *Erlang) Rand() float64 {
	var rnd func() float64
	if e.src != nil {
		rnd = rand.New(e.src).Float64
	} else {
		rnd = rand.Float64
	}

	mul := 1.0
	for i := 0; i < e.shape; i++ {
		mul *= rnd()
	}

	return -(1 / e.rate) * math.Log(mul)
}
