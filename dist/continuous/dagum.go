package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Dagum distribution (Burr III distribution or Inverse Burr)
// https://en.wikipedia.org/wiki/Dagum_distribution
type Dagum struct {
	baseContinuousWithSource
	p, a, scale float64 // scale, shape.p, shape.a
}

func NewDagum(p, a, scale float64) (*Dagum, error) {
	return NewDagumWithSource(p, a, scale, nil)
}

func NewDagumWithSource(p, a, scale float64, src rand.Source) (*Dagum, error) {
	if scale <= 0 || p <= 0 || a <= 0 {
		return nil, err.Invalid()
	}

	r := new(Dagum)
	r.p = p
	r.a = a
	r.scale = scale
	r.src = src

	return r, nil
}

func (d *Dagum) String() string {
	return "Dagum: Parameters - " + d.Parameters().String() + ", Support(x) - " + d.Support().String()
}

// p ∈ (0,∞)
// a ∈ (0,∞)
// b ∈ (0,∞)
func (d *Dagum) Parameters() stats.Limits {
	return stats.Limits{
		"p": stats.Interval{0, math.Inf(1), true, true},
		"a": stats.Interval{0, math.Inf(1), true, true},
		"b": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (d *Dagum) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

func (d *Dagum) Probability(x float64) float64 {
	if d.Support().IsWithinInterval(x) {
		z := x / d.scale
		num := math.Pow(z, d.a*d.p-1)
		denom := math.Pow(math.Pow(z, d.a)+1, d.p+1)
		return ((d.a * d.p) / d.scale) * (num / denom)
	}

	return 0
}

func (d *Dagum) Distribution(x float64) float64 {
	if d.Support().IsWithinInterval(x) {
		z := x / d.scale
		return math.Pow(1+math.Pow(z, -d.a), -d.p)
	}

	return 0
}

func (d *Dagum) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return d.scale * math.Pow(-1+math.Pow(p, -1/d.p), -1/d.a)
}

func (d *Dagum) Mean() float64 {
	if d.a > 1 {
		return d.rm(1)
	}

	return math.NaN()
}

func (d *Dagum) Median() float64 {
	return d.scale * math.Pow(-1+math.Pow(2., 1/d.p), -1/d.a)
}

func (d *Dagum) Mode() float64 {
	return d.scale * math.Pow((d.a*d.p-1)/(d.a+1), 1/d.a)
}

func (d *Dagum) Variance() float64 {
	if d.a > 2 {
		m1 := d.rm(1)
		m2 := d.rm(2)
		return -math.Pow(m1, 2) + m2
	}

	return math.NaN()
}

func (d *Dagum) Skewness() float64 {
	if d.a > 3 {
		m1 := d.rm(1)
		m2 := d.rm(2)
		m3 := d.rm(3)
		return (m1*(2*(m1*m1)-3*m2) + m3) / math.Pow(m2-(m1*m1), 3./2)
	}

	return math.NaN()
}

func (d *Dagum) ExKurtosis() float64 {
	if d.a > 4 {
		m1 := d.rm(1)
		m2 := d.rm(2)
		m3 := d.rm(3)
		m4 := d.rm(4)
		return (-3*(m1*m1*m1*m1) + 6*(m1*m1)*m2 - 4*m1*m3 + m4 - 3*((m2-(m1*m1))*(m2-(m1*m1)))) / ((m2 - (m1 * m1)) * (m2 - (m1 * m1)))
	}

	return math.NaN()
}

func (d *Dagum) rm(n float64) float64 {
	//return math.Pow(d.scale, n) * d.p * specfunc.Beta(d.p+(n/d.a), 1-(n/d.a))
	return math.Pow(d.scale, n) * d.p * specfunc.Beta((d.a*d.p+n)/d.a, (d.a-n)/d.a)
}

func (d *Dagum) Rand() float64 {
	var rnd float64
	if d.src != nil {
		rnd = rand.New(d.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return d.Inverse(rnd)
}
