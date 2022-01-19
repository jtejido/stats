package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Birnbaum–Saunders distribution (fatigue life distribution)
// https://en.wikipedia.org/wiki/Birnbaum%E2%80%93Saunders_distribution
type BirnbaumSaunders struct {
	baseContinuousWithSource
	shape, scale float64 // α, β
}

func NewBirnbaumSaunders(shape, scale float64) (*BirnbaumSaunders, error) {
	return NewBirnbaumSaundersWithSource(shape, scale, nil)
}

func NewBirnbaumSaundersWithSource(shape, scale float64, src rand.Source) (*BirnbaumSaunders, error) {
	if shape <= 0 || scale <= 0 {
		return nil, err.Invalid()
	}

	ret := new(BirnbaumSaunders)
	ret.shape = shape
	ret.scale = scale
	ret.src = src

	return ret, nil
}

func (bs *BirnbaumSaunders) String() string {
	return "BirnbaumSaunders: Parameters - " + bs.Parameters().String() + ", Support(x) - " + bs.Support().String()
}

// α ∈ (0,∞)
// β ∈ (0,∞)
func (bs *BirnbaumSaunders) Parameters() stats.Limits {
	return stats.Limits{
		"α": stats.Interval{0, math.Inf(1), true, true},
		"γ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (bs *BirnbaumSaunders) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

func (bs *BirnbaumSaunders) Probability(x float64) float64 {
	if bs.Support().IsWithinInterval(x) {
		num := math.Exp(-(math.Pow(-1+x*bs.scale, 2) / (2 * x * (bs.shape * bs.shape) * bs.scale))) * (1 + x*bs.scale)
		denom := 2 * math.Sqrt(2*math.Pi) * bs.shape * math.Sqrt((x*x*x)*bs.scale)
		return num / denom
	}

	return 0
}

func (bs *BirnbaumSaunders) Distribution(x float64) float64 {
	if bs.Support().IsWithinInterval(x) {
		return .5 * (1 + specfunc.Erf((-1+x*bs.scale)/(math.Sqrt(2)*bs.shape*math.Sqrt(x*bs.scale))))
	}

	return 0
}

func (bs *BirnbaumSaunders) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return (1 + (bs.shape*bs.shape)*math.Pow(math.Erfcinv(2*p), 2) - bs.shape*math.Erfcinv(2*p)*math.Sqrt(2+(bs.shape*bs.shape)*math.Pow(math.Erfcinv(2*p), 2))) / bs.scale

}

func (bs *BirnbaumSaunders) Mean() float64 {
	return (2 + (bs.shape * bs.shape)) / (2 * bs.scale)
}

func (bs *BirnbaumSaunders) Variance() float64 {
	return ((bs.shape * bs.shape) * (4 + 5*(bs.shape*bs.shape)) / (4 * (bs.scale * bs.scale)))
}

func (bs *BirnbaumSaunders) Median() float64 {
	return 1 / bs.scale
}

func (bs *BirnbaumSaunders) ExKurtosis() float64 {
	return (48 + 360*(bs.shape*bs.shape) + 633*(bs.shape*bs.shape*bs.shape*bs.shape)) / math.Pow(4+5*(bs.shape*bs.shape), 2)
}

func (bs *BirnbaumSaunders) Skewness() float64 {
	return (4 * bs.shape * (6 + 11*(bs.shape*bs.shape))) / math.Pow(4*5*(bs.shape*bs.shape), 3./2)
}

func (bs *BirnbaumSaunders) Rand() float64 {
	var rnd float64
	if bs.src != nil {
		rnd = rand.New(bs.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return bs.Inverse(rnd)
}
