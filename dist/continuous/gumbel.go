package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Gumbel distribution
// https://en.wikipedia.org/wiki/Gumbel_distribution
type Gumbel struct {
	location, scale float64 // mu, beta
	src             rand.Source
}

func NewGumbel(location, scale float64) (*Gumbel, error) {
	return NewGumbelWithSource(location, scale, nil)
}

func NewGumbelWithSource(location, scale float64, src rand.Source) (*Gumbel, error) {
	if scale <= 0 {
		return nil, err.Invalid()
	}

	return &Gumbel{location, scale, src}, nil
}

// μ ∈ (-∞,∞)
// β ∈ (0,∞)
func (g *Gumbel) Parameters() stats.Limits {
	return stats.Limits{
		"μ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"β": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (-∞,∞)
func (g *Gumbel) Support() stats.Interval {
	return stats.Interval{math.Inf(-1), math.Inf(1), true, true}
}

func (g *Gumbel) z(x float64) float64 {
	return (x - g.location) / g.scale
}

func (g *Gumbel) Probability(x float64) float64 {
	z := g.z(x)
	return (1 / g.scale) * math.Exp(-(z + math.Exp(-z)))
}

func (g *Gumbel) Distribution(x float64) float64 {
	z := g.z(x)
	return math.Exp(-math.Exp(-z))
}

func (g *Gumbel) Entropy() float64 {
	return math.Log(g.scale) + gsl.Euler + 1
}

func (g *Gumbel) ExKurtosis() float64 {
	return 12.0 / 5
}

func (g *Gumbel) Mean() float64 {
	return g.location + g.scale*gsl.Euler
}

func (g *Gumbel) Median() float64 {
	return g.location - g.scale*math.Log(math.Ln2)
}

func (g *Gumbel) Mode() float64 {
	return g.location
}

func (g *Gumbel) Inverse(p float64) float64 {
	if p <= 0 {
		return math.Inf(-1)
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return g.location - g.scale*math.Log(-math.Log(p))
}

func (g *Gumbel) Skewness() float64 {
	return 12 * math.Sqrt(6) * gsl.Apery / (math.Pi * math.Pi * math.Pi)
}

func (g *Gumbel) Variance() float64 {
	return ((math.Pi * math.Pi) / 6) * (g.scale * g.scale)
}

func (g *Gumbel) Rand() float64 {
	var rnd float64
	if g.src != nil {
		rnd = rand.New(g.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return g.Inverse(rnd)
}
