package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// F-distribution
// https://en.wikipedia.org/wiki/F-distribution
type F struct {
	baseContinuousWithSource
	d1, d2 int
}

func NewF(d1, d2 int) (*F, error) {
	return NewFWithSource(d1, d2, nil)
}

func NewFWithSource(d1, d2 int, src rand.Source) (*F, error) {
	if d1 <= 0 || d2 <= 0 {
		return nil, err.Invalid()
	}

	f := new(F)
	f.d1 = d1
	f.d2 = d2
	f.src = src

	return f, nil
}

func (f *F) String() string {
	return "F: Parameters - " + f.Parameters().String() + ", Support(x) - " + f.Support().String()
}

// d₁ ∈ (0,∞)
// d₂ ∈ (0,∞)
func (f *F) Parameters() stats.Limits {
	return stats.Limits{
		"d₁": stats.Interval{0, math.Inf(1), true, true},
		"d₂": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x  ∈ [0,∞)
func (f *F) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), false, true}
}

func (f *F) Probability(x float64) float64 {
	if f.Support().IsWithinInterval(x) {
		a := math.Pow(float64(f.d1)*x, float64(f.d1)) * math.Pow(float64(f.d2), float64(f.d2))
		b := math.Pow(float64(f.d1)*x+float64(f.d2), float64(f.d1)+float64(f.d2))
		num := math.Sqrt(a / b)
		denom := x * specfunc.Beta(float64(f.d1)/2, float64(f.d2)/2)

		return num / denom
	}

	return 0
}

func (f *F) Distribution(x float64) float64 {
	if f.Support().IsWithinInterval(x) {
		return specfunc.Beta_inc(float64(f.d1)/2, float64(f.d2)/2, (float64(f.d1)*x)/(float64(f.d1)*x+float64(f.d2)))
	}

	return 0
}

func (f *F) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	res := 0.

	w := specfunc.Beta_inc(0.5*float64(f.d2), 0.5*float64(f.d1), 0.5)

	if w > p || p < 0.001 {
		w = smath.InverseRegularizedIncompleteBeta(0.5*float64(f.d1), 0.5*float64(f.d2), p)
		res = float64(f.d2) * w / (float64(f.d1) * (1.0 - w))
	} else {
		w = smath.InverseRegularizedIncompleteBeta(0.5*float64(f.d2), 0.5*float64(f.d1), 1.0-p)
		res = (float64(f.d2) - float64(f.d2)*w) / (float64(f.d1) * w)
	}

	return res
}

func (f *F) Entropy() float64 {
	lgd1 := specfunc.Lngamma(float64(f.d1) / 2.)
	lgd2 := specfunc.Lngamma(float64(f.d2) / 2.)
	lgd1pd2 := specfunc.Lngamma((float64(f.d1) + float64(f.d2)) / 2.)
	return lgd1 + lgd2 - lgd1pd2 + (1-(float64(f.d1)/2.))*specfunc.Psi(1+(float64(f.d1)/2.)) - (1+(float64(f.d2)/2.))*specfunc.Psi(1+(float64(f.d2)/2.)) + ((float64(f.d1)+float64(f.d2))/2.)*specfunc.Psi((float64(f.d1)+float64(f.d2))/2.) + math.Log(float64(f.d1)/float64(f.d2))
}

func (f *F) ExKurtosis() float64 {
	if float64(f.d2) <= 8 {
		return math.NaN()
	}
	return (12 / (float64(f.d2) - 6)) * ((5*float64(f.d2)-22)/(float64(f.d2)-8) + ((float64(f.d2)-4)/float64(f.d1))*((float64(f.d2)-2)/(float64(f.d2)-8))*((float64(f.d2)-2)/(float64(f.d1)+float64(f.d2)-2)))
}

func (f *F) Skewness() float64 {
	if float64(f.d2) <= 6 {
		return math.NaN()
	}
	num := (2*float64(f.d1) + float64(f.d2) - 2) * math.Sqrt(8*(float64(f.d2)-4))
	den := (float64(f.d2) - 6) * math.Sqrt(float64(f.d1)*(float64(f.d1)+float64(f.d2)-2))
	return num / den
}

func (f *F) Mean() float64 {
	if float64(f.d2) > 2 {
		return float64(f.d2) / (float64(f.d2) - 2)
	}

	return math.NaN()
}

func (f *F) Median() float64 {
	return f.Inverse(.5)
}

func (f *F) Mode() float64 {
	if float64(f.d1) <= 2 {
		return math.NaN()
	}

	return ((float64(f.d1) - 2) / float64(f.d1)) * (float64(f.d2) / (float64(f.d2) + 2))
}

func (f *F) Variance() float64 {
	if float64(f.d2) <= 4 {
		return math.NaN()
	}

	twoD2pow2d1pd2m2 := (2 * (float64(f.d2) * float64(f.d2))) * (float64(f.d1) + float64(f.d2) - 2)
	d1d2m2Pow2d2m4 := (float64(f.d1) * ((float64(f.d2) - 2) * (float64(f.d2) - 2))) * (float64(f.d2) - 4)

	return twoD2pow2d1pd2m2 / d1d2m2Pow2d2m4
}

func (f *F) Rand() float64 {
	c1 := ChiSquared{dof: f.d1}
	c2 := ChiSquared{dof: f.d2}
	c1.src = f.src
	c2.src = f.src

	return (c1.Rand() / float64(f.d1)) / (c2.Rand() / float64(f.d2))
}
