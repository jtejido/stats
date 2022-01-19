package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Rayleigh distribution
// https://en.wikipedia.org/wiki/Rayleigh_distribution
type Rayleigh struct {
	scale float64 // σ
	src   rand.Source
}

func NewRayleigh(scale float64) (*Rayleigh, error) {
	return NewRayleighWithSource(scale, nil)
}

func NewRayleighWithSource(scale float64, src rand.Source) (*Rayleigh, error) {
	if scale <= 0 {
		return nil, err.Invalid()
	}

	return &Rayleigh{scale, src}, nil
}

// σ ∈ (0,∞)
func (r *Rayleigh) Parameters() stats.Limits {
	return stats.Limits{
		"σ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [0,∞)
func (r *Rayleigh) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), false, true}
}

func (r *Rayleigh) Probability(x float64) float64 {
	if r.Support().IsWithinInterval(x) {
		return (x / (r.scale * r.scale)) * math.Exp((-(x * x))/(2*(r.scale*r.scale)))
	}

	return 0
}

func (r *Rayleigh) Distribution(x float64) float64 {
	if r.Support().IsWithinInterval(x) {
		return 1.0 - math.Exp((-(x * x))/(2*(r.scale*r.scale)))
	}

	return 0
}

func (r *Rayleigh) Entropy() float64 {
	return 1.0 + math.Log(r.scale) - math.Log(math.Sqrt(2.0)) + gsl.Euler/2.0
}

func (r *Rayleigh) ExKurtosis() float64 {
	return -(6.0*(math.Pi*math.Pi) - 24.0*math.Pi + 16.0) / ((4.0 - math.Pi) * (4.0 - math.Pi))
}

func (r *Rayleigh) Skewness() float64 {
	num := 2 * math.Sqrt(math.Pi) * (math.Pi - 3)
	denom := math.Pow(4-math.Pi, 1.5)
	return num / denom
}

func (r *Rayleigh) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return math.Sqrt(-2.0 * (r.scale * r.scale) * math.Log(1.0-p))
}

func (r *Rayleigh) Mean() float64 {
	return r.scale * math.Sqrt(math.Pi/2.)
}

func (r *Rayleigh) Median() float64 {
	return r.scale * math.Sqrt(2*math.Log(2.))
}

func (r *Rayleigh) Mode() float64 {
	return r.scale
}

func (r *Rayleigh) Variance() float64 {
	return ((4 - math.Pi) / 2) * (r.scale * r.scale)
}

func (r *Rayleigh) Rand() float64 {
	var rnd float64
	if r.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(r.src).Float64()
	}

	return r.Inverse(rnd)
}
