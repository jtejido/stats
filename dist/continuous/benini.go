package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Benini distribution
// https://en.wikipedia.org/wiki/Benini_distribution
type Benini struct {
	baseContinuousWithSource
	alpha, beta, sigma float64
}

func NewBenini(alpha, beta, sigma float64) (*Benini, error) {
	return NewBeniniWithSource(alpha, beta, sigma, nil)
}

func NewBeniniWithSource(alpha, beta, sigma float64, src rand.Source) (*Benini, error) {
	if alpha < 0 || beta <= 0 || sigma <= 0 {
		return nil, err.Invalid()
	}

	ret := new(Benini)
	ret.alpha = alpha
	ret.beta = beta
	ret.sigma = sigma
	ret.src = src
	return ret, nil
}

func (b *Benini) String() string {
	return "Benini: Parameters - " + b.Parameters().String() + ", Support(x) - " + b.Support().String()
}

// α ∈ [0,∞), this allows for 2-parameter Benini instead of α ∈ (0,∞) without a need for another type
// β ∈ (0,∞)
// σ ∈ (0,∞)
func (b *Benini) Parameters() stats.Limits {
	return stats.Limits{
		"α": stats.Interval{0, math.Inf(1), false, true},
		"β": stats.Interval{0, math.Inf(1), true, true},
		"σ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [σ,∞)
func (b *Benini) Support() stats.Interval {
	return stats.Interval{b.sigma, math.Inf(1), false, true}
}

func (b *Benini) Probability(x float64) float64 {
	if b.Support().IsWithinInterval(x) {
		return math.Exp(-b.alpha*math.Log(x/b.sigma)-b.beta*math.Pow(math.Log(x/b.sigma), 2.)) * ((b.alpha / x) + ((2 * b.beta * math.Log(x/b.sigma)) / x))
	}

	return 0
}

func (b *Benini) Distribution(x float64) float64 {
	if b.Support().IsWithinInterval(x) {
		return 1 - math.Exp(-b.alpha*math.Log(x/b.sigma)-b.beta*math.Pow(math.Log(x/b.sigma), 2.))
	}

	return 0
}

func (b *Benini) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	if b.alpha == 0 {
		return b.sigma * math.Exp(math.Sqrt(-(1/b.beta)*math.Log(1-p)))
	}

	stats.NotImplementedError()
	return math.NaN()

}

func (b *Benini) Mean() float64 {
	return b.sigma + (b.sigma/math.Sqrt(2.*b.beta))*specfunc.Hermite_prob(-1, (-1+b.alpha)/math.Sqrt(2*b.beta))
}

func (b *Benini) Median() float64 {
	return b.sigma * (math.Exp((-b.alpha + math.Sqrt((b.alpha*b.alpha)+b.beta*math.Log(16))) / (2 * b.beta)))
}

func (b *Benini) Variance() float64 {
	mean := b.Mean()
	return ((b.sigma * b.sigma) + ((2*(b.sigma*b.sigma))/math.Sqrt(2*b.beta))*specfunc.Hermite_prob(-1, (-2+b.alpha)/math.Sqrt(2*b.beta))) - (mean * mean)
}

func (b *Benini) Entropy() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (b *Benini) Mode() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (b *Benini) ExKurtosis() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (b *Benini) Skewness() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (b *Benini) Rand() float64 {
	if b.alpha == 0 {
		stats.NotImplementedError()
		return math.NaN()
	}

	var rnd float64
	if b.src != nil {
		rnd = rand.New(b.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return b.Inverse(rnd)
}
