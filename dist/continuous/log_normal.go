package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// LogNormal distribution
// https://en.wikipedia.org/wiki/Log-normal_distribution
type LogNormal struct {
	location, scale float64 // μ, σ
	src             rand.Source
}

func NewLogNormal(location, scale float64) (*LogNormal, error) {
	return NewLogNormalWithSource(location, scale, nil)
}

func NewLogNormalWithSource(location, scale float64, src rand.Source) (*LogNormal, error) {
	if scale <= 0 {
		return nil, err.Invalid()
	}

	return &LogNormal{location, scale, src}, nil
}

// μ ∈ (-∞,∞)
// σ ∈ (0,∞)
func (ln *LogNormal) Parameters() stats.Limits {
	return stats.Limits{
		"μ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"σ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (ln *LogNormal) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

func (ln *LogNormal) Probability(x float64) float64 {
	if ln.Support().IsWithinInterval(x) {
		xσsqrt2π := x * ln.scale * math.Sqrt(2.*math.Pi)
		lnxmμ_sqrd := (math.Log(x) - ln.location) * (math.Log(x) - ln.location)
		σ_sqrd := ln.scale * ln.scale
		return (1. / xσsqrt2π) * math.Exp(-(lnxmμ_sqrd / (2. * σ_sqrd)))
	}

	return 0
}

func (ln *LogNormal) Distribution(x float64) float64 {
	if ln.Support().IsWithinInterval(x) {
		d := &Normal{ln.location, ln.scale, nil, nil}
		return d.Distribution(math.Log(x))
	}

	return 0
}

func (ln *LogNormal) Entropy() float64 {
	return 0.5 + 0.5*math.Log(2*math.Pi*ln.scale*ln.scale) + ln.location
}

func (ln *LogNormal) ExKurtosis() float64 {
	s2 := ln.scale * ln.scale
	return math.Exp(4*s2) + 2*math.Exp(3*s2) + 3*math.Exp(2*s2) - 6
}

func (ln *LogNormal) Skewness() float64 {
	s2 := ln.scale * ln.scale
	return (math.Exp(s2) + 2.0) * math.Sqrt(math.Exp(s2)-1.0)
}

func (ln *LogNormal) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	d := &Normal{ln.location, ln.scale, nil, nil}
	return math.Exp(d.Inverse(p))
}

func (ln *LogNormal) Mean() float64 {
	return math.Exp(ln.location + ((ln.scale * ln.scale) / 2.))
}

func (ln *LogNormal) Median() float64 {
	return math.Exp(ln.location)
}

func (ln *LogNormal) Mode() float64 {
	return math.Exp(ln.location - (ln.scale * ln.scale))
}

func (ln *LogNormal) Variance() float64 {
	σ_sqrd := ln.scale * ln.scale
	μ2 := 2. * ln.location

	return (math.Exp(σ_sqrd) - 1.) * math.Exp(μ2+σ_sqrd)
}

func (ln *LogNormal) Rand() float64 {
	var rnd float64
	if ln.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(ln.src).Float64()
	}

	return ln.Inverse(rnd)
}
