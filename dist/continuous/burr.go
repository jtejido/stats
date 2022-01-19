package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Burr distribution (Burr XII or Singh–Maddala Distribution)
// https://en.wikipedia.org/wiki/Burr_distribution
type Burr struct {
	baseContinuousWithSource
	c, k, scale float64
}

func NewBurr(c, k, scale float64) (*Burr, error) {
	return NewBurrWithSource(c, k, scale, nil)
}

func NewBurrWithSource(c, k, scale float64, src rand.Source) (*Burr, error) {
	if c <= 0 || k <= 0 || scale <= 0 {
		return nil, err.Invalid()
	}

	ret := new(Burr)
	ret.c = c
	ret.k = k
	ret.scale = scale
	ret.src = src

	return ret, nil
}

func (b *Burr) String() string {
	return "Burr: Parameters - " + b.Parameters().String() + ", Support(x) - " + b.Support().String()
}

// c ∈ (0,∞)
// k ∈ (0,∞)
func (b *Burr) Parameters() stats.Limits {
	return stats.Limits{
		"c": stats.Interval{0, math.Inf(1), true, true},
		"k": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (b *Burr) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

func (b *Burr) rm(r float64) float64 {
	return math.Pow(b.scale, r) * b.k * specfunc.Beta((b.c*b.k-r)/b.c, (b.c+r)/b.c)
}

func (b *Burr) Probability(x float64) float64 {
	if b.Support().IsWithinInterval(x) {
		return b.c * math.Pow(b.scale, -b.c) * b.k * (math.Pow(x, b.c-1) / math.Pow(1+math.Pow(x/b.scale, b.c), b.k+1))
	}

	return 0
}

func (b *Burr) Distribution(x float64) float64 {
	if b.Support().IsWithinInterval(x) {
		return 1 - math.Pow(1+math.Pow(x/b.scale, b.c), -b.k)
	}

	return 0
}

func (b *Burr) Skewness() float64 {
	m1 := b.rm(1)
	m2 := b.rm(2)
	m3 := b.rm(3)
	return (m1*(2*(m1*m1)-3*m2) + m3) / math.Pow(m2-(m1*m1), 3./2)
}

func (b *Burr) Mean() float64 {
	return b.rm(1)
}

func (b *Burr) Median() float64 {
	return math.Pow(math.Pow(2., 1/b.k)-1, 1/b.c) * b.scale
}

func (b *Burr) Mode() float64 {
	return math.Pow((b.c-1)/(b.k*b.c+1), 1/b.c) * b.scale
}

func (b *Burr) Variance() float64 {
	return -math.Pow(b.rm(1), 2.) + b.rm(2)
}

func (b *Burr) ExKurtosis() float64 {
	m1 := b.rm(1)
	m2 := b.rm(2)
	m3 := b.rm(3)
	m4 := b.rm(4)
	return (-3*(m1*m1*m1*m1) + 6*(m1*m1)*m2 - 4*m1*m3 + m4 - 3*((m2-(m1*m1))*(m2-(m1*m1)))) / ((m2 - (m1 * m1)) * (m2 - (m1 * m1)))
}

func (b *Burr) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return math.Pow(math.Expm1((-1/b.k)*math.Log1p(-p)), 1/b.c) * b.scale
}

func (b *Burr) Entropy() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (b *Burr) Rand() float64 {
	var rnd float64
	if b.src != nil {
		rnd = rand.New(b.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return b.Inverse(rnd)
}
