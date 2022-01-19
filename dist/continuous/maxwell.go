package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// Maxwell–Boltzmann distribution
// https://en.wikipedia.org/wiki/Maxwell%E2%80%93Boltzmann_distribution
type MaxwellBoltzmann struct {
	scale float64
	src   rand.Source
}

func NewMaxwellBoltzmann(scale float64) (*MaxwellBoltzmann, error) {
	return NewMaxwellBoltzmannWithSource(scale, nil)
}

func NewMaxwellBoltzmannWithSource(scale float64, src rand.Source) (*MaxwellBoltzmann, error) {
	if scale <= 0 {
		return nil, err.Invalid()
	}

	return &MaxwellBoltzmann{scale, src}, nil
}

// σ ∈ (0,∞)
func (mb *MaxwellBoltzmann) Parameters() stats.Limits {
	return stats.Limits{
		"σ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (mb *MaxwellBoltzmann) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

func (mb *MaxwellBoltzmann) Probability(x float64) float64 {
	if mb.Support().IsWithinInterval(x) {
		num := (x * x) * math.Exp(-(x*x)/(2*(mb.scale*mb.scale)))
		denom := mb.scale * mb.scale * mb.scale

		return math.Sqrt(2/math.Pi) * (num / denom)
	}

	return 0
}

func (mb *MaxwellBoltzmann) Distribution(x float64) float64 {
	if mb.Support().IsWithinInterval(x) {
		num := x * math.Exp(-(x*x)/(2.*(mb.scale*mb.scale)))
		denom := mb.scale

		return math.Erf(x/(math.Sqrt(2)*mb.scale)) - (math.Sqrt(2./math.Pi) * (num / denom))
	}

	return 0
}

func (mb *MaxwellBoltzmann) Entropy() float64 {
	return math.Log(mb.scale*(math.Sqrt(2.*math.Pi))) + gsl.Euler - (1. / 2)
}

func (mb *MaxwellBoltzmann) ExKurtosis() float64 {
	num := -96 + (40 * math.Pi) - (3 * (math.Pi * math.Pi))
	denom := math.Pow((3*math.Pi)-8, 2.)

	return 4 * (num / denom)
}

func (mb *MaxwellBoltzmann) Skewness() float64 {
	num := 2. * math.Sqrt(2.) * (16 - (5. * math.Pi))
	denom := math.Pow((3*math.Pi)-8, 3./2)

	return num / denom
}

func (mb *MaxwellBoltzmann) Mean() float64 {
	return 2 * mb.scale * math.Sqrt(2./math.Pi)
}

func (mb *MaxwellBoltzmann) Mode() float64 {
	return math.Sqrt(2) * mb.scale
}

func (mb *MaxwellBoltzmann) Variance() float64 {
	return ((mb.scale * mb.scale) * ((3 * math.Pi) - 8)) / math.Pi
}

func (mb *MaxwellBoltzmann) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return math.Sqrt(2) * mb.scale * math.Sqrt(smath.InverseRegularizedUpperIncompleteGamma(3./2, 1-p))
}

func (mb *MaxwellBoltzmann) Rand() float64 {
	var rnd float64
	if mb.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(mb.src).Float64()
	}

	return mb.Inverse(rnd)
}
