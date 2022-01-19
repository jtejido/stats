package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Johnson SN Distribution (Normal)
// https://reference.wolfram.com/language/ref/JohnsonDistribution.html
type JohnsonSN struct {
	gamma, delta, location, scale float64 // γ, δ, location μ, and scale σ
	src                           rand.Source
}

func NewJohnsonSN(gamma, delta, location, scale float64) (*JohnsonSN, error) {
	return NewJohnsonSNWithSource(gamma, delta, location, scale, nil)
}

func NewJohnsonSNWithSource(gamma, delta, location, scale float64, src rand.Source) (*JohnsonSN, error) {
	if delta <= 0 && scale <= 0 {
		return nil, err.Invalid()
	}

	return &JohnsonSN{gamma, delta, location, scale, src}, nil
}

// γ ∈ (-∞,∞)
// δ ∈ (0,∞)
// μ ∈ (-∞,∞)
// σ ∈ (0,∞)
func (j *JohnsonSN) Parameters() stats.Limits {
	return stats.Limits{
		"γ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"δ": stats.Interval{0, math.Inf(1), true, true},
		"μ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"σ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (-∞,∞)
func (j *JohnsonSN) Support() stats.Interval {
	return stats.Interval{math.Inf(-1), math.Inf(1), true, true}
}

func (j *JohnsonSN) Probability(x float64) float64 {
	return (math.Exp(-.5*math.Pow(j.gamma+(j.delta*(x-j.location))/j.scale, 2)) * j.delta) / (math.Sqrt(2*math.Pi) * j.scale)
}

func (j *JohnsonSN) Distribution(x float64) float64 {
	return .5 * math.Erfc(-((j.gamma + (j.delta*(x-j.location))/j.scale) / math.Sqrt(2)))
}

func (j *JohnsonSN) Mean() float64 {
	return (j.delta*j.location - j.gamma*j.scale) / j.delta
}

func (j *JohnsonSN) Variance() float64 {
	return (j.scale * j.scale) / (j.delta * j.delta)
}

func (j *JohnsonSN) Median() float64 {
	return j.location - (j.gamma*j.scale)/j.delta
}

func (j *JohnsonSN) ExKurtosis() float64 {
	return 0
}

func (j *JohnsonSN) Entropy() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (j *JohnsonSN) Inverse(q float64) float64 {
	if q <= 0 {
		return math.Inf(-1)
	}

	if q >= 1 {
		return math.Inf(1)
	}

	return j.location + (j.scale*(-j.gamma-math.Sqrt(2)*math.Erfcinv(2*q)))/j.delta

}

func (j *JohnsonSN) Rand() float64 {
	var rnd float64
	if j.src != nil {
		rnd = rand.New(j.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return j.Inverse(rnd)
}
