package continuous

import (
	// gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Gompertz distribution
// 蓑谷千凰彦：『統計分布ハンドブック　増補版』,朝倉書店,2010,pp.707-713.
// https://en.wikipedia.org/wiki/Gompertz_distribution
type Gompertz struct {
	baseContinuousWithSource
	shape, scale float64 // η, b
}

func NewGompertz(shape, scale float64) (*Gompertz, error) {
	return NewGompertzWithSource(shape, scale, nil)
}

func NewGompertzWithSource(shape, scale float64, src rand.Source) (*Gompertz, error) {
	if shape <= 0 || scale <= 0 {
		return nil, err.Invalid()
	}

	g := new(Gompertz)
	g.shape = shape
	g.scale = scale
	g.src = src

	return g, nil
}

func (g *Gompertz) String() string {
	return "Gompertz: Parameters - " + g.Parameters().String() + ", Support(x) - " + g.Support().String()
}

// η ∈ (0,∞)
// b ∈ (0,∞)
func (g *Gompertz) Parameters() stats.Limits {
	return stats.Limits{
		"η": stats.Interval{0, math.Inf(1), true, true},
		"b": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [0,∞)
func (g *Gompertz) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), false, true}
}

func (g *Gompertz) Probability(x float64) float64 {
	if g.Support().IsWithinInterval(x) {
		return g.shape * math.Exp(g.scale*x) * math.Exp(-g.shape/g.scale*(math.Expm1(g.scale*x)))
	}

	return 0
}

func (g *Gompertz) Distribution(x float64) float64 {
	if g.Support().IsWithinInterval(x) {
		return 1 - math.Exp(-g.shape/g.scale*(math.Expm1(g.scale*x)))
	}

	return 0
}

func (g *Gompertz) Mean() float64 {
	return (1 / g.scale) * math.Exp(g.shape) * specfunc.Expint_Ei(-g.shape)
}

func (g *Gompertz) Median() float64 {
	return (1 / g.scale) * math.Log((-1/g.shape)*math.Log(1./2)+1)
}

func (g *Gompertz) Mode() float64 {
	if g.shape >= 1 {
		return 0
	}

	return (1 / g.scale) * math.Log(1/g.shape)
}

// func (g *Gompertz) Variance() float64 {
// 	threef3 := stats.Hyperg_3F3(1, 1, 1, 2, 2, 2, -g.shape)
// 	return ((1 / g.scale) * (1 / g.scale)) * math.Exp(g.shape) * (-2*g.shape*threef3 + math.Pow(gsl.Euler, 2.) + (math.Pow(math.Pi, 2.) / 6) + 2*gsl.Euler*math.Log(g.shape) + math.Pow(math.Log(g.shape), 2.) - math.Exp(g.shape)*math.Pow(specfunc.Expint_Ei(-g.shape), 2.))
// }

func (g *Gompertz) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return math.Log(1.0-g.scale/g.shape*math.Log(1.0-p)) / g.scale
}

func (g *Gompertz) Rand() float64 {
	var rnd float64
	if g.src != nil {
		rnd = rand.New(g.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return g.Inverse(rnd)
}
