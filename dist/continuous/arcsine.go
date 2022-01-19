package continuous

import (
	"github.com/jtejido/stats"
	"math"
	"math/rand"
)

// Arcsine distribution
// https://en.wikipedia.org/wiki/Arcsine_distribution
type Arcsine struct {
	baseContinuousWithSource
}

func NewArcsine() (*Arcsine, error) {
	return NewArcsineWithSource(nil)
}

func NewArcsineWithSource(src rand.Source) (*Arcsine, error) {
	r := new(Arcsine)
	r.src = src
	return r, nil
}

func (as *Arcsine) String() string {
	return "Arcsine: Support(x) - " + as.Support().String()
}

func (as *Arcsine) Parameters() stats.Limits {
	return stats.Limits{}
}

// x  ∈ [0,1]
func (as *Arcsine) Support() stats.Interval {
	return stats.Interval{0, 1, false, false}
}

func (as *Arcsine) Probability(x float64) float64 {
	if as.Support().IsWithinInterval(x) {
		return 1 / (math.Pi * math.Sqrt(x*(1-x)))
	}

	return 0
}

func (as *Arcsine) Distribution(x float64) float64 {
	if as.Support().IsWithinInterval(x) {
		return (2. / math.Pi) * math.Asin(math.Sqrt(x))
	}

	return 0
}

func (as *Arcsine) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return math.Pow(math.Sin((math.Pi*p)/2), 2)
}

func (as *Arcsine) ExKurtosis() float64 {
	return -(3. / 2.)
}

func (as *Arcsine) Skewness() float64 {
	return 0
}

func (as *Arcsine) Mean() float64 {
	return 1. / 2.
}

func (as *Arcsine) Median() float64 {
	return 1. / 2.
}

func (as *Arcsine) Mode(x float64) float64 {
	return x
}

func (as *Arcsine) Variance() float64 {
	return 1. / 8.
}

func (as *Arcsine) Entropy() float64 {
	return math.Log(math.Pi / 4)
}

func (as *Arcsine) Rand() float64 {
	var rnd float64
	if as.src != nil {
		rnd = rand.New(as.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return as.Inverse(rnd)
}
