package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// Rice distribution
// https://en.wikipedia.org/wiki/Rice_distribution
type Rice struct {
	distance, spread float64 // v, σ
	src              rand.Source
}

func NewRice(distance, spread float64) (*Rice, error) {
	return NewRiceWithSource(distance, spread, nil)
}

func NewRiceWithSource(distance, spread float64, src rand.Source) (*Rice, error) {
	if distance < 0 || spread < 0 {
		return nil, err.Invalid()
	}

	return &Rice{distance, spread, src}, nil
}

// μ ∈ [0,∞)
// σ ∈ [0,∞)
func (r *Rice) Parameters() stats.Limits {
	return stats.Limits{
		"v": stats.Interval{0, math.Inf(1), true, true},
		"σ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [0,∞)
func (r *Rice) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), false, true}
}

func (r *Rice) Probability(x float64) float64 {
	if r.Support().IsWithinInterval(x) {
		return (x / (r.spread * r.spread)) * math.Exp(-((x*x)+(r.distance*r.distance))/(2*(r.spread*r.spread))) * specfunc.Bessel_I0((x*r.distance)/(r.spread*r.spread))
	}

	return 0
}

func (r *Rice) Distribution(x float64) float64 {
	if r.Support().IsWithinInterval(x) {
		return 1 - smath.MarcumQ(1, r.distance/r.spread, x/r.spread)
	}

	return 0
}

func (r *Rice) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	ncs := NonCentralChiSquared{2, math.Pow(r.distance/r.spread, 2), nil}
	return math.Sqrt(ncs.Inverse(p)) * r.spread
}

func (r *Rice) Mean() float64 {
	return r.spread * math.Sqrt(math.Pi/2) * smath.Laguerre(1./2, math.Pow(-r.distance, 2.)/(2*(r.spread*r.spread)))
}

func (r *Rice) Variance() float64 {
	return 2*(r.spread*r.spread) + (r.distance * r.distance) - ((math.Pi*(r.spread*r.spread))/2)*math.Pow(smath.Laguerre(1./2, math.Pow(-r.distance, 2.)/(2*(r.spread*r.spread))), 2.)
}

func (r *Rice) Skewness() float64 {
	m1 := r.rm(1)
	m2 := r.rm(2)
	m3 := r.rm(3)
	return (m1*(2*(m1*m1)-3*m2) + m3) / math.Pow(m2-(m1*m1), 3./2)
}

func (r *Rice) ExKurtosis() float64 {
	m1 := r.rm(1)
	m2 := r.rm(2)
	m3 := r.rm(3)
	m4 := r.rm(4)
	return (-3*(m1*m1*m1*m1) + 6*(m1*m1)*m2 - 4*m1*m3 + m4 - 3*((m2-(m1*m1))*(m2-(m1*m1)))) / ((m2 - (m1 * m1)) * (m2 - (m1 * m1)))

}

func (r *Rice) rm(k float64) float64 {
	return math.Pow(r.spread, k) * math.Pow(2, k/2) * specfunc.Gamma(1+k/2) * smath.Laguerre(k/2, -(r.distance*r.distance)/(2*(r.spread*r.spread)))
}

func (r *Rice) Rand() float64 {
	n := &Normal{0, 1, r.src, nil}
	x := r.spread*n.Rand() + r.distance
	y := r.spread * n.Rand()
	return math.Sqrt((x * x) + (y * y))
}
