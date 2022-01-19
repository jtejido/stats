package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/linear"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// Gamma distribution
// https://en.wikipedia.org/wiki/Gamma_distribution
type InverseGamma struct {
	shape, scale float64 // α, β
	src          rand.Source
	natural      linear.RealVector
}

func NewInverseGamma(shape, scale float64) (*InverseGamma, error) {
	return NewInverseGammaWithSource(shape, scale, nil)
}

func NewInverseGammaWithSource(shape, scale float64, src rand.Source) (*InverseGamma, error) {
	if shape <= 0 || scale <= 0 {
		return nil, err.Invalid()
	}

	return &InverseGamma{shape, scale, src, nil}, nil
}

// α ∈ (0,∞)
// β ∈ (0,∞)
func (ig *InverseGamma) Parameters() stats.Limits {
	return stats.Limits{
		"α": stats.Interval{0, math.Inf(1), true, true},
		"β": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (ig *InverseGamma) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

func (ig *InverseGamma) Probability(x float64) float64 {
	if ig.Support().IsWithinInterval(x) {
		return (math.Pow(ig.scale, ig.shape) / specfunc.Gamma(ig.shape)) * math.Pow(x, -ig.shape-1) * math.Exp(-ig.scale/x)
	}

	return 0
}

func (ig *InverseGamma) Distribution(x float64) float64 {
	if ig.Support().IsWithinInterval(x) {
		return specfunc.Gamma_inc_Q(ig.shape, ig.scale/x)
	}

	return 0
}

func (ig *InverseGamma) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return (1 / smath.InverseRegularizedLowerIncompleteGamma(ig.shape, p)) * ig.scale
}

func (ig *InverseGamma) Entropy() float64 {
	return ig.shape + math.Log(ig.scale*specfunc.Gamma(ig.shape)) - (1+ig.shape)*specfunc.Psi(ig.shape)
}

func (ig *InverseGamma) ExKurtosis() float64 {
	return (6 * (5*ig.shape - 11)) / ((ig.shape - 3) * (ig.shape - 4))
}

func (ig *InverseGamma) Skewness() float64 {
	return (4 * math.Sqrt(ig.shape-2)) / (ig.shape - 3)
}

func (ig *InverseGamma) Mean() float64 {
	return ig.scale / (ig.shape - 1)
}
func (ig *InverseGamma) Median() float64 {
	return ig.scale / smath.InverseRegularizedLowerIncompleteGamma(ig.shape, .5)
}

func (ig *InverseGamma) Mode() float64 {
	return ig.scale / (ig.shape + 1)
}

func (ig *InverseGamma) Variance() float64 {
	return (ig.scale * ig.scale) / (math.Pow(ig.shape-1, 2) * (ig.shape - 2))
}

func (ig *InverseGamma) Rand() float64 {
	var rnd float64
	if ig.src != nil {
		rnd = rand.New(ig.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return ig.Inverse(rnd)
}

func (ig *InverseGamma) ToExponential() {
	vec, _ := linear.NewArrayRealVectorFromSlice([]float64{-ig.shape - 1, -ig.scale})
	ig.natural = vec

	// vec2, _ := linear.NewSizedArrayRealVector(2)
	// vec2.SetEntry(0, math.Log(ig.scale)-specfunc.Psi(ig.shape))
	// vec2.SetEntry(1, ig.shape/ig.scale)
	// ig.Moment = vec2
}

func (ig *InverseGamma) SufficientStatistics(x float64) linear.RealVector {
	vec, _ := linear.NewArrayRealVectorFromSlice([]float64{math.Log(x), 1 / x})
	return vec
}
