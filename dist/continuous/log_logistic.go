package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"github.com/jtejido/trig"
	"math"
	"math/rand"
)

// Log-logistic distribution
// Also known as the Fisk distribution.
// https://en.wikipedia.org/wiki/Log-logistic_distribution
// https://en.wikipedia.org/wiki/Shifted_log-logistic_distribution
type LogLogistic struct {
	scale, shape, location float64 // α, β, γ
	src                    rand.Source
}

func NewLogLogistic(scale, shape, location float64) (*LogLogistic, error) {
	return NewLogLogisticWithSource(scale, shape, location, nil)
}

func NewLogLogisticWithSource(scale, shape, location float64, src rand.Source) (*LogLogistic, error) {
	if scale <= 0 || shape <= 0 {
		return nil, err.Invalid()
	}

	return &LogLogistic{scale, shape, location, nil}, nil
}

// α ∈ (0,∞)
// β ∈ (0,∞)
// γ ∈ (-∞,∞)
func (ll *LogLogistic) Parameters() stats.Limits {
	return stats.Limits{
		"α": stats.Interval{0, math.Inf(1), true, true},
		"β": stats.Interval{0, math.Inf(1), true, true},
		"γ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
	}
}

// x ∈ [𝛿,∞)
func (ll *LogLogistic) Support() stats.Interval {
	return stats.Interval{ll.location, math.Inf(1), false, true}
}

func (ll *LogLogistic) Probability(x float64) float64 {
	if ll.Support().IsWithinInterval(x) {
		x -= ll.location
		num := (ll.shape / ll.scale) * math.Pow(x/ll.scale, ll.shape-1)
		denom := (1 + math.Pow(x/ll.scale, ll.shape)) * (1 + math.Pow(x/ll.scale, ll.shape))
		return num / denom
	}

	return 0
}

func (ll *LogLogistic) Distribution(x float64) float64 {
	if ll.Support().IsWithinInterval(x) {
		x -= ll.location
		xOveraPownegB := math.Pow(x/ll.scale, -ll.shape)
		return 1 / (1 + xOveraPownegB)
	}

	return 0
}

func (ll *LogLogistic) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return (ll.scale * math.Pow(p/(1-p), 1/ll.shape)) + ll.location
}

func (ll *LogLogistic) Mean() float64 {
	if ll.shape <= 1 {
		return math.NaN()
	}

	theta := math.Pi / ll.shape
	return ll.scale*theta*trig.Csc(theta) + ll.location
}

func (ll *LogLogistic) Median() float64 {
	return ll.scale
}

func (ll *LogLogistic) Mode() float64 {
	if ll.shape <= 1 {
		return ll.location
	}

	return ll.location + (ll.scale * math.Pow((ll.shape-1)/(ll.shape+1), 1/ll.shape))
}

func (ll *LogLogistic) Variance() float64 {
	if ll.shape <= 2 {
		return math.NaN()
	}

	theta := math.Pi / ll.shape
	return (ll.scale * ll.scale) * theta * (2*trig.Csc(2*theta) - theta*(trig.Csc(theta)*trig.Csc(theta)))
}

func (ll *LogLogistic) Skewness() float64 {
	if ll.shape <= 3 {
		return math.NaN()
	}

	theta := math.Pi / ll.shape
	num := 3*trig.Csc(3*theta) - 6*theta*trig.Csc(2*theta)*trig.Csc(theta) + 2*(theta*theta)*(trig.Csc(theta)*trig.Csc(theta)*trig.Csc(theta))
	denom := math.Sqrt(theta) * math.Pow(2*trig.Csc(2*theta)-theta*(trig.Csc(theta)*trig.Csc(theta)), 3./2)
	return num / denom
}

func (ll *LogLogistic) ExKurtosis() float64 {
	if ll.shape <= 4 {
		return math.NaN()
	}
	theta := math.Pi / ll.shape
	num := 4*trig.Csc(4*theta) - 12*theta*trig.Csc(3*theta)*trig.Csc(theta) + 12*(theta*theta)*trig.Csc(2*theta)*(trig.Csc(theta)*trig.Csc(theta)) - 3*(theta*theta*theta)*(trig.Csc(theta)*trig.Csc(theta)*trig.Csc(theta)*trig.Csc(theta))
	denom := theta * math.Pow(2*trig.Csc(2*theta)-theta*(trig.Csc(theta)*trig.Csc(theta)), 2)
	return (num / denom) - 3
}

func (ll *LogLogistic) Rand() float64 {
	var rnd float64
	if ll.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(ll.src).Float64()
	}

	return ll.Inverse(rnd)
}
