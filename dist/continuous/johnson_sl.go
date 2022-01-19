package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Johnson SL Distribution (Semi-bounded)
// https://reference.wolfram.com/language/ref/JohnsonDistribution.html
type JohnsonSL struct {
	gamma, delta, location, scale float64 // γ, δ, location μ, and scale σ
	src                           rand.Source
}

func NewJohnsonSL(gamma, delta, location, scale float64) (*JohnsonSL, error) {
	return NewJohnsonSLWithSource(gamma, delta, location, scale, nil)
}

func NewJohnsonSLWithSource(gamma, delta, location, scale float64, src rand.Source) (*JohnsonSL, error) {
	if delta <= 0 && scale <= 0 {
		return nil, err.Invalid()
	}

	return &JohnsonSL{gamma, delta, location, scale, src}, nil
}

// γ ∈ (-∞,∞)
// δ ∈ (0,∞)
// μ ∈ (-∞,∞)
// σ ∈ (0,∞)
func (j *JohnsonSL) Parameters() stats.Limits {
	return stats.Limits{
		"γ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"δ": stats.Interval{0, math.Inf(1), true, true},
		"μ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"σ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (μ,∞)
func (j *JohnsonSL) Support() stats.Interval {
	return stats.Interval{j.location, math.Inf(1), true, true}
}

func (j *JohnsonSL) Probability(x float64) float64 {
	if x > j.location {
		return j.delta / (math.Exp(.5*math.Pow(j.gamma+j.delta*math.Log((x-j.location)/j.scale), 2)) * (math.Sqrt(2*math.Pi) * (x - j.location)))
	}

	return 0

}

func (j *JohnsonSL) Distribution(x float64) float64 {
	if x >= j.location && x <= j.location+j.scale {
		return .5 * (math.Erfc(-(j.gamma + j.delta*math.Log((x-j.location)/j.scale)/math.Sqrt(2))))
	} else if x > j.location+j.scale {
		return .5 * (1 + math.Erf(j.gamma+j.delta*math.Log((x-j.location)/j.scale)/math.Sqrt(2)))
	}

	return 0

}

func (j *JohnsonSL) Mean() float64 {
	return j.location + math.Exp((1-2*j.gamma*j.delta)/(2*(j.delta*j.delta)))*j.scale
}

func (j *JohnsonSL) Variance() float64 {
	return math.Exp((1-2*j.gamma*j.delta)/(j.delta*j.delta)) * (-1 + math.Exp(1/(j.delta*j.delta))) * (j.scale * j.scale)
}

func (j *JohnsonSL) Median() float64 {
	return j.location + j.scale/math.Exp(j.gamma/j.delta)
}

func (j *JohnsonSL) ExKurtosis() float64 {
	return math.Exp(2/(j.delta*j.delta)) * (3 + math.Exp(1/(j.delta*j.delta))*(2+math.Exp(1/(j.delta*j.delta))))
}

func (j *JohnsonSL) Entropy() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (j *JohnsonSL) Inverse(q float64) float64 {
	if q <= 0 {
		return j.location
	}

	if q >= 1 {
		return math.Inf(1)
	}

	return j.location + math.Exp((-j.gamma-math.Sqrt(2)*math.Erfcinv(2*q))/j.delta)*j.scale
}

func (j *JohnsonSL) Rand() float64 {
	var rnd float64
	if j.src != nil {
		rnd = rand.New(j.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return j.Inverse(rnd)
}
