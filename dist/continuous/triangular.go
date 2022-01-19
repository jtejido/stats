package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Triangular distribution
// https://en.wikipedia.org/wiki/Triangular_distribution
type Triangular struct {
	min, max, mode float64 // a, b, c
	src            rand.Source
}

func NewTriangular(min, max, mode float64) (*Triangular, error) {
	return NewTriangularWithSource(min, max, mode, nil)
}

func NewTriangularWithSource(min, max, mode float64, src rand.Source) (*Triangular, error) {
	if min >= max || mode > min || mode > max {
		return nil, err.Invalid()
	}

	return &Triangular{min, max, mode, src}, nil
}

// a ∈ (-∞,∞)
// b ∈ (a,∞)
// c ∈ [a,b]
func (t *Triangular) Parameters() stats.Limits {
	return stats.Limits{
		"min":  stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"max":  stats.Interval{t.min, math.Inf(1), true, true},
		"mode": stats.Interval{t.min, t.max, false, false},
	}
}

// x ∈ [a,b]
func (t *Triangular) Support() stats.Interval {
	return stats.Interval{t.min, t.max, false, false}
}

func (t *Triangular) Probability(x float64) float64 {
	switch {
	case x < t.min:
		return 0
	case x < t.mode:
		return 2 * (x - t.min) / ((t.max - t.min) * (t.mode - t.min))
	case x == t.mode:
		return 2 / (t.max - t.min)
	case x <= t.max:
		return 2 * (t.max - x) / ((t.max - t.min) * (t.max - t.mode))
	default:
		return 0
	}
}

func (t *Triangular) Distribution(x float64) float64 {
	switch {
	case x <= t.min:
		return 0
	case x <= t.mode:
		d := x - t.min
		return (d * d) / ((t.max - t.min) * (t.mode - t.min))
	case x < t.max:
		d := t.max - x
		return 1 - (d*d)/((t.max-t.min)*(t.max-t.mode))
	default:
		return 1
	}
}

func (t *Triangular) Mean() float64 {
	return (t.min + t.max + t.mode) / 3
}

func (t *Triangular) Median() float64 {
	if t.mode >= (t.min+t.max)/2 {
		return t.min + math.Sqrt((t.max-t.min)*(t.mode-t.min)/2)
	}
	return t.max - math.Sqrt((t.max-t.min)*(t.max-t.mode)/2)
}

func (t *Triangular) Mode() float64 {
	return t.mode
}

func (t *Triangular) Skewness() float64 {
	num := math.Sqrt2 * (t.min + t.max - 2*t.mode) * (2*t.min - t.max - t.mode) * (t.min - 2*t.max + t.mode)
	denom := 5 * math.Pow(t.min*t.min+t.max*t.max+t.mode*t.mode-t.min*t.max-t.min*t.mode-t.max*t.mode, 3.0/2.0)

	return num / denom
}

func (t *Triangular) ExKurtosis() float64 {
	return -3.0 / 5.0
}

func (t *Triangular) Entropy() float64 {
	return .5 * math.Log((t.max-t.min)/2)
}

func (t *Triangular) Variance() float64 {
	return (t.min*t.min + t.max*t.max + t.mode*t.mode - t.min*t.max - t.min*t.mode - t.max*t.mode) / 18
}

func (t *Triangular) Inverse(p float64) float64 {

	f := (t.mode - t.min) / (t.max - t.min)

	if 0 <= p && p <= f {
		return t.min + math.Sqrt(p*(t.max-t.min)*(t.mode-t.min))
	}

	if 1 >= p && p > f {
		return t.max - math.Sqrt((1-p)*(t.max-t.min)*(t.max-t.mode))
	}

	if p < 0 {
		return t.min
	}

	return t.max

}

func (t *Triangular) Rand() float64 {
	var rnd float64
	if t.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(t.src).Float64()
	}

	return t.Inverse(rnd)
}
