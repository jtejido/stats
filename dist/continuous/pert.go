package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// PERT distribution
// https://en.wikipedia.org/wiki/PERT_distribution
type PERT struct {
	min, max, mode float64 // a,c,b
	src            rand.Source
}

func NewPERT(min, max, mode float64) (*PERT, error) {
	return NewPERTWithSource(min, max, mode, nil)
}

func NewPERTWithSource(min, max, mode float64, src rand.Source) (*PERT, error) {
	if min <= 0 || mode <= min || max <= mode {
		return nil, err.Invalid()
	}

	return &PERT{min, max, mode, src}, nil
}

// min ∈ (0,∞)
// mode ∈ (min,max)
// max ∈ (mode,∞)
func (p *PERT) Parameters() stats.Limits {
	return stats.Limits{
		"a": stats.Interval{0, math.Inf(1), true, true},
		"b": stats.Interval{p.min, p.max, false, true},
		"c": stats.Interval{p.mode, math.Inf(1), false, true},
	}
}

// x ∈ (min,max)
func (p *PERT) Support() stats.Interval {
	return stats.Interval{p.min, p.max, true, true}
}

func (p *PERT) alpha() float64 {
	return (4*p.mode + p.max - 5*p.min) / (p.max - p.min)
}

func (p *PERT) beta() float64 {
	return (5*p.max - p.min - 4*p.mode) / (p.max - p.min)
}

func (p *PERT) Probability(x float64) float64 {
	if p.Support().IsWithinInterval(x) {
		a := p.alpha()
		b := p.beta()
		return (math.Pow(x-p.min, a-1) * math.Pow(p.max-x, b-1)) / (specfunc.Beta(a, b) * math.Pow(p.max-p.min, a+b-1))
	} else if x >= p.max {
		return 1
	}

	return 0
}

func (p *PERT) Distribution(x float64) float64 {
	if p.Support().IsWithinInterval(x) {
		a := p.alpha()
		b := p.beta()
		z := (x - p.min) / (p.max - p.min)
		return specfunc.Beta_inc(a, b, z)
	}

	return 0
}

func (p *PERT) Entropy() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (p *PERT) ExKurtosis() float64 {
	a := p.alpha()
	b := p.beta()
	return (6 * (math.Pow(a-b, 2)*(a+b+1) - a*b*(a+b+2))) / (a * b * (a + b + 2) * (a + b + 3))
}

func (p *PERT) Skewness() float64 {
	a := p.alpha()
	b := p.beta()
	return (2 * (b - a) * math.Sqrt(a+b+1)) / ((a + b + 2) * math.Sqrt(a*b))
}

func (p *PERT) Inverse(q float64) float64 {
	if q >= 1 {
		return p.max
	}

	if q <= 0 {
		return p.min
	}

	return p.min + (p.max-p.min)*smath.InverseRegularizedIncompleteBeta(1+((p.mode-p.min)/(p.max-p.min)), 1+((-p.mode+p.max)/(p.max-p.min)), q)
}

func (p *PERT) Mean() float64 {
	return (p.min + 4*p.mode + p.max) / 6
}

func (p *PERT) Median() float64 {
	a := p.alpha()
	b := p.beta()
	return p.min + (p.max-p.min)*smath.InverseRegularizedIncompleteBeta(a, b, 1/2.)
}

func (p *PERT) Mode() float64 {
	return p.mode
}

func (p *PERT) Variance() float64 {
	mean := (p.min + 4*p.mode + p.max) / 6
	return ((mean - p.min) * (p.max - mean)) / 7
}

func (p *PERT) Rand() float64 {
	var rnd float64
	if p.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(p.src).Float64()
	}

	return p.Inverse(rnd)
}
