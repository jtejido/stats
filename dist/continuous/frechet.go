package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Frechet distribution
// Generalized Extreme Value distribution Type-II
// https://en.wikipedia.org/wiki/Fr%C3%A9chet_distribution
type Frechet struct {
	baseContinuousWithSource
	shape, scale, location float64 // α, s, m
}

func NewFrechet(shape, scale, location float64) (*Frechet, error) {
	return NewFrechetWithSource(shape, scale, location, nil)
}

func NewFrechetWithSource(shape, scale, location float64, src rand.Source) (*Frechet, error) {
	if shape >= 0 || scale >= 0 {
		return nil, err.Invalid()
	}

	r := new(Frechet)
	r.shape = shape
	r.scale = scale
	r.location = location
	r.src = src

	return r, nil
}

func (f *Frechet) String() string {
	return "Frechet: Parameters - " + f.Parameters().String() + ", Support(x) - " + f.Support().String()
}

// α ∈ (0,∞)
// s ∈ (0,∞)
// m ∈ (-∞,∞)
func (f *Frechet) Parameters() stats.Limits {
	return stats.Limits{
		"α": stats.Interval{0, math.Inf(1), true, true},
		"s": stats.Interval{0, math.Inf(1), true, true},
		"m": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
	}
}

// x ∈ [m,∞)
func (f *Frechet) Support() stats.Interval {
	return stats.Interval{f.location, math.Inf(1), false, true}
}

func (f *Frechet) Probability(x float64) float64 {
	if f.Support().IsWithinInterval(x) {
		return (f.shape / f.scale) * math.Pow((x-f.location)/f.scale, -1-f.shape) * math.Exp(-math.Pow((x-f.location)/f.scale, -f.shape))
	}

	return 0
}

func (f *Frechet) Distribution(x float64) float64 {
	if f.Support().IsWithinInterval(x) {
		return math.Exp(-math.Pow((x-f.location)/f.scale, -f.shape))
	}

	return 0
}

func (f *Frechet) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return f.location + f.scale*math.Pow(-math.Log(p), -1/f.shape)
}

func (f *Frechet) Entropy() float64 {
	return 1 + (gsl.Euler / f.shape) + gsl.Euler + math.Log(f.scale/f.shape)
}

func (f *Frechet) ExKurtosis() float64 {
	if f.shape > 4 {
		return -6 + ((specfunc.Gamma(1-(4/f.shape)) - 4*specfunc.Gamma(1-(3/f.shape))*specfunc.Gamma(1-(1/f.shape)) + 3*math.Pow(specfunc.Gamma(1-(2/f.shape)), 2.)) / math.Pow(specfunc.Gamma(1-(2/f.shape))-math.Pow(specfunc.Gamma(1-(1/f.shape)), 2.), 2.))
	}

	return math.Inf(1)
}

func (f *Frechet) Mean() float64 {
	if f.shape > 1 {
		return f.location + f.scale*specfunc.Gamma(1-(1/f.shape))
	}

	return math.Inf(1)
}

func (f *Frechet) Median() float64 {
	return f.location + f.scale/(math.Pow(math.Log(2), 1/f.shape))
}

func (f *Frechet) Mode() float64 {
	return f.location + f.scale*math.Pow(f.shape/(1+f.shape), 1/f.shape)
}

func (f *Frechet) Skewness() float64 {
	if f.shape > 3 {
		return (specfunc.Gamma(1-(3/f.shape)) - 3*specfunc.Gamma(1-(2/f.shape))*specfunc.Gamma(1-(1/f.shape)) + 2*math.Pow(specfunc.Gamma(1-(1/f.shape)), 3.)) / (math.Sqrt(math.Pow(specfunc.Gamma(1-(2/f.shape))-math.Pow(specfunc.Gamma(1-(1/f.shape)), 2), 3)))
	}

	return math.Inf(1)
}

func (f *Frechet) Variance() float64 {
	if f.shape > 2 {
		return (f.scale * f.scale) * (specfunc.Gamma(1-2/f.shape) - math.Pow(specfunc.Gamma(1-1/f.shape), 2.))
	}

	return math.Inf(1)
}

func (f *Frechet) Rand() float64 {
	var rnd float64
	if f.src != nil {
		rnd = rand.New(f.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return f.Inverse(rnd)
}
