package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Johnson SU Distribution (Unbounded)
// https://reference.wolfram.com/language/ref/JohnsonDistribution.html
type JohnsonSU struct {
	gamma, delta, location, scale float64 // γ, δ, location μ, and scale σ
	src                           rand.Source
}

func NewJohnsonSU(gamma, delta, location, scale float64) (*JohnsonSU, error) {
	return NewJohnsonSUWithSource(gamma, delta, location, scale, nil)
}

func NewJohnsonSUWithSource(gamma, delta, location, scale float64, src rand.Source) (*JohnsonSU, error) {
	if delta <= 0 && scale <= 0 {
		return nil, err.Invalid()
	}

	return &JohnsonSU{gamma, delta, location, scale, src}, nil
}

// γ ∈ (-∞,∞)
// δ ∈ (0,∞)
// μ ∈ (-∞,∞)
// σ ∈ (0,∞)
func (j *JohnsonSU) Parameters() stats.Limits {
	return stats.Limits{
		"γ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"δ": stats.Interval{0, math.Inf(1), true, true},
		"μ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"σ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (-∞,∞)
func (j *JohnsonSU) Support() stats.Interval {
	return stats.Interval{math.Inf(-1), math.Inf(1), true, true}
}

func (j *JohnsonSU) Probability(x float64) float64 {
	return (math.Exp(-.5*math.Pow(j.gamma+j.delta*math.Asinh((x-j.location)/j.scale), 2)) * j.delta) / (math.Sqrt(2*math.Pi) * math.Sqrt(((x-j.location)*(x-j.location))+(j.scale*j.scale)))
}

func (j *JohnsonSU) Distribution(x float64) float64 {
	return .5 * (1 + math.Erf(j.gamma+j.delta*math.Asinh((x-j.location)/j.scale)/math.Sqrt(2)))
}

func (j *JohnsonSU) Mean() float64 {
	return j.location - math.Exp(1/(2*(j.delta*j.delta)))*j.scale*math.Sinh(j.gamma/j.delta)
}

func (j *JohnsonSU) Variance() float64 {
	return 1. / 4 * math.Exp(-((2 * j.gamma) / j.delta)) * (-1 + math.Exp(1/(j.delta*j.delta))) * (math.Exp(1/(j.delta*j.delta)) + 2*math.Exp((2*j.gamma)/j.delta) + math.Exp((1+4*j.gamma*j.delta)/(j.delta*j.delta))) * (j.scale * j.scale)
}

func (j *JohnsonSU) Median() float64 {
	return j.location - j.scale*math.Sinh(j.gamma/j.delta)
}

func (j *JohnsonSU) ExKurtosis() float64 {
	num := (4*math.Exp((2+2*j.gamma*j.delta)/(j.delta*j.delta))*(2+math.Exp(1/(j.delta*j.delta))) + 4*math.Exp((2+6*j.gamma*j.delta)/(j.delta*j.delta))*(2+math.Exp(1/(j.delta*j.delta))) + 6*math.Exp((4*j.gamma)/j.delta)*(1+2*math.Exp(1/(j.delta*j.delta))) + math.Exp(2/(j.delta*j.delta))*(-3+math.Exp(2/(j.delta*j.delta))*(3+math.Exp(1/(j.delta*j.delta))*(2+math.Exp(1/(j.delta*j.delta))))) + math.Exp((2+8*j.gamma*j.delta)/(j.delta*j.delta))*(-3+math.Exp(2/(j.delta*j.delta))*(3+math.Exp(1/(j.delta*j.delta))*(2+math.Exp(1/(j.delta*j.delta))))))
	denom := math.Pow(math.Exp(1/(j.delta*j.delta))+2*math.Exp((2*j.gamma)/j.delta)+math.Exp((1+4*j.gamma*j.delta)/(j.delta*j.delta)), 2)

	return (num / denom) - 3
}

func (j *JohnsonSU) Entropy() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (j *JohnsonSU) Inverse(q float64) float64 {
	if q <= 0 {
		return math.Inf(-1)
	}

	if q >= 1 {
		return math.Inf(1)
	}

	return j.location + j.scale*math.Sinh((-j.gamma+math.Sqrt(2)*math.Erfinv(-1+2*q))/j.delta)
}

func (j *JohnsonSU) Rand() float64 {
	var rnd float64
	if j.src != nil {
		rnd = rand.New(j.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	n := Normal{0, 1, j.src, nil}
	return j.scale*math.Sinh((n.Inverse(rnd)-j.gamma)/j.delta) + j.location
}
