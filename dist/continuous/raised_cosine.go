package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"github.com/jtejido/trig"
	"math"
	"math/rand"
)

// Raised Cosine distribution
// https://en.wikipedia.org/wiki/Raised_cosine_distribution
type RaisedCosine struct {
	location, scale float64
	src             rand.Source
}

func NewRaisedCosine(location, scale float64) (*RaisedCosine, error) {
	return NewRaisedCosineWithSource(location, scale, nil)
}

func NewRaisedCosineWithSource(location, scale float64, src rand.Source) (*RaisedCosine, error) {
	if scale <= 0 {
		return nil, err.Invalid()
	}

	return &RaisedCosine{location, scale, src}, nil
}

// Distribution parameter bounds limits
// μ  ∈ (-∞,∞)
// s  ∈ (0,∞)
func (rs *RaisedCosine) Parameters() stats.Limits {
	return stats.Limits{
		"μ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"s": stats.Interval{0, math.Inf(1), true, true},
	}
}

// Distribution support bounds limits
// x  ∈ [μ-s,μ+s]
func (rs *RaisedCosine) Support() stats.Interval {
	return stats.Interval{rs.location - rs.scale, rs.location + rs.scale, false, false}
}

// Probability density function
func (rs *RaisedCosine) Probability(x float64) float64 {
	if rs.Support().IsWithinInterval(x) {
		return (1 / rs.scale) * trig.Havercosin(((x-rs.location)/rs.scale)*math.Pi)
	}

	return 0
}

// Cumulative distribution function
func (rs *RaisedCosine) Distribution(x float64) float64 {
	if rs.Support().IsWithinInterval(x) {
		return .5 * (1 + ((x - rs.location) / rs.scale) + (1/math.Pi)*math.Sin(((x-rs.location)/rs.scale)*math.Pi))
	}

	return 0
}

// ExKurtosis of the distribution.
func (rs *RaisedCosine) ExKurtosis() float64 {
	return (6 * (90 - math.Pow(math.Pi, 4.))) / (5 * math.Pow(math.Pow(math.Pi, 2.)-6, 2.))
}

// Skewness of the distribution
func (rs *RaisedCosine) Skewness() float64 {
	return 0
}

// Mean of the distribution
func (rs *RaisedCosine) Mean() float64 {
	return rs.location
}

// Median of the distribution
func (rs *RaisedCosine) Median() float64 {
	return rs.location
}

// Mode of the distribution
func (rs *RaisedCosine) Mode() float64 {
	return rs.location
}

func (rs *RaisedCosine) Entropy() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

// Variance of the distribution
func (rs *RaisedCosine) Variance() float64 {
	return (rs.scale * rs.scale) * ((1. / 3) - (2. / (math.Pi * math.Pi)))
}

func (rs *RaisedCosine) Rand() float64 {
	var rnd func() float64
	if rs.src == nil {
		rnd = rand.Float64
	} else {
		rnd = rand.New(rs.src).Float64
	}

	x := math.Pi*rnd() - gsl.PiOver2
	xSq := x * x
	u := rnd()
	u += u
	a := 0
	b := -1
	w := 0.0
	v := 1.0
	iter := 0
	for iter <= iter_rejection {
		a += 2
		b += 2
		v *= xSq / float64(a*b)
		w += v
		if u >= w {
			return x*rs.location + rs.scale
		}

		a += 2
		b += 2
		v *= xSq / float64(a*b)
		w -= v
		if u <= w {
			if x == 0.0 {
				return rs.location + rs.scale
			}

			if x > 0.0 {
				return (math.Pi-x)*rs.location + rs.scale
			}

			return (-math.Pi-x)*rs.location + rs.scale
		}

		iter++
	}

	return math.NaN()
}
