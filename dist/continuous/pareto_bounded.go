package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Pareto distribution
// https://en.wikipedia.org/wiki/Pareto_distribution#Bounded_Pareto_distribution
type ParetoBounded struct {
	min, max, shape float64 // L, H, α
	src             rand.Source
}

func NewParetoBounded(min, max, shape float64) (*ParetoBounded, error) {
	return NewParetoBoundedWithSource(min, max, shape, nil)
}

func NewParetoBoundedWithSource(min, max, shape float64, src rand.Source) (*ParetoBounded, error) {
	if min >= max || shape <= 0 {
		return nil, err.Invalid()
	}

	return &ParetoBounded{min, max, shape, src}, nil
}

// L ∈ (0,∞)
// H ∈ (L,∞)
// α ∈ (0,∞)
func (p ParetoBounded) Parameters() stats.Limits {
	return stats.Limits{
		"L": stats.Interval{0, math.Inf(1), true, true},
		"H": stats.Interval{p.min, math.Inf(1), true, true},
		"α": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (p ParetoBounded) Support() stats.Interval {
	return stats.Interval{p.min, p.max, false, false}
}

func (p ParetoBounded) Probability(x float64) float64 {
	if p.Support().IsWithinInterval(x) {
		a := p.shape
		num := a * math.Pow(p.min, a) * math.Pow(x, -a-1)
		denom := 1 - math.Pow(p.min/p.max, a)
		return num / denom
	}

	return 0
}

func (p ParetoBounded) Distribution(x float64) float64 {
	if p.Support().IsWithinInterval(x) {
		a := p.shape
		num := 1 - math.Pow(p.min, a)*math.Pow(x, -a)
		denom := 1 - math.Pow(p.min/p.max, a)
		return num / denom
	}

	return 0
}

func (p ParetoBounded) Entropy() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (p ParetoBounded) Inverse(q float64) float64 {
	if q > 0 && q < 1 {
		num := -q*math.Pow(p.max, p.shape) - q*math.Pow(p.min, p.shape) - math.Pow(p.max, p.shape)
		denom := math.Pow(p.max, p.shape) * math.Pow(p.min, p.shape)
		return math.Pow(-(num / denom), -1/p.shape)
	}

	if q <= 0 {
		return 0
	}

	return math.Inf(1)

}

func (p ParetoBounded) Mean() float64 {
	// p.rm(1)
	if p.shape == 1 {
		return (p.max * p.min) / (p.max - p.min) * math.Log(p.max/p.min)
	}

	a := math.Pow(p.min, p.shape) / (1 - math.Pow(p.min/p.max, p.shape))
	b := p.shape / (p.shape - 1)
	c := (1 / math.Pow(p.min, p.shape-1)) - (1 / math.Pow(p.max, p.shape-1))
	return a * b * c
}

func (p ParetoBounded) Median() float64 {
	return p.min * math.Pow(1-(.5*(1-math.Pow(p.min/p.max, p.shape))), -1/p.shape)
}

func (p ParetoBounded) Mode() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (p ParetoBounded) Variance() float64 {
	m1 := p.rm(1)
	return -(m1 * m1) + p.rm(2)
}

func (p ParetoBounded) Skewness() float64 {
	m1 := p.rm(1)
	m2 := p.rm(2)
	m3 := p.rm(3)
	return (m1*(2*(m1*m1)-3*m2) + m3) / math.Pow(m2-(m1*m1), 3./2)
}

func (p ParetoBounded) ExKurtosis() float64 {
	m1 := p.rm(1)
	m2 := p.rm(2)
	m3 := p.rm(3)
	m4 := p.rm(4)
	return (-3*(m1*m1*m1*m1) + 6*(m1*m1)*m2 - 4*m1*m3 + m4 - 3*((m2-(m1*m1))*(m2-(m1*m1)))) / ((m2 - (m1 * m1)) * (m2 - (m1 * m1)))
}

func (p ParetoBounded) rm(k float64) float64 {
	a := math.Pow(p.min, p.shape) / (1 - math.Pow(p.min/p.max, p.shape))
	b := (p.shape * (math.Pow(p.min, k-p.shape) - math.Pow(p.max, k-p.shape))) / (p.shape - k)
	return a * b
}

func (p ParetoBounded) Rand() float64 {
	var rnd float64
	if p.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(p.src).Float64()
	}

	return p.Inverse(rnd)
}
