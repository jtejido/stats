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
type ModifiedPERT struct {
	min, max, mode, shape float64 // a,c,b
	src                   rand.Source
}

func NewModifiedPERT(min, max, mode, shape float64) (*ModifiedPERT, error) {
	return NewModifiedPERTWithSource(min, max, mode, shape, nil)
}

func NewModifiedPERTWithSource(min, max, mode, shape float64, src rand.Source) (*ModifiedPERT, error) {
	if shape <= 0 || min <= 0 || mode <= min || max <= mode {
		return nil, err.Invalid()
	}

	return &ModifiedPERT{min, max, mode, shape, src}, nil
}

// min ∈ (0,∞)
// mode ∈ (min,max)
// max ∈ (mode,∞)
// λ ∈ (0,∞)
func (p *ModifiedPERT) Parameters() stats.Limits {
	return stats.Limits{
		"a": stats.Interval{0, math.Inf(1), true, true},
		"b": stats.Interval{p.min, p.max, false, true},
		"c": stats.Interval{p.mode, math.Inf(1), false, true},
		"λ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (min,max)
func (p *ModifiedPERT) Support() stats.Interval {
	return stats.Interval{p.min, p.max, true, true}
}

func (p *ModifiedPERT) alpha() float64 {
	return (p.shape * (p.mode - p.min)) / (p.max - p.min)
}

func (p *ModifiedPERT) beta() float64 {
	return (p.shape * (p.max - p.mode)) / (p.max - p.min)
}

func (p *ModifiedPERT) Probability(x float64) float64 {
	if p.Support().IsWithinInterval(x) {
		a := p.alpha()
		b := p.beta()
		num := math.Pow(p.max-p.min, -1-p.shape) * math.Pow(p.max-x, b) * math.Pow(x-p.min, a)
		denom := specfunc.Beta(1+a, 1+b)
		return num / denom
	}

	return 0
}

func (p *ModifiedPERT) Distribution(x float64) float64 {
	if p.Support().IsWithinInterval(x) {
		a := p.alpha()
		b := p.beta()
		z := (x - p.min) / (p.max - p.min)
		return specfunc.Beta_inc(1+a, 1+b, z)
	} else if x >= p.max {
		return 1
	}

	return 0
}

func (p *ModifiedPERT) Entropy() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (p *ModifiedPERT) ExKurtosis() float64 {
	num := 3 * (3 + p.shape) * (-(p.mode*p.mode)*(-4+p.shape)*(p.shape*p.shape) + p.mode*p.max*(-4+p.shape)*(p.shape*p.shape) + p.mode*p.min*(-4+p.shape)*(p.shape*p.shape) + (p.max*p.max)*(4+p.shape*(5+3*p.shape)) + (p.min*p.min)*(4+p.shape*(5+3*p.shape)) - p.max*p.min*(8+p.shape*(10+p.shape*(2+p.shape))))
	denom := (4 + p.shape) * (5 + p.shape) * (p.max - p.min - p.mode*p.shape + p.max*p.shape) * (p.max + p.mode*p.shape - p.min*(1+p.shape))
	return (num / denom) - 3
}

func (p *ModifiedPERT) Skewness() float64 {
	num := 2 * (-2*p.mode + p.max + p.min) * p.shape * math.Sqrt(3+p.shape)
	denom := (4 + p.shape) * math.Sqrt((p.max-p.min-p.mode*p.shape+p.max*p.shape)*(p.max+p.mode*p.shape-p.min*(1+p.shape)))
	return num / denom
}

func (p *ModifiedPERT) Inverse(q float64) float64 {
	if q >= 1 {
		return p.max
	}

	if q <= 0 {
		return p.min
	}

	a := p.alpha()
	b := p.beta()

	return p.min + (p.max-p.min)*smath.InverseRegularizedIncompleteBeta(1+a, 1+b, q)
}

func (p *ModifiedPERT) Mean() float64 {
	return (p.max + p.min + p.mode*p.shape) / (2 + p.shape)
}

func (p *ModifiedPERT) Median() float64 {
	a := p.alpha()
	b := p.beta()
	return p.min + (p.max-p.min)*smath.InverseRegularizedIncompleteBeta(1+a, 1+b, 1/2.)
}

func (p *ModifiedPERT) Mode() float64 {
	return p.mode
}

func (p *ModifiedPERT) Variance() float64 {
	num := (p.max - p.min - p.mode*p.shape + p.max*p.shape) * (p.max + p.mode*p.shape - p.min*(1+p.shape))
	denom := math.Pow(2+p.shape, 2) * (3 + p.shape)
	return num / denom
}

func (p *ModifiedPERT) Rand() float64 {
	var rnd float64
	if p.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(p.src).Float64()
	}

	return p.Inverse(rnd)
}
