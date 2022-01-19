package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Pareto type-II
// At μ = 0, see https://en.wikipedia.org/wiki/Lomax_distribution
// https://reference.wolfram.com/language/ref/ParetoDistribution.html
type ParetoType2 struct {
	xmin, shape, location float64 // xm, α, μ
	src                   rand.Source
}

func NewParetoType2(xmin, shape, location float64) (*ParetoType2, error) {
	return NewParetoType2WithSource(xmin, shape, location, nil)
}

func NewParetoType2WithSource(xmin, shape, location float64, src rand.Source) (*ParetoType2, error) {
	if xmin <= 0 || shape <= 0 {
		return nil, err.Invalid()
	}

	return &ParetoType2{xmin, shape, location, src}, nil
}

// xm ∈ (0,∞)
// α ∈ (0,∞)
// μ ∈ (-∞,∞)
func (p *ParetoType2) Parameters() stats.Limits {
	return stats.Limits{
		"xm": stats.Interval{0, math.Inf(1), true, true},
		"α":  stats.Interval{0, math.Inf(1), true, true},
		"μ":  stats.Interval{math.Inf(-1), math.Inf(1), true, true},
	}
}

// x ∈ [μ,∞)
func (p *ParetoType2) Support() stats.Interval {
	return stats.Interval{p.location, math.Inf(1), false, true}
}

func (p *ParetoType2) Probability(x float64) float64 {
	if p.Support().IsWithinInterval(x) {
		return (p.shape * math.Pow((p.xmin+x-p.location)/p.xmin, -1-p.shape)) / p.xmin
	}

	return 0
}

func (p *ParetoType2) Distribution(x float64) float64 {
	if p.Support().IsWithinInterval(x) {
		return 1 - math.Pow(1+((x-p.location)/p.xmin), -p.shape)
	}

	return 0
}

func (p *ParetoType2) Inverse(q float64) float64 {
	if q <= 0 {
		return p.location
	}

	if q >= 1 {
		return math.Inf(1)
	}

	return p.xmin*math.Pow(-1+(1-q), (-1/p.shape)) + p.location
}

func (p *ParetoType2) Mean() float64 {
	if p.shape > 1 {
		return (p.xmin / (p.shape - 1)) + p.location
	}

	return math.NaN()
}

func (p *ParetoType2) Median() float64 {
	return (p.xmin * (math.Pow(2., 1/p.shape) - 1)) + p.location
}

func (p *ParetoType2) Mode() float64 {
	return p.location
}

func (p *ParetoType2) Variance() float64 {
	if p.shape > 2 {
		return (math.Pow(p.xmin, 2.) * p.shape) / (math.Pow(p.shape-1, 2.) * (p.shape - 2))
	}

	if p.shape > 1 && p.shape <= 2 {
		return math.Inf(1)
	}

	return math.NaN()
}

func (p *ParetoType2) ExKurtosis() float64 {
	if p.shape > 4 {
		return (6 * ((p.shape * p.shape * p.shape) + (p.shape * p.shape) - 6*p.shape - 2)) / (p.shape * (p.shape - 3) * (p.shape - 4))
	}

	return math.NaN()
}

func (p *ParetoType2) Skewness() float64 {
	if p.shape > 3 {
		return ((2 * (1 + p.shape)) / (p.shape - 3)) * math.Sqrt((p.shape-2)/p.shape)
	}

	return math.NaN()
}

func (p *ParetoType2) Rand() float64 {
	var rnd float64
	if p.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(p.src).Float64()
	}

	return p.Inverse(rnd)
}
