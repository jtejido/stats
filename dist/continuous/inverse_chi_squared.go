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

// (Scaled) Inverse chi-squared distribution
// https://en.wikipedia.org/wiki/Scaled_inverse_chi-squared_distribution
// https://en.wikipedia.org/wiki/Inverse-chi-squared_distribution
type InverseChiSquared struct {
	dof, scale float64 // v, σ2
	src        rand.Source
	natural    linear.RealVector
}

func NewInverseChiSquared(dof, scale float64) (*InverseChiSquared, error) {
	return NewInverseChiSquaredWithSource(dof, scale, nil)
}

func NewInverseChiSquaredWithSource(dof, scale float64, src rand.Source) (*InverseChiSquared, error) {
	if dof <= 0 || scale <= 0 {
		return nil, err.Invalid()
	}

	return &InverseChiSquared{dof, scale, src, nil}, nil
}

// v ∈ (0,∞)
// σ2 ∈ (0,∞)
func (i *InverseChiSquared) Parameters() stats.Limits {
	return stats.Limits{
		"v":  stats.Interval{0, math.Inf(1), true, true},
		"σ2": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (i *InverseChiSquared) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

func (i *InverseChiSquared) Probability(x float64) float64 {
	if i.Support().IsWithinInterval(x) {
		return (math.Pow(2, -i.dof/2) * math.Exp(-(i.dof*i.scale)/(2*x)) * math.Pow((i.dof*i.scale)/x, i.dof/2)) / (x * specfunc.Gamma(i.dof/2))
	}

	return 0
}

func (i *InverseChiSquared) Distribution(x float64) float64 {
	if i.Support().IsWithinInterval(x) {
		return specfunc.Gamma_inc_Q(i.dof/2, (i.scale*i.dof)/(2*x))
	}

	return 0
}

func (i *InverseChiSquared) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return (i.dof * i.scale) / (2 * smath.InverseRegularizedLowerIncompleteGamma(i.dof/2, p))
}

func (i *InverseChiSquared) Entropy() float64 {
	return (i.dof / 2) + math.Log(((i.scale*i.dof)/2)*specfunc.Gamma(i.dof/2)) - (1+(i.dof/2))*specfunc.Psi(i.dof/2)
}

func (i *InverseChiSquared) ExKurtosis() float64 {
	if i.dof > 8 {
		return (12 * (5*i.dof - 22)) / ((i.dof - 6) * (i.dof - 8))
	}

	return math.Inf(1)
}

func (i *InverseChiSquared) Mean() float64 {
	if i.dof > 2 {
		return (i.dof * i.scale) / (i.dof - 2)
	}

	return math.Inf(1)
}

func (i *InverseChiSquared) Median() float64 {
	return (i.dof * i.scale) / (2 * smath.InverseRegularizedLowerIncompleteGamma(i.dof/2, .5))
}

func (i *InverseChiSquared) Mode() float64 {
	return (i.dof * i.scale) / (i.dof + 2)
}

func (i *InverseChiSquared) Skewness() float64 {
	if i.dof > 6 {
		return (4 * math.Sqrt(2) * math.Sqrt(-4+i.dof)) / (-6 + i.dof)
	}

	return math.Inf(1)
}

func (i *InverseChiSquared) Variance() float64 {
	if i.dof > 4 {
		return (2 * (i.dof * i.dof) * (i.scale * i.scale)) / (math.Pow(i.dof-2, 2) * (i.dof - 4))
	}

	return math.Inf(1)
}

func (i *InverseChiSquared) Rand() float64 {
	var rnd float64
	if i.src != nil {
		rnd = rand.New(i.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return i.Inverse(rnd)
}

func (i *InverseChiSquared) ToExponential() {
	vec, _ := linear.NewArrayRealVectorFromSlice([]float64{-(i.dof / 2) - 1, -(i.dof * i.scale) / 2})
	i.natural = vec
	// n1 := vec.At(0)
	// n2 := vec.At(1)
	// vec2, _ := linear.NewSizedArrayRealVector(2)
	// vec2.SetEntry(0, math.Log((i.dof * i.scale) / 2)-specfunc.Psi((i.dof / 2) + 1))
	// vec2.SetEntry(1, -(i.dof + 4)/(i.dof * i.scale))
	// i.Moment = vec2
}

func (i *InverseChiSquared) SufficientStatistics(x float64) linear.RealVector {
	vec, _ := linear.NewArrayRealVectorFromSlice([]float64{math.Log(x), 1 / x})
	return vec
}
