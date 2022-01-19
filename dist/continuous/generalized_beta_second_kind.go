package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Generalized Beta of the second kind
// https://en.wikipedia.org/wiki/Generalized_beta_distribution#Generalized_beta_of_the_second_kind_(GB2)
type GB2 struct {
	baseContinuousWithSource
	alpha, beta, p, q float64 // α, β, p, q
}

func NewGB2(alpha, beta, p, q float64) (*GB2, error) {
	return NewGB2WithSource(alpha, beta, p, q, nil)
}

func NewGB2WithSource(alpha, beta, p, q float64, src rand.Source) (*GB2, error) {
	if alpha <= 0 || beta <= 0 || p <= 0 || q <= 0 {
		return nil, err.Invalid()
	}

	ret := new(GB2)
	ret.alpha = alpha
	ret.beta = beta
	ret.p = p
	ret.q = q
	ret.src = src

	return ret, nil
}

func (b *GB2) String() string {
	return "GB2: Parameters - " + b.Parameters().String() + ", Support(x) - " + b.Support().String()
}

// α ∈ (0,∞)
// β ∈ (0,∞)
// p ∈ (0,∞)
// q ∈ (0,∞)
func (b *GB2) Parameters() stats.Limits {
	return stats.Limits{
		"α": stats.Interval{0, math.Inf(1), true, true},
		"β": stats.Interval{0, math.Inf(1), true, true},
		"p": stats.Interval{0, math.Inf(1), true, true},
		"q": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (b *GB2) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

func (b *GB2) Probability(x float64) float64 {
	if b.Support().IsWithinInterval(x) {
		num := math.Abs(b.alpha) * math.Pow(x, b.alpha*b.p-1)
		denom := math.Pow(b.beta, b.alpha*b.p) * specfunc.Beta(b.p, b.q) * math.Pow(1+math.Pow(x/b.beta, b.alpha), b.p+b.q)
		return num / denom
	}

	return 0
}

func (b *GB2) Distribution(x float64) float64 {
	if b.Support().IsWithinInterval(x) {
		beta := Beta{alpha: b.p, beta: b.q}
		y := math.Pow(x/b.beta, b.alpha)
		z := y / (1 + y)
		return beta.Distribution(z)
	}

	return 0
}

func (b *GB2) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	beta1 := Beta{alpha: b.p, beta: b.q}
	z1 := beta1.Inverse(p)
	return b.beta * math.Pow(z1/(1-z1), 1/b.alpha)
}

func (b *GB2) Mean() float64 {
	return b.rm(1)
}

func (b *GB2) Variance() float64 {
	m1 := b.rm(1)
	m2 := b.rm(2)
	return -math.Pow(m1, 2) + m2
}

func (b *GB2) Skewness() float64 {
	m1 := b.rm(1)
	m2 := b.rm(2)
	m3 := b.rm(3)
	return (m1*(2*(m1*m1)-3*m2) + m3) / math.Pow(m2-(m1*m1), 3./2)
}

func (b *GB2) ExKurtosis() float64 {
	m1 := b.rm(1)
	m2 := b.rm(2)
	m3 := b.rm(3)
	m4 := b.rm(4)
	return (-3*(m1*m1*m1*m1) + 6*(m1*m1)*m2 - 4*m1*m3 + m4 - 3*((m2-(m1*m1))*(m2-(m1*m1)))) / ((m2 - (m1 * m1)) * (m2 - (m1 * m1)))
}

func (b *GB2) rm(k float64) float64 {
	num := math.Pow(b.beta, k) * specfunc.Beta(b.p+k/b.alpha, b.q-k/b.alpha)
	denom := specfunc.Beta(b.p, b.q)
	return num / denom
}

func (b *GB2) Rand() float64 {
	var rnd float64
	if b.src != nil {
		rnd = rand.New(b.src).Float64()
	} else {
		rnd = rand.Float64()
	}

	beta := Beta{alpha: b.p, beta: b.q}
	z := beta.Inverse(rnd)
	y := z / (1 - z)
	return b.beta * math.Pow(y, 1/b.alpha)
}
