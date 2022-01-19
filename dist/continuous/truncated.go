package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Truncated distribution
// https://en.wikipedia.org/wiki/Truncated_distribution
type Truncated struct {
	baseContinuousWithSource
	dist     Truncatable
	min, max float64
}

func NewTruncated(dist Truncatable, min, max float64) (*Truncated, error) {
	return NewTruncatedWithSource(dist, min, max, nil)
}

func NewTruncatedWithSource(dist Truncatable, min, max float64, src rand.Source) (*Truncated, error) {
	if dist == nil || max < min {
		return nil, err.Invalid()
	}

	ret := new(Truncated)
	ret.dist = dist
	ret.min = min
	ret.max = max
	ret.src = src

	return ret, nil
}

// a ∈ (-∞,∞)
// b ∈ (a,∞)
func (t *Truncated) Parameters() stats.Limits {
	return stats.Limits{
		"min": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"max": stats.Interval{t.min, math.Inf(1), true, true},
	}
}

// x ∈ (a,b]
func (t *Truncated) Support() stats.Interval {
	return stats.Interval{t.min, t.max, true, false}
}

func (t *Truncated) Probability(x float64) float64 {
	if t.Support().IsWithinInterval(x) {
		return t.dist.Probability(x) / (t.dist.Distribution(t.max) - t.dist.Distribution(t.min))
	}

	return 0
}

func (t *Truncated) Distribution(x float64) float64 {
	if t.Support().IsWithinInterval(x) {
		return (t.dist.Distribution(x) - t.dist.Distribution(t.min)) / (t.dist.Distribution(t.max) - t.dist.Distribution(t.min))
	}

	return 0
}

func (t *Truncated) Inverse(q float64) float64 {
	if q <= 0 {
		return t.min
	}

	if q >= 1 {
		return t.max
	}

	return t.dist.Inverse(t.dist.Distribution(t.min) + q*t.dist.Distribution(t.max) - t.dist.Distribution(t.min))
}

func (t *Truncated) Rand() float64 {
	var rnd float64
	if t.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(t.src).Float64()
	}

	return t.Inverse(rnd)
}
