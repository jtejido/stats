package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/stats"
	"github.com/jtejido/trig"
	"math"
	"math/rand"
)

// Hyperbolic secant distribution
// https://en.wikipedia.org/wiki/Hyperbolic_secant_distribution
type HyperbolicSecant struct {
	src rand.Source
}

func NewHyperbolicSecantWithSource(src rand.Source) (*HyperbolicSecant, error) {
	return &HyperbolicSecant{src}, nil
}

func (hs *HyperbolicSecant) Parameters() stats.Limits {
	return stats.Limits{}
}

// x ∈ (-∞,∞)
func (hs *HyperbolicSecant) Support() stats.Interval {
	return stats.Interval{math.Inf(-1), math.Inf(1), true, true}
}

func (hs *HyperbolicSecant) Probability(x float64) float64 {
	return .5 * trig.Sech((math.Pi/2)*x)
}

func (hs *HyperbolicSecant) Distribution(x float64) float64 {
	return (2. / math.Pi) * math.Atan(math.Exp((math.Pi/2)*x))
}

func (hs *HyperbolicSecant) ExKurtosis() float64 {
	return 2
}

func (hs *HyperbolicSecant) Skewness() float64 {
	return 0
}

func (hs *HyperbolicSecant) Entropy() float64 {
	return (4.0 / math.Pi) * gsl.Catalan
}

func (hs *HyperbolicSecant) Mean() float64 {
	return 0
}

func (hs *HyperbolicSecant) Median() float64 {
	return 0
}

func (hs *HyperbolicSecant) Mode() float64 {
	return 0
}

func (hs *HyperbolicSecant) Variance() float64 {
	return 1
}

func (hs *HyperbolicSecant) Inverse(p float64) float64 {
	if p <= 0 {
		return math.Inf(-1)
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return (2 / math.Pi) * math.Log(math.Tan((math.Pi/2)*p))
}

func (hs *HyperbolicSecant) Rand() float64 {
	var rnd float64
	if hs.src != nil {
		rnd = rand.New(hs.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return hs.Inverse(rnd)
}
