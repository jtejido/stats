package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// Chi distribution
// https://en.wikipedia.org/wiki/Chi_distribution
type Chi struct {
	baseContinuousWithSource
	dof int // degrees of freedom
}

func NewChi(dof int) (*Chi, error) {
	return NewChiWithSource(dof, nil)
}

func NewChiWithSource(dof int, src rand.Source) (*Chi, error) {
	if dof <= 0 {
		return nil, err.Invalid()
	}

	ret := new(Chi)
	ret.dof = dof
	ret.src = src

	return ret, nil
}

func (c *Chi) String() string {
	return "Chi: Parameters - " + c.Parameters().String() + ", Support(x) - " + c.Support().String()
}

// k ∈ (0,∞)
func (c *Chi) Parameters() stats.Limits {
	return stats.Limits{
		"k": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [0,∞)
func (c *Chi) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), false, true}
}

func (c *Chi) Probability(x float64) float64 {
	if x > 0 {
		return (math.Pow(x, float64(c.dof-1)) * math.Exp(-(x*x)/2)) / (math.Pow(2, ((float64(c.dof)/2.)-1.)) * specfunc.Gamma(float64(c.dof)/2.))
	}

	return 0

}

func (c *Chi) Distribution(x float64) float64 {
	if c.Support().IsWithinInterval(x) {
		return specfunc.Gamma_inc_P(float64(c.dof)/2, (x*x)/2)
	}

	return 0
}

func (c *Chi) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return math.Sqrt(smath.InverseRegularizedLowerIncompleteGamma(float64(c.dof)/2.0, p)) * math.Sqrt(2)

}

func (c *Chi) Entropy() float64 {
	return math.Log(specfunc.Gamma(float64(c.dof)/2.0)) + 0.5*(float64(c.dof)-gsl.Ln2-(float64(c.dof)-1.0)*specfunc.Psi(float64(c.dof)/2.0))
}

func (c *Chi) Mode() float64 {
	return math.Max(math.Sqrt(float64(c.dof)-1), 0.)
}

func (c *Chi) Median() float64 {
	return math.Sqrt(float64(c.dof) * math.Pow(1-(2/(9*float64(c.dof))), 3.))
}

// these moments should be expressed from raw moment
func (c *Chi) ExKurtosis() float64 {
	variance := c.Variance()
	mean := c.Mean()
	skew := c.Skewness()
	sd := math.Sqrt(variance)
	return (2 / variance) * (1 - mean*sd*skew - variance)
}

func (c *Chi) Skewness() float64 {
	variance := c.Variance()
	sd := math.Sqrt(variance)
	mean := c.Mean()
	return (mean / (sd * sd * sd)) * (1 - 2*variance)
}

func (c *Chi) Mean() float64 {
	return math.Sqrt(2) * (specfunc.Gamma((float64(c.dof)+1)/2) / specfunc.Gamma(float64(c.dof)/2))
}

func (c *Chi) Variance() float64 {
	mean := c.Mean()
	return float64(c.dof) * (mean * mean)
}

func (c *Chi) Rand() float64 {
	var rnd float64
	if c.src != nil {
		rnd = rand.New(c.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return c.Inverse(rnd)
}
