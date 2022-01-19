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
type Gamma struct {
	baseContinuousWithSource
	shape, rate float64 // α, β
	natural     linear.RealVector
}

func NewGamma(shape, rate float64) (*Gamma, error) {
	return NewGammaWithSource(shape, rate, nil)
}

func NewGammaWithSource(shape, rate float64, src rand.Source) (*Gamma, error) {
	if shape <= 0 || rate <= 0 {
		return nil, err.Invalid()
	}

	r := new(Gamma)
	r.shape = shape
	r.rate = rate
	r.src = src

	return r, nil
}

func (g *Gamma) String() string {
	return "Gamma: Parameters - " + g.Parameters().String() + ", Support(x) - " + g.Support().String()
}

// k ∈ (0,∞)
// θ ∈ (0,∞)
func (g *Gamma) Parameters() stats.Limits {
	return stats.Limits{
		"k": stats.Interval{0, math.Inf(1), true, true},
		"θ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (g *Gamma) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

func (g *Gamma) Probability(x float64) float64 {
	if g.Support().IsWithinInterval(x) {
		return (math.Pow(g.rate, g.shape) / specfunc.Gamma(g.shape)) * math.Pow(x, g.shape-1) * math.Exp(-g.rate*x)
	}

	return 0

}

func (g *Gamma) Distribution(x float64) float64 {
	if g.Support().IsWithinInterval(x) {
		return specfunc.Gamma_inc_P(g.shape, g.rate*x)
	}

	return 0
}

func (g *Gamma) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return smath.InverseRegularizedLowerIncompleteGamma(g.shape, p) / g.rate
}

func (g *Gamma) Entropy() float64 {
	return g.shape - math.Log(g.rate) + specfunc.Lngamma(g.shape) + (1-g.shape)*specfunc.Psi(g.shape)
}

func (g *Gamma) ExKurtosis() float64 {
	return 6 / g.shape
}

func (g *Gamma) Skewness() float64 {
	return 2 / math.Sqrt(g.shape)
}

func (g *Gamma) Mean() float64 {
	return g.shape / g.rate
}
func (g *Gamma) Median() float64 {
	μ := g.Mean()
	k3 := 3 * g.shape

	return μ * ((k3 - 0.8) / (k3 + 0.2))
}

func (g *Gamma) Mode() float64 {
	if g.shape < 1 {
		return math.NaN()
	}

	return (g.shape - 1) / g.rate
}

func (g *Gamma) Variance() float64 {
	return g.shape / (g.rate * g.rate)
}

func (g *Gamma) Rand() float64 {
	var d, c float64
	if g.shape < 1 {
		d = g.shape + 1.0 - 1.0/3.0
		c = (1.0 / 3.0) / math.Sqrt(d)
		var v float64

		if g.src != nil {
			u := rand.New(g.src)
			v = u.Float64()
		} else {
			v = rand.Float64()
		}

		return (marsaglia(g.src, d, c) * math.Pow(v, 1.0/g.shape)) / g.rate
	} else {
		d = g.shape - 1.0/3.0
		c = (1.0 / 3.0) / math.Sqrt(d)

		return marsaglia(g.src, d, c) / g.rate
	}
}

func marsaglia(src rand.Source, d, c float64) float64 {
	for {
		var x, t, v float64

		for {
			n, _ := NewNormalWithSource(0, 1, src)
			x = n.Rand()
			t = (1.0 + c*x)
			v = t * t * t
			if v > 0 {
				break
			}
		}

		u := rand.Float64()

		x2 := x * x
		if u < 1-0.0331*x2*x2 {
			return d * v
		}

		if math.Log(u) < 0.5*x2+d*(1.0-v+math.Log(v)) {
			return d * v
		}

	}
}

func (g *Gamma) ToExponential() {
	vec, _ := linear.NewArrayRealVectorFromSlice([]float64{g.shape - 1, -g.rate})
	g.natural = vec

	// vec2, _ := linear.NewSizedArrayRealVector(2)
	// vec2.SetEntry(0, specfunc.Psi(g.shape)+math.Log(g.scale))
	// vec2.SetEntry(1, g.shape*g.scale)
	// g.Moment = vec2
}

func (g *Gamma) SufficientStatistics(x float64) linear.RealVector {
	vec, _ := linear.NewArrayRealVectorFromSlice([]float64{math.Log(x), x})
	return vec
}
