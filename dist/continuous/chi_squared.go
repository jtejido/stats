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

// Chi-Squared distribution
// https://en.wikipedia.org/wiki/Chi-squared_distribution
type ChiSquared struct {
	baseContinuousWithSource
	dof     int // degrees of freedom
	natural linear.RealVector
}

func NewChiSquared(dof int) (*ChiSquared, error) {
	return NewChiSquaredWithSource(dof, nil)
}

func NewChiSquaredWithSource(dof int, src rand.Source) (*ChiSquared, error) {
	if dof <= 0 {
		return nil, err.Invalid()
	}

	ret := new(ChiSquared)
	ret.dof = dof
	ret.src = src

	return ret, nil
}

func (cs *ChiSquared) String() string {
	return "ChiSquared: Parameters - " + cs.Parameters().String() + ", Support(x) - " + cs.Support().String()
}

// k ∈ (0,∞)
func (cs *ChiSquared) Parameters() stats.Limits {
	return stats.Limits{
		"k": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [0,∞)
func (cs *ChiSquared) Support() stats.Interval {
	if cs.dof == 1 {
		return stats.Interval{0, math.Inf(1), true, true}
	}

	return stats.Interval{0, math.Inf(1), false, true}
}

func (cs *ChiSquared) Probability(x float64) float64 {
	if cs.Support().IsWithinInterval(x) {
		xpowKhalfm1 := math.Pow(x, ((float64(cs.dof) / 2.) - 1.))
		expPowmxHalf := math.Exp(-(x / 2.))
		Twokhalf := math.Pow(2, (float64(cs.dof) / 2.))
		gammaKHalf := specfunc.Gamma(float64(cs.dof) / 2)

		return (xpowKhalfm1 * expPowmxHalf) / (Twokhalf * gammaKHalf)
	}

	return 0
}

func (cs *ChiSquared) Distribution(x float64) float64 {
	if cs.Support().IsWithinInterval(x) {
		return specfunc.Gamma_inc_P(float64(cs.dof)/2, x/2)
	}

	return 0
}

func (cs *ChiSquared) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return smath.InverseRegularizedLowerIncompleteGamma(float64(cs.dof)/2.0, p) * 2.0
}

func (cs *ChiSquared) Entropy() float64 {
	return (float64(cs.dof) / 2.) + math.Log(2*specfunc.Gamma(float64(cs.dof)/2)) + (1-(float64(cs.dof)/2.))*specfunc.Psi(float64(cs.dof)/2.)
}

func (cs *ChiSquared) ExKurtosis() float64 {
	return 12 / float64(cs.dof)
}

func (cs *ChiSquared) Skewness() float64 {
	return math.Sqrt(8. / float64(cs.dof))
}

func (cs *ChiSquared) Mean() float64 {
	return float64(cs.dof)
}

func (cs *ChiSquared) Median() float64 {
	return float64(cs.dof) * math.Pow(1-(2/(9*float64(cs.dof))), 3.)
}

func (cs *ChiSquared) Mode() float64 {
	return math.Max(float64(cs.dof)-2., 0.)
}

func (cs *ChiSquared) Variance() float64 {
	return 2 * float64(cs.dof)
}

func (cs *ChiSquared) Rand() float64 {
	var rnd float64
	if cs.src != nil {
		rnd = rand.New(cs.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return cs.Inverse(rnd)
}

func (cs *ChiSquared) ToExponential() {
	vec, _ := linear.NewArrayRealVectorFromSlice([]float64{(float64(cs.dof) / 2) - 1})
	cs.natural = vec
}

func (cs *ChiSquared) SufficientStatistics(x float64) linear.RealVector {
	vec, _ := linear.NewArrayRealVectorFromSlice([]float64{math.Log(x)})
	return vec
}
