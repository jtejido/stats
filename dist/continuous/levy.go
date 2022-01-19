package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Levy distribution
// https://en.wikipedia.org/wiki/L%C3%A9vy_distribution
type Levy struct {
	location, scale float64 // μ, c
	src             rand.Source
}

func NewLevy(location, scale float64) (*Levy, error) {
	return NewLevyWithSource(location, scale, nil)
}

func NewLevyWithSource(location, scale float64, src rand.Source) (*Levy, error) {
	if location < 0 || scale <= 0 {
		return nil, err.Invalid()
	}

	return &Levy{location, scale, src}, nil
}

// μ ∈ [0,∞)
// c ∈ (0,∞)
func (l *Levy) Parameters() stats.Limits {
	return stats.Limits{
		"μ": stats.Interval{0, math.Inf(1), false, true},
		"c": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [μ,∞)
func (l *Levy) Support() stats.Interval {
	return stats.Interval{l.location, math.Inf(1), false, true}
}

func (l *Levy) Probability(x float64) float64 {
	if l.Support().IsWithinInterval(x) {
		return math.Sqrt(l.scale/(2.0*math.Pi)) * math.Exp(-(l.scale / (2.0 * (x - l.location)))) / math.Pow(x-l.location, 1.5)
	}

	return 0
}

func (l *Levy) Distribution(x float64) float64 {
	if l.Support().IsWithinInterval(x) {
		return specfunc.Erfc(math.Sqrt(l.scale / (2 * (x - l.location))))
	}

	return 0
}

func (l *Levy) Entropy() float64 {
	return (1.0 - 3.0*gsl.Euler + math.Log(16.0*math.Pi*l.scale*l.scale)) / 2.0
}

func (l *Levy) ExKurtosis() float64 {
	return math.NaN()
}

func (l *Levy) Skewness() float64 {
	return math.NaN()
}

func (l *Levy) Inverse(p float64) float64 {
	if p <= 0 {
		return l.location
	}

	if p >= 1 {
		return math.Inf(1)
	}

	sn := &Normal{0, 1, nil, nil}
	return l.location + (l.scale / math.Pow(sn.Inverse(1-p/2), 2))
}

func (l *Levy) Mean() float64 {
	return math.Inf(1)
}

func (l *Levy) Median() float64 {
	return l.location + (l.scale / (2 * math.Pow(math.Erfcinv(.5), 2)))
}

func (l *Levy) Mode() float64 {
	return l.location + (l.scale / 3)
}

func (l *Levy) Variance() float64 {
	return math.Inf(1)
}

func (l *Levy) Rand() float64 {
	var rnd float64
	if l.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(l.src).Float64()
	}

	return l.Inverse(rnd)
}
