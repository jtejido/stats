package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Beta Prime distribution (Beta Distribution of the second kind)
// https://en.wikipedia.org/wiki/Beta_prime_distribution
type BetaPrime struct {
	baseContinuousWithSource
	alpha, beta float64
}

func NewBetaPrime(alpha, beta float64) (*BetaPrime, error) {
	return NewBetaPrimeWithSource(alpha, beta, nil)
}

func NewBetaPrimeWithSource(alpha, beta float64, src rand.Source) (*BetaPrime, error) {
	if alpha <= 0 || beta <= 0 {
		return nil, err.Invalid()
	}

	ret := new(BetaPrime)
	ret.alpha = alpha
	ret.beta = beta
	ret.src = src

	return ret, nil
}

func (bp *BetaPrime) String() string {
	return "BetaPrime: Parameters - " + bp.Parameters().String() + ", Support(x) - " + bp.Support().String()
}

// α ∈ (0,∞)
// β ∈ (0,∞)
func (bp *BetaPrime) Parameters() stats.Limits {
	return stats.Limits{
		"α": stats.Interval{0, math.Inf(1), true, true},
		"β": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (bp *BetaPrime) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

func (bp *BetaPrime) Probability(x float64) float64 {
	if bp.Support().IsWithinInterval(x) {
		return (math.Pow(x, bp.alpha-1) * math.Pow(1+x, -bp.alpha-bp.beta)) / specfunc.Beta(bp.alpha, bp.beta)
	}

	return 0
}

func (bp *BetaPrime) Distribution(x float64) float64 {
	if bp.Support().IsWithinInterval(x) {
		return specfunc.Beta_inc(bp.alpha, bp.beta, x/(1+x))
	}

	return 0
}

func (bp *BetaPrime) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	b, _ := NewBeta(bp.alpha, bp.beta)
	x := b.Inverse(p)
	return x / (1. - x)

}

func (bp *BetaPrime) Skewness() float64 {
	if bp.beta > 3 {
		return (2 * (2*bp.alpha + bp.beta - 1)) / (bp.beta - 3) * math.Sqrt((bp.beta-2)/(bp.alpha*(bp.alpha+bp.beta-1)))
	}

	return math.NaN()

}

func (bp *BetaPrime) Mean() float64 {
	if bp.beta > 1 {
		return bp.alpha / (bp.beta - 1)
	}

	return math.NaN()
}

func (bp *BetaPrime) Mode() float64 {

	if bp.alpha >= 1 {
		return bp.alpha - 1/bp.beta + 1
	}

	return 0
}

func (bp *BetaPrime) Median() float64 {
	return bp.Inverse(.5)
}

func (bp *BetaPrime) Variance() float64 {
	if bp.beta > 2 {
		return (bp.alpha * (bp.alpha + bp.beta - 1)) / ((bp.beta - 2) * math.Pow(bp.beta-1, 2.))
	}

	return math.NaN()
}

func (bp *BetaPrime) ExKurtosis() float64 {
	num := 3*(bp.alpha*bp.alpha*bp.alpha)*(bp.beta*bp.beta) + 69*(bp.alpha*bp.alpha*bp.alpha)*bp.beta - 30*(bp.alpha*bp.alpha*bp.alpha) + 6*(bp.alpha*bp.alpha)*(bp.beta*bp.beta*bp.beta) + 12*(bp.alpha*bp.alpha)*(bp.beta*bp.beta) - 78*(bp.alpha*bp.alpha)*bp.beta + 60*(bp.alpha*bp.alpha) + 3*bp.alpha*(bp.beta*bp.beta*bp.beta*bp.beta) + 9*bp.alpha*(bp.beta*bp.beta*bp.beta) - 69*bp.alpha*(bp.beta*bp.beta) + 99*bp.alpha*bp.beta - 42*bp.alpha + 6*(bp.beta*bp.beta*bp.beta*bp.beta) - 30*(bp.beta*bp.beta*bp.beta) + 54*(bp.beta*bp.beta) - 42*bp.beta + 12
	denom := (bp.alpha + bp.beta - 1) * (bp.beta - 3) * (bp.beta - 4)
	return num / denom
}

func (bp *BetaPrime) Rand() float64 {
	var rnd float64
	if bp.src != nil {
		rnd = rand.New(bp.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	return bp.Inverse(rnd)
}
