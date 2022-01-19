package continuous

import (
	"github.com/jtejido/linear"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Laplace distribution
// https://en.wikipedia.org/wiki/Laplace_distribution
type Laplace struct {
	location, scale float64 // μ, b
	src             rand.Source
	natural         linear.RealVector
}

func NewLaplace(location, scale float64) (*Laplace, error) {
	return NewLaplaceWithSource(location, scale, nil)
}

func NewLaplaceWithSource(location, scale float64, src rand.Source) (*Laplace, error) {
	if scale <= 0 {
		return nil, err.Invalid()
	}

	return &Laplace{location, scale, src, nil}, nil
}

// μ ∈ (-∞,∞)
// b ∈ (0,∞)
func (l *Laplace) Parameters() stats.Limits {
	return stats.Limits{
		"μ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"b": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (l *Laplace) Support() stats.Interval {
	return stats.Interval{math.Inf(-1), math.Inf(1), true, true}
}

func (l *Laplace) Probability(x float64) float64 {
	if l.Support().IsWithinInterval(x) {
		return (1 / (2 * l.scale)) * math.Exp(-(math.Abs(x-l.location) / l.scale))
	}

	return 0
}

func (l *Laplace) Distribution(x float64) float64 {
	if l.Support().IsWithinInterval(x) {
		z := (x - l.location) / l.scale

		if z < 0 {
			return (1. / 2) * math.Exp(z)
		}

		return 1. - (1./2)*math.Exp(-z)

	}

	return 0
}

func (l *Laplace) Entropy() float64 {
	return 1 + math.Log(2*l.scale)
}

func (l *Laplace) ExKurtosis() float64 {
	return 3
}

func (l *Laplace) Skewness() float64 {
	return 0
}

func (l *Laplace) Inverse(p float64) float64 {
	if p <= 0 {
		return math.Inf(-1)
	}

	if p >= 1 {
		return math.Inf(1)
	}

	if p < .5 {
		return l.location + math.Log(2*p)*l.scale
	}

	return l.location - math.Log(2*(1-p))*l.scale
}

func (l *Laplace) Mean() float64 {
	return l.location
}

func (l *Laplace) Median() float64 {
	return l.location
}

func (l *Laplace) Mode() float64 {
	return l.location
}

func (l *Laplace) Variance() float64 {
	return 2 * (l.scale * l.scale)
}

func (l *Laplace) Rand() float64 {
	var rnd float64
	if l.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(l.src).Float64()
	}
	u := rnd - 0.5

	if u < 0 {
		return l.location + l.scale*math.Log(1+2*u)
	}

	return l.location - l.scale*math.Log(1-2*u)
}

func (l *Laplace) ToExponential() {
	vec, _ := linear.NewArrayRealVectorFromSlice([]float64{-1 / l.scale})
	l.natural = vec
}

func (l *Laplace) SufficientStatistics(x float64) linear.RealVector {
	vec, _ := linear.NewArrayRealVectorFromSlice([]float64{math.Abs(x - l.location)})
	return vec
}
