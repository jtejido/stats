package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Shifted-Gompertz distribution
// https://en.wikipedia.org/wiki/Shifted_Gompertz_distribution
type ShiftedGompertz struct {
	scale, shape float64 // b, η
	src          rand.Source
}

func NewShiftedGompertz(scale, shape float64) (*ShiftedGompertz, error) {
	return NewShiftedGompertzWithSource(scale, shape, nil)
}

func NewShiftedGompertzWithSource(scale, shape float64, src rand.Source) (*ShiftedGompertz, error) {
	if scale < 0 || shape < 0 {
		return nil, err.Invalid()
	}

	return &ShiftedGompertz{scale, shape, src}, nil
}

// η ∈ [0,∞)
// b ∈ [0,∞)
func (sg *ShiftedGompertz) Parameters() stats.Limits {
	return stats.Limits{
		"η": stats.Interval{0, math.Inf(1), true, true},
		"b": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [0,∞)
func (sg *ShiftedGompertz) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), false, true}
}

func (sg *ShiftedGompertz) Probability(x float64) float64 {
	if sg.Support().IsWithinInterval(x) {
		return sg.scale * math.Exp(-sg.scale*x) * math.Exp(-sg.shape*math.Exp(-sg.scale*x)) * (1 + sg.shape*(1-math.Exp(-sg.scale*x)))
	}

	return 0
}

func (sg *ShiftedGompertz) Distribution(x float64) float64 {
	if sg.Support().IsWithinInterval(x) {
		return (1 - math.Exp(-sg.scale*x)) * math.Exp(-sg.shape*math.Exp(-sg.scale*x))
	}

	return 0
}

func (sg *ShiftedGompertz) z() float64 {
	return (3 + sg.shape - math.Pow((sg.shape*sg.shape)+2*sg.shape+5, 1./2)) / (2 * sg.shape)
}

func (sg *ShiftedGompertz) Mean() float64 {
	return (1 - math.Exp(-sg.shape) + sg.shape*(gsl.Euler+specfunc.Gamma(sg.shape)+math.Log(sg.shape))) / (sg.scale * sg.shape)
}

func (sg *ShiftedGompertz) Median() float64 {
	return sg.Inverse(.5)
}

func (sg *ShiftedGompertz) Mode() float64 {
	if sg.shape > 0 && sg.shape <= 0.5 {
		return 0
	}

	if sg.shape > 0.5 {
		return (1 / sg.scale) * math.Log(sg.z())
	}

	return math.NaN()
}

func (sg *ShiftedGompertz) Inverse(p float64) float64 {
	if p >= 1 {
		return math.Inf(1)
	}

	if p <= 0 {
		return 0
	}

	return -(math.Log(1-(specfunc.Lambert_W0(math.Exp(sg.shape)*p*sg.shape)/sg.shape)) / sg.scale)
}

func (sg *ShiftedGompertz) Rand() float64 {
	var rnd float64
	if sg.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(sg.src).Float64()
	}

	return sg.Inverse(rnd)
}
