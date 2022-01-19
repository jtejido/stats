package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Wigner semicircle distribution
// https://en.wikipedia.org/wiki/Wigner_semicircle_distribution
type WignerSemiCircle struct {
	radius, center float64
	src            rand.Source
}

func NewWignerSemiCircle(radius, center float64) (*WignerSemiCircle, error) {
	return NewWignerSemiCircleWithSource(radius, center, nil)
}

func NewWignerSemiCircleWithSource(radius, center float64, src rand.Source) (*WignerSemiCircle, error) {
	if radius <= 0 {
		return nil, err.Invalid()
	}

	return &WignerSemiCircle{radius, center, src}, nil
}

// a ∈ (-∞,∞)
// R ∈ (0,∞)
func (ws *WignerSemiCircle) Parameters() stats.Limits {
	return stats.Limits{
		"a": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"R": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [a-R,a+R]
func (ws *WignerSemiCircle) Support() stats.Interval {
	return stats.Interval{ws.center - ws.radius, ws.center + ws.radius, false, false}
}

func (ws *WignerSemiCircle) Probability(x float64) float64 {
	if ws.Support().IsWithinInterval(x) {
		return (2. / math.Pi * (ws.radius * ws.radius)) * math.Sqrt((ws.radius*ws.radius)-math.Pow(-ws.center+x, 2))
	} else if x >= ws.center+ws.radius {
		return 1
	}

	return 0
}

func (ws *WignerSemiCircle) Distribution(x float64) float64 {
	if ws.Support().IsWithinInterval(x) {
		return .5 + ((-ws.center+x)*math.Sqrt((ws.radius*ws.radius)-math.Pow(-ws.center+x, 2)))/(math.Pi*(ws.radius*ws.radius)) + math.Asin((-ws.center+x)/ws.radius)/math.Pi
	} else if x >= ws.center+ws.radius {
		return 1
	}

	return 0
}

func (ws *WignerSemiCircle) Mean() float64 {
	return ws.center
}

func (ws *WignerSemiCircle) Median() float64 {
	return ws.center
}

func (ws *WignerSemiCircle) Mode() float64 {
	return 0
}

func (ws *WignerSemiCircle) Skewness() float64 {
	return 0
}

func (ws *WignerSemiCircle) ExKurtosis() float64 {
	return -1
}

func (ws *WignerSemiCircle) Entropy() float64 {
	return math.Log(math.Pi*ws.radius) - (1. / 2)
}

func (ws *WignerSemiCircle) Variance() float64 {
	return (ws.radius * ws.radius) / 4
}

func (ws *WignerSemiCircle) Rand() float64 {
	var rnd float64
	if ws.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(ws.src).Float64()
	}

	rnd += rnd - 1
	return ws.radius * rnd
}
