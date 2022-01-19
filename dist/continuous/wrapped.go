package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
)

// Wrapped distribution
// https://en.wikipedia.org/wiki/Wrapped_distribution
// K. V. Mardia and P. E. Jupp, Directional Statistics, 1st ed. Wiley, 1999
type Wrapped struct {
	dist    Wrappable
	support stats.Interval // allows flexibility with Support as long as x ∈ [any interval of length 2π].
	k       int
}

func NewWrapped(dist Wrappable, k int, support stats.Interval) (*Wrapped, error) {
	if k <= 0 {
		k = 1000
	}

	if support.Lower >= support.Upper {
		return nil, err.Error("lower cannot be greater than or equal upper", err.EINVAL)
	}

	if !support.IsEqualLength(defaultLength) {
		return nil, err.Error("length not equals 2π", err.EINVAL)
	}

	return &Wrapped{dist, support, k}, nil
}

// k ∈ (0,∞)
func (w *Wrapped) Parameters() stats.Limits {
	return stats.Limits{
		"k": stats.Interval{0, math.Inf(1), true, true},
	}
}

// θ ∈ (any interval of length 2π]
func (w *Wrapped) Support() stats.Interval {
	return w.support
}

// via knb summation
func (w *Wrapped) Probability(θ float64) float64 {
	if w.Support().IsWithinInterval(θ) {
		var c float64
		sum := w.dist.Probability(θ + 2*math.Pi*float64(-w.k))
		for k := -w.k - 1; k < w.k; k++ {
			d := w.dist.Probability(θ + 2*math.Pi*float64(k))
			t := sum + d
			if math.Abs(sum) >= math.Abs(d) {
				c += (sum - t) + d
			} else {
				c += (d - t) + sum
			}

			sum = t
		}

		return sum + c
	}

	return 0
}

// via knb summation
func (w *Wrapped) Distribution(θ float64) float64 {
	if w.Support().IsWithinInterval(θ) {
		var c float64
		fwk := float64(-w.k)
		sum := w.dist.Distribution(θ+2*math.Pi*fwk) - w.dist.Distribution(-math.Pi+2*math.Pi*fwk)
		for k := -w.k - 1; k < w.k; k++ {
			fk := float64(k)
			d := w.dist.Distribution(θ+2*math.Pi*fk) - w.dist.Distribution(-math.Pi+2*math.Pi*fk)
			t := sum + d
			if math.Abs(sum) >= math.Abs(d) {
				c += (sum - t) + d
			} else {
				c += (d - t) + sum
			}

			sum = t
		}

		return sum + c
	}

	return 0
}

func (w *Wrapped) Rand() float64 {
	sup := w.Support()
	return smath.WrapRange(w.dist.Rand(), sup.Lower, sup.Upper, false)
}
