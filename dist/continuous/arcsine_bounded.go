package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Arcsine distribution
// https://en.wikipedia.org/wiki/Arcsine_distribution
type ArcsineBounded struct {
	Arcsine
	min, max float64
}

func NewArcsineBounded(min, max float64) (*ArcsineBounded, error) {
	return NewArcsineBoundedWithSource(min, max, nil)
}

func NewArcsineBoundedWithSource(min, max float64, src rand.Source) (*ArcsineBounded, error) {
	if max <= min {
		return nil, err.Invalid()
	}

	r := new(ArcsineBounded)
	r.min = min
	r.max = max
	r.src = src

	return r, nil
}

func (asb *ArcsineBounded) String() string {
	return "ArcsineBounded: Parameters - " + asb.Parameters().String() + ", Support(x) - " + asb.Support().String()
}

// a  ∈ (-∞,∞)
// b  ∈ (a,∞)
func (asb *ArcsineBounded) Parameters() stats.Limits {
	return stats.Limits{
		"A": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"B": stats.Interval{asb.min, math.Inf(1), true, true},
	}
}

// x  ∈ [a,b]
func (asb *ArcsineBounded) Support() stats.Interval {
	return stats.Interval{asb.min, asb.max, false, false}
}

func (asb *ArcsineBounded) Probability(x float64) float64 {
	if asb.Support().IsWithinInterval(x) {
		return 1 / (math.Pi * math.Sqrt((x-asb.min)*(asb.max-x)))
	}

	return 0
}

func (asb *ArcsineBounded) Distribution(x float64) float64 {
	if asb.Support().IsWithinInterval(x) {
		return (2. / math.Pi) * math.Asin(math.Sqrt((x-asb.min)/(asb.max-asb.min)))
	}

	return 0
}

func (asb *ArcsineBounded) Inverse(p float64) float64 {
	if p <= 0 {
		return asb.min
	}

	if p >= 1 {
		return asb.max
	}

	return asb.min + (asb.max-asb.min)*math.Pow(math.Sin((math.Pi*p)/2), 2)
}

func (asb *ArcsineBounded) Mean() float64 {
	return (asb.min + asb.max) / 2.
}

func (asb *ArcsineBounded) Median() float64 {
	return (asb.min + asb.max) / 2.
}

func (asb *ArcsineBounded) Mode() float64 {
	return (asb.min + asb.max) / 2.
}

func (asb *ArcsineBounded) Variance() float64 {
	return (1. / 8) * math.Pow(asb.max-asb.min, 2.)
}

func (asb *ArcsineBounded) ExKurtosis() float64 {
	return -(3. / 2.)
}

func (asb *ArcsineBounded) Entropy() float64 {
	return -0.24156447527049044469 + math.Log(asb.max-asb.min)
}

func (asb *ArcsineBounded) Rand() float64 {
	var rnd float64
	if asb.src != nil {
		rnd = rand.New(asb.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return asb.Inverse(rnd)
}
