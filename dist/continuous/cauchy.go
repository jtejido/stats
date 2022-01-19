package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Cauchy distribution
// https://en.wikipedia.org/wiki/Cauchy_distribution
type Cauchy struct {
	baseContinuousWithSource
	location, scale float64 // x₀, γ
}

func NewCauchy(location, scale float64) (*Cauchy, error) {
	return NewCauchyWithSource(location, scale, nil)
}

func NewCauchyWithSource(location, scale float64, src rand.Source) (*Cauchy, error) {
	if scale <= 0 {
		return nil, err.Invalid()
	}

	ret := new(Cauchy)
	ret.location = location
	ret.scale = scale
	ret.src = src

	return ret, nil
}

func (c *Cauchy) String() string {
	return "Cauchy: Parameters - " + c.Parameters().String() + ", Support(x) - " + c.Support().String()
}

// x₀ ∈ (-∞,∞)
// γ  ∈ (0,∞)
func (c *Cauchy) Parameters() stats.Limits {
	return stats.Limits{
		"x₀": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"γ":  stats.Interval{0, math.Inf(1), true, true},
	}
}

// x  ∈ (-∞,∞)
func (c *Cauchy) Support() stats.Interval {
	return stats.Interval{math.Inf(-1), math.Inf(1), true, true}
}

func (c *Cauchy) Probability(x float64) float64 {
	return 1 / (math.Pi * c.scale * (1. + math.Pow(((x-c.location)/c.scale), 2.)))
}

func (c *Cauchy) Distribution(x float64) float64 {
	return 1/math.Pi*math.Atan((x-c.location)/c.scale) + 0.5
}

func (c *Cauchy) Entropy() float64 {
	return math.Log(4 * math.Pi * c.scale)
}

func (c *Cauchy) ExKurtosis() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (c *Cauchy) Skewness() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (c *Cauchy) Inverse(p float64) float64 {
	if p <= 0 {
		return math.Inf(-1)
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return c.location + c.scale*math.Tan(math.Pi*(p-0.5))
}

func (c *Cauchy) Mean() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (c *Cauchy) Median() float64 {
	return c.location
}

func (c *Cauchy) Mode() float64 {
	return c.location
}

func (c *Cauchy) Variance() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (c *Cauchy) Rand() float64 {
	var rnd func() float64
	if c.src == nil {
		rnd = rand.Float64
	} else {
		rnd = rand.New(c.src).Float64
	}

	var x, y float64
	for y == 0.0 || x*x+y*y > 1.0 {
		x = 2*rnd() - 1
		y = 2*rnd() - 1
	}

	return c.location + c.scale*(x/y)
}
