package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// Nakagami distribution
// https://en.wikipedia.org/wiki/Nakagami_distribution
type Nakagami struct {
	shape, spread float64 // m, Ω
	src           rand.Source
}

func NewNakagami(shape, spread float64) (*Nakagami, error) {
	return NewNakagamiWithSource(shape, spread, nil)
}

func NewNakagamiWithSource(shape, spread float64, src rand.Source) (*Nakagami, error) {
	if shape < 0.5 || spread <= 0 {
		return nil, err.Invalid()
	}

	return &Nakagami{shape, spread, src}, nil
}

// m ∈ [0.5,∞)
// Ω ∈ (0,∞)
func (n *Nakagami) Parameters() stats.Limits {
	return stats.Limits{
		"m": stats.Interval{0.5, math.Inf(1), false, true},
		"Ω": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (n *Nakagami) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

func (n *Nakagami) Probability(x float64) float64 {
	if n.Support().IsWithinInterval(x) {
		return ((2 * math.Pow(n.shape, n.shape)) / (specfunc.Gamma(n.shape) * math.Pow(n.spread, n.shape))) * math.Pow(x, 2*n.shape-1) * math.Exp(-n.shape/n.spread*(x*x))
	}

	return 0
}

func (n *Nakagami) Distribution(x float64) float64 {
	if n.Support().IsWithinInterval(x) {
		return specfunc.Gamma_inc_P(n.shape, (n.shape*(x*x))/n.spread)
	}

	return 0
}

func (n *Nakagami) Mean() float64 {
	return (math.Gamma(n.shape+1./2) / specfunc.Gamma(n.shape)) * math.Pow(n.spread/n.shape, 1./2)
}

func (n *Nakagami) Mode() float64 {
	return math.Sqrt(2) / 2 * math.Pow((((2*n.shape)-1)*n.spread)/n.shape, 1./2)
}

func (n *Nakagami) Variance() float64 {
	return n.spread * (1 - 1./n.shape*math.Pow(specfunc.Gamma(n.shape+1./2)/specfunc.Gamma(n.shape), 2.))
}

func (n *Nakagami) Skewness() float64 {
	return (specfunc.Poch(n.shape, .5) * (.5 - 2*(n.shape-math.Pow(specfunc.Poch(n.shape, .5), 2)))) / math.Pow(n.shape-math.Pow(specfunc.Poch(n.shape, .5), 2), 3./2)
}

func (n *Nakagami) ExKurtosis() float64 {
	num := n.shape*(1+4*n.shape) - 2*(1+2*n.shape)*math.Pow(specfunc.Poch(n.shape, .5), 2)
	denom := math.Pow(n.shape-math.Pow(specfunc.Poch(n.shape, .5), 2), 2)
	return (num / denom) - 6
}

func (n *Nakagami) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return math.Sqrt((n.spread * smath.InverseRegularizedLowerIncompleteGamma(n.shape, p)) / n.shape)
}

func (n *Nakagami) Rand() float64 {
	var g Gamma
	g.shape = n.shape
	g.rate = n.spread / n.shape
	g.src = n.src

	return math.Sqrt(g.Rand())
}
