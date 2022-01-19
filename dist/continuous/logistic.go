package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Logistic distribution
// https://en.wikipedia.org/wiki/Logistic_distribution
type Logistic struct {
	location, scale float64 // μ, s
	src             rand.Source
}

func NewLogistic(location, scale float64) (*Logistic, error) {
	return NewLogisticWithSource(location, scale, nil)
}

func NewLogisticWithSource(location, scale float64, src rand.Source) (*Logistic, error) {
	if scale <= 0 {
		return nil, err.Invalid()
	}

	return &Logistic{location, scale, src}, nil
}

func (l *Logistic) String() string {
	return "Logistic: " + l.Parameters().String()
}

// μ ∈ (-∞,∞)
// s ∈ (0,∞)
func (l *Logistic) Parameters() stats.Limits {
	return stats.Limits{
		"μ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"s": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (l *Logistic) Support() stats.Interval {
	return stats.Interval{math.Inf(-1), math.Inf(1), true, true}
}

func (l *Logistic) Probability(x float64) float64 {
	if l.Support().IsWithinInterval(x) {
		a := math.Exp(-(x - l.location) / l.scale)
		return a / (l.scale * ((1 + a) * (1 + a)))
	}

	return 0
}

func (l *Logistic) Distribution(x float64) float64 {
	if l.Support().IsWithinInterval(x) {
		a := math.Exp(-(x - l.location) / l.scale)
		return 1. / (1. + a)
	}

	return 0
}

func (l *Logistic) Entropy() float64 {
	return math.Log(l.scale) + 2.
}

func (l *Logistic) ExKurtosis() float64 {
	return 6. / 5.
}

func (l *Logistic) Skewness() float64 {
	return 0
}

func (l *Logistic) Inverse(p float64) float64 {
	if p <= 0 {
		return math.Inf(-1)
	}

	if p >= 1 {
		return math.Inf(1)
	}

	// return l.location + l.scale*math.Log(p/(1-p))
	return l.location - l.scale*(math.Log1p(-p)-math.Log(p))
}

func (l *Logistic) Mean() float64 {
	return l.location
}

func (l *Logistic) Median() float64 {
	return l.location
}

func (l *Logistic) Mode() float64 {
	return l.location
}

func (l *Logistic) Variance() float64 {
	s_sqrd := l.scale * l.scale
	pi_sqrd := math.Pi * math.Pi

	return (s_sqrd * pi_sqrd) / 3
}

func (l *Logistic) Rand() float64 {
	var rnd float64
	if l.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(l.src).Float64()
	}

	return l.Inverse(rnd)
}
