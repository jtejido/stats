package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// Q-Weibull distribution
// https://en.wikipedia.org/wiki/Q-exponential_distribution
type QWeibull struct {
	rate, shape, q float64 // λ, κ, q
	src            rand.Source
}

func NewQWeibull(rate, shape, q float64) (*QWeibull, error) {
	return NewQWeibullWithSource(rate, shape, q, nil)
}

func NewQWeibullWithSource(rate, shape, q float64, src rand.Source) (*QWeibull, error) {
	if rate <= 0 || shape <= 0 || q >= 2 {
		return nil, err.Invalid()
	}

	r := new(QWeibull)
	r.rate = rate
	r.shape = shape
	r.q = q
	r.src = src

	return r, nil
}

func (q *QWeibull) String() string {
	return "QWeibull: Parameters - " + q.Parameters().String() + ", Support(x) - " + q.Support().String()
}

// λ  ∈ (0,∞)
// κ  ∈ (0,∞)
// q  ∈ (-∞,3)
func (q *QWeibull) Parameters() stats.Limits {
	return stats.Limits{
		"λ": stats.Interval{0, math.Inf(1), true, true},
		"κ": stats.Interval{0, math.Inf(1), true, true},
		"q": stats.Interval{math.Inf(-1), 2, true, true},
	}
}

// x  ∈ [0,∞) for q >= 1
// x  ∈ [0,λ/(1-q)^1/κ]
func (q *QWeibull) Support() stats.Interval {
	if q.q >= 1 {
		return stats.Interval{0, math.Inf(1), false, true}
	}

	return stats.Interval{0, q.rate / math.Pow(1-q.q, 1/q.shape), false, true}
}

func (q *QWeibull) Probability(x float64) float64 {
	if x >= 0 {
		return (2 - q.q) * (q.shape / q.rate) * math.Pow(x/q.rate, q.shape-1) * smath.Expq(-math.Pow(x/q.rate, q.shape), q.q)
	}

	return 0
}

func (q *QWeibull) Distribution(x float64) float64 {
	if x >= 0 {
		qp := 1 / (2 - q.q)
		rp := q.rate / math.Pow(2-q.q, 1/q.shape)
		return 1 - smath.Expq(-math.Pow(x/rp, q.shape), qp)
	}

	return 0
}

func (q *QWeibull) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		if q.q >= 1 {
			return math.Inf(1)
		}

		return q.rate / math.Pow(1-q.q, 1/q.shape)
	}

	return math.Pow((1-math.Pow(1-p, (1-q.q)/(2-q.q)))/(1-q.q), 1/q.shape) * q.rate
}

func (q *QWeibull) Mean() float64 {
	if q.q < 1 {
		return q.rate * (2 + (1 / (1 - q.q)) + (1 / q.shape)) * math.Pow(1-q.q, -1/q.shape) * specfunc.Beta(1+(1/q.shape), 2+(1/(1-q.q)))
	} else if q.q == 1 {
		return q.rate * specfunc.Gamma(1+(1/q.shape))
	} else if 1 < q.q && q.q < 1+((1+2*q.shape)/(1+q.shape)) {
		return q.rate * (2 - q.q) * math.Pow(q.q-1, -(1+q.shape)/q.shape) * specfunc.Beta(1+(1/q.shape), -(1+(1/(1-q.q))+(1/q.shape)))
	} else if 1+(q.shape/(q.shape+1)) <= q.q && q.q < 2 {
		return math.Inf(1)
	} else {
		return math.NaN()
	}
}

func (q *QWeibull) Rand() float64 {
	var rnd func() float64
	if q.src != nil {
		rnd = rand.New(q.src).Float64
	} else {
		rnd = rand.Float64
	}

	return q.Inverse(rnd())
}
