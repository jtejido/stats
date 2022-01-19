package continuous

import (
	"github.com/jtejido/linear"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Pareto distribution
// https://en.wikipedia.org/wiki/Pareto_distribution
type Pareto struct {
	shape, xmin float64 // α, xm
	src         rand.Source
	natural     linear.RealVector
}

func NewPareto(shape, xmin float64) (*Pareto, error) {
	return NewParetoWithSource(shape, xmin, nil)
}

func NewParetoWithSource(shape, xmin float64, src rand.Source) (*Pareto, error) {
	if shape <= 0 || xmin <= 0 {
		return nil, err.Invalid()
	}

	return &Pareto{shape, xmin, src, nil}, nil
}

// a ∈ (0,∞)
// xm ∈ (0,∞)
func (p *Pareto) Parameters() stats.Limits {
	return stats.Limits{
		"α":  stats.Interval{0, math.Inf(1), true, true},
		"xm": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (p *Pareto) Support() stats.Interval {
	return stats.Interval{p.xmin, math.Inf(1), false, true}
}

func (p *Pareto) Probability(x float64) float64 {
	if p.Support().IsWithinInterval(x) {
		a := p.shape
		b := p.xmin
		num := a * math.Pow(b, a)
		denom := math.Pow(x, a+1)
		return num / denom
	}

	return 0
}

func (p *Pareto) Distribution(x float64) float64 {
	if p.Support().IsWithinInterval(x) {
		return 1. - math.Pow(p.xmin/x, p.shape)
	}

	return 0
}

func (p *Pareto) Entropy() float64 {
	return math.Log(p.xmin) - math.Log(p.shape) + (1 + 1/p.shape)
}

func (p *Pareto) ExKurtosis() float64 {
	if p.shape <= 4 {
		return 0
	}
	return 6 * (p.shape*p.shape*p.shape + p.shape*p.shape - 6*p.shape - 2) / (p.shape * (p.shape - 3) * (p.shape - 4))

}

func (p *Pareto) Skewness() float64 {
	if p.shape < 3.0 {
		return 0
	}

	return 2 * (1 + p.shape) / (p.shape - 3) * math.Sqrt((p.shape-2)/p.shape)
}

func (p *Pareto) Inverse(q float64) float64 {
	if q > 0 && q < 1 {
		return p.xmin / math.Pow(1-q, 1/p.shape)
	}

	if q <= 0 {
		return 0
	}

	return math.Inf(1)
}

func (p *Pareto) Mean() float64 {
	if p.shape <= 1 {
		return math.Inf(1)
	}

	return p.shape * p.xmin / (p.shape - 1)
}

func (p *Pareto) Median() float64 {
	return p.shape * math.Pow(2, 1/p.xmin)
}

func (p *Pareto) Mode() float64 {
	return p.shape
}

func (p *Pareto) Variance() float64 {
	if p.shape <= 2 {
		return math.Inf(1)
	}

	return (p.shape * (p.xmin * p.xmin)) / (((p.shape - 1) * (p.shape - 1)) * (p.shape - 2))
}

func (p *Pareto) Rand() float64 {
	var rnd float64
	if p.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(p.src).Float64()
	}

	return p.Inverse(rnd)
}

func (p *Pareto) ToExponential() {
	vec, _ := linear.NewSizedArrayRealVector(1)
	vec.SetEntry(0, -p.shape-1)
	p.natural = vec

	// moment
	// E[logx] = d/dn((n + 1) log(p) - log(-n - 1)) = 1/(-n - 1) + log(p)
}

func (p *Pareto) SufficientStatistics(x float64) linear.RealVector {
	vec, _ := linear.NewArrayRealVectorFromSlice([]float64{math.Log(x)})
	return vec
}
