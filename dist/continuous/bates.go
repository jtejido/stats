package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"math"
	"math/rand"
)

// Bates distribution
// https://en.wikipedia.org/wiki/Bates_distribution
type Bates struct {
	baseContinuousWithSource
	a, b float64
	n    uint
}

func NewBates(a, b float64, n uint) (*Bates, error) {
	return NewBatesWithSource(a, b, n, nil)
}

func NewBatesWithSource(a, b float64, n uint, src rand.Source) (*Bates, error) {
	ret := new(Bates)
	ret.a = a
	ret.b = b
	ret.n = n
	ret.src = src
	return ret, nil
}

func (b *Bates) String() string {
	return "Bates: Parameters - " + b.Parameter().String() + ", Support(x) - " + b.Support().String()
}

// a ∈ (-∞,∞)
// b ∈ (-∞,∞)
// n ∈ [0,∞)
func (b *Bates) Parameter() stats.Limits {
	return stats.Limits{
		"a": stats.Interval{0, math.Inf(1), true, true},
		"b": stats.Interval{0, math.Inf(1), true, true},
		"n": stats.Interval{0, math.Inf(1), false, true},
	}
}

// x ∈ [a,b]
func (b *Bates) Support() stats.Interval {
	return stats.Interval{b.a, b.b, false, false}
}

func (b *Bates) Probability(x float64) float64 {
	if b.Support().IsWithinInterval(x) {
		if x > b.a && x < b.b {
			sum := 0.
			var k uint
			for k <= b.n {
				pow := math.Pow(((x-b.a)/(b.b-b.a))-float64(k)/float64(b.n), float64(b.n)-1)
				sgn := gsl.Sign(((x - b.a) / (b.b - b.a)) - float64(k)/float64(b.n))
				sum += math.Pow(-1., float64(k)) * specfunc.Choose(b.n, k) * pow * float64(sgn)
				k++
			}

			return sum
		}
	}

	return 0
}

func (b *Bates) Mean() float64 {
	return .5 * (b.a + b.b)
}

func (b *Bates) Variance() float64 {
	return 1. / (12. * float64(b.n)) * math.Pow(b.b-b.a, 2.)
}

func (b *Bates) Skewness() float64 {
	return 0
}

func (b *Bates) ExKurtosis() float64 {
	return -(6 / (5. * float64(b.n)))
}

func (b *Bates) Entropy() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (b *Bates) Distribution(x float64) float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (b *Bates) Rand() float64 {
	ih := &IrwinHall{b.n, b.src}
	return (b.b-b.a)*ih.Rand()/float64(b.n) + b.a
}
