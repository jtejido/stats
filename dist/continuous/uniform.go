package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Continuous uniform distribution
// https://en.wikipedia.org/wiki/Uniform_distribution_(continuous)
type Uniform struct {
	min, max float64 // a, b
	src      rand.Source
}

func NewUniform(min, max float64) (*Uniform, error) {
	return NewUniformWithSource(min, max, nil)
}

func NewUniformWithSource(min, max float64, src rand.Source) (*Uniform, error) {
	if min >= max {
		return nil, err.Invalid()
	}

	return &Uniform{min, max, src}, nil
}

// a ∈ (-∞,∞)
// b ∈ (a,∞)
func (u *Uniform) Parameters() stats.Limits {
	return stats.Limits{
		"min": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"max": stats.Interval{u.min, math.Inf(1), true, true},
	}
}

// x ∈ [a,b]
func (u *Uniform) Support() stats.Interval {
	return stats.Interval{u.min, u.max, false, false}
}

func (u *Uniform) Probability(x float64) float64 {
	if u.Support().IsWithinInterval(x) {
		return 1. / (u.max - u.min)
	}

	return 0
}

func (u *Uniform) Distribution(x float64) float64 {
	if x < u.min {
		return 0
	}

	if x > u.max {
		return 1
	}

	return (x - u.min) / (u.max - u.min)
}

func (u *Uniform) Entropy() float64 {
	return math.Log(u.max - u.min)
}

func (u *Uniform) ExKurtosis() float64 {
	return -(6. / 5)
}

func (u *Uniform) Skewness() float64 {
	return 0
}

func (u *Uniform) Inverse(p float64) float64 {
	if p <= 0 {
		return u.min
	}

	if p >= 1 {
		return u.max
	}

	return p*(u.max-u.min) + u.min
}

func (u *Uniform) Mean() float64 {
	return (u.min + u.max) / 2
}

func (u *Uniform) Median() float64 {
	return (u.min + u.max) / 2
}

func (u *Uniform) Mode() float64 {
	return u.min
}

func (u *Uniform) Variance() float64 {
	return ((u.max - u.min) * (u.max - u.min)) / 12.
}

func (u *Uniform) Rand() float64 {
	var rnd float64
	if u.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(u.src).Float64()
	}

	return u.Inverse(rnd)
}
