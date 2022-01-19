package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Assymetric Laplace distribution
// https://en.wikipedia.org/wiki/Asymmetric_Laplace_distribution
type AssymetricLaplace struct {
	baseContinuousWithSource
	location, scale, assymetry float64 // m, λ, κ
}

func NewAssymetricLaplace(m, λ, κ float64) (*AssymetricLaplace, error) {
	return NewAssymetricLaplaceWithSource(m, λ, κ, nil)
}

func NewAssymetricLaplaceWithSource(m, λ, κ float64, src rand.Source) (*AssymetricLaplace, error) {
	if λ <= 0 || κ <= 0 {
		return nil, err.Invalid()
	}

	ret := new(AssymetricLaplace)
	ret.location = m
	ret.scale = λ
	ret.assymetry = κ
	ret.src = src
	return ret, nil
}

func (al *AssymetricLaplace) String() string {
	return "AssymetricLaplace: Parameters - " + al.Parameters().String() + ", Support(x) - " + al.Support().String()
}

func (al *AssymetricLaplace) Parameters() stats.Limits {
	return stats.Limits{
		"m": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"λ": stats.Interval{0, math.Inf(1), true, true},
		"κ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (-∞,∞)
func (al *AssymetricLaplace) Support() stats.Interval {
	return stats.Interval{math.Inf(-1), math.Inf(1), true, true}
}

func (al *AssymetricLaplace) Probability(x float64) float64 {
	a := al.scale / (al.assymetry + 1/al.assymetry)
	if x < al.location {
		return a * math.Exp((al.scale/al.assymetry)*(x-al.location))
	}

	return a * math.Exp(-al.scale*al.assymetry*(x-al.location))
}

func (al *AssymetricLaplace) Distribution(x float64) float64 {
	if x <= al.location {
		return ((al.assymetry * al.assymetry) / (1 + (al.assymetry * al.assymetry))) * math.Exp((al.scale/al.assymetry)*(x-al.location))
	}

	return 1 - (1/(1+(al.assymetry*al.assymetry)))*math.Exp(-al.scale*al.assymetry*(x-al.location))
}

func (al *AssymetricLaplace) Inverse(p float64) float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (al *AssymetricLaplace) ExKurtosis() float64 {
	k4 := al.assymetry * al.assymetry * al.assymetry * al.assymetry
	k8 := k4 * al.assymetry * al.assymetry * al.assymetry * al.assymetry
	return (6 * (1 + k8)) / math.Pow(1+k4, 2)
}

func (al *AssymetricLaplace) Skewness() float64 {
	k4 := al.assymetry * al.assymetry * al.assymetry * al.assymetry
	k6 := k4 * al.assymetry * al.assymetry
	return (2 * (1 - k6)) / math.Pow(k4+1, 3./2)
}

func (al *AssymetricLaplace) Mean() float64 {
	return al.location + ((1 - (al.assymetry * al.assymetry)) / (al.scale * al.assymetry))
}

func (al *AssymetricLaplace) Median() float64 {
	return al.location + (al.assymetry/al.scale)*math.Log((1+(al.assymetry*al.assymetry))/(2*(al.assymetry*al.assymetry)))
}

func (al *AssymetricLaplace) Variance() float64 {
	s2 := al.scale * al.scale
	k2 := al.assymetry * al.assymetry
	k4 := k2 * al.assymetry * al.assymetry * al.assymetry * al.assymetry
	return (1 + k4) / (s2 * k2)
}

func (al *AssymetricLaplace) Entropy() float64 {
	k2 := al.assymetry * al.assymetry
	return math.Log(math.Exp((1 + k2) / (al.assymetry * al.scale)))
}

func (al *AssymetricLaplace) Rand() float64 {
	var rnd float64
	if al.src != nil {
		rnd = rand.New(al.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	s := gsl.Sign(rnd)
	return al.location - (1/(al.scale*s*math.Pow(al.assymetry, s)))*math.Log(1-rnd*s*math.Pow(al.assymetry, s))
}
