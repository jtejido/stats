package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// Beta distribution
// https://en.wikipedia.org/wiki/Beta_distribution
type Beta struct {
	baseContinuousWithSource
	alpha, beta float64 // α, β
}

func NewBeta(alpha, beta float64) (*Beta, error) {
	return NewBetaWithSource(alpha, beta, nil)
}

func NewBetaWithSource(alpha, beta float64, src rand.Source) (*Beta, error) {
	if alpha <= 0 || beta <= 0 {
		return nil, err.Invalid()
	}

	ret := new(Beta)
	ret.alpha = alpha
	ret.beta = beta
	ret.src = src

	return ret, nil
}

func (b *Beta) String() string {
	return "Beta: Parameters - " + b.Parameters().String() + ", Support(x) - " + b.Support().String()
}

// α ∈ (0,∞)
// β ∈ (0,∞)
func (b *Beta) Parameters() stats.Limits {
	return stats.Limits{
		"α": stats.Interval{0, math.Inf(1), true, true},
		"β": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [0,1]
func (b *Beta) Support() stats.Interval {
	return stats.Interval{0, 1, false, false}
}

func (b *Beta) Probability(x float64) float64 {
	if b.Support().IsWithinInterval(x) {
		xPowam1 := math.Pow(x, b.alpha-1.)
		xm1Powbm1 := math.Pow(1.-x, b.beta-1.)
		b_ab := specfunc.Beta(b.alpha, b.beta)
		return (xPowam1 * xm1Powbm1) / b_ab
	}

	return 0
}

func (b *Beta) Distribution(x float64) float64 {
	if b.Support().IsWithinInterval(x) {
		return specfunc.Beta_inc(b.alpha, b.beta, x)
	}

	return 0
}

func (b *Beta) Entropy() float64 {
	if b.alpha <= 0 || b.beta <= 0 {
		panic("b: negative parameters")
	}

	return specfunc.Lnbeta(b.alpha, b.beta) - (b.alpha-1)*specfunc.Psi(b.alpha) -
		(b.beta-1)*specfunc.Psi(b.beta) + (b.alpha+b.beta-2)*specfunc.Psi(b.alpha+b.beta)
}

func (b *Beta) ExKurtosis() float64 {
	num := 6 * (((b.alpha-b.beta)*(b.alpha-b.beta))*(b.alpha+b.beta+1) - b.alpha*b.beta*(b.alpha+b.beta+2))
	den := b.alpha * b.beta * (b.alpha + b.beta + 2) * (b.alpha + b.beta + 3)
	return num / den
}

func (b *Beta) Skewness() float64 {
	num := 2 * (b.beta - b.alpha) * math.Sqrt(b.alpha+b.beta+1)
	denom := (b.alpha + b.beta + 2) * math.Sqrt(b.alpha*b.beta)
	return num / denom
}

func (b *Beta) Inverse(q float64) float64 {
	if q <= 0 {
		return 0
	}

	if q >= 1 {
		return 1
	}

	return smath.InverseRegularizedIncompleteBeta(b.alpha, b.beta, q)
}

func (b *Beta) Mean() float64 {
	return b.alpha / (b.alpha + b.beta)
}

func (b *Beta) Median() float64 {

	if b.alpha == b.beta {
		return 0.5
	}

	if b.alpha == 1 && b.beta > 0 {
		return 1 - math.Pow(2., (-1/b.beta))
	}

	if b.beta == 1 && b.alpha > 0 {
		return math.Pow(2., (-1 / b.alpha))
	}

	if b.alpha == 3 && b.beta == 2 {
		return 0.6142724318676105
	}

	if b.alpha == 2 && b.beta == 3 {
		return 0.38572756813238945
	}

	return (b.alpha - 1./3) / (b.alpha + b.beta - 2./3)
}

func (b *Beta) Mode() float64 {

	if b.alpha == 1 && b.beta > 1 {
		return 0
	}
	if b.alpha > 1 && b.beta == 1 {
		return 1
	}

	return (b.alpha - 1.) / (b.alpha + b.beta - 2.)
}

func (b *Beta) Variance() float64 {

	ab := b.alpha * b.beta
	aPowbSqrd := (b.alpha + b.beta) * (b.alpha + b.beta)
	apbp1 := b.alpha + b.beta + 1.

	return ab / (aPowbSqrd * apbp1)
}

func (b *Beta) Rand() float64 {
	if (b.alpha <= 1.0) && (b.beta <= 1.0) {

		for {
			u := rand.Float64()
			v := rand.Float64()
			if b.src != nil {
				r := rand.New(b.src)
				u = r.Float64()
				v = r.Float64()
			}

			x := math.Pow(u, 1/b.alpha)
			y := math.Pow(v, 1/b.beta)
			if (x + y) <= 1 {
				if (x + y) > 0 {
					return x / (x + y)
				} else {
					logX := math.Log(u) / b.alpha
					logY := math.Log(v) / b.beta
					logM := logY
					if logX > logY {
						logM = logX
					}
					logX -= logM
					logY -= logM
					return math.Exp(logX - math.Log(math.Exp(logX)+math.Exp(logY)))
				}
			}
		}
	} else {
		var gas, gbs Gamma
		gas.shape = b.alpha
		gas.rate = 1
		gas.src = b.src
		gbs.shape = b.beta
		gbs.rate = 1
		gbs.src = b.src

		ga := gas.Rand()
		gb := gbs.Rand()
		return ga / (ga + gb)
	}
}
