package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// Q-Exponential distribution
// https://en.wikipedia.org/wiki/Q-exponential_distribution
type QExponential struct {
	rate, q float64 // λ, q
	src     rand.Source
}

func NewQExponential(rate, q float64) (*QExponential, error) {
	return NewQExponentialWithSource(rate, q, nil)
}

func NewQExponentialWithSource(rate, q float64, src rand.Source) (*QExponential, error) {
	if rate <= 0 || q >= 2 {
		return nil, err.Invalid()
	}

	r := new(QExponential)
	r.rate = rate
	r.q = q
	r.src = src

	return r, nil
}

func (q *QExponential) String() string {
	return "QExponential: Parameters - " + q.Parameters().String() + ", Support(x) - " + q.Support().String()
}

// λ  ∈ (0,∞)
// q  ∈ (-∞,2)
func (q *QExponential) Parameters() stats.Limits {
	return stats.Limits{
		"λ": stats.Interval{0, math.Inf(1), true, true},
		"q": stats.Interval{math.Inf(-1), 2, true, true},
	}
}

// x  ∈ [0,∞) for q >= 1
// x  ∈ [0,1/(λ(1-q))]
func (q *QExponential) Support() stats.Interval {
	if q.q >= 1 {
		return stats.Interval{0, math.Inf(1), false, true}
	}

	return stats.Interval{0, 1 / (q.rate * (1 - q.q)), false, true}
}

func (q *QExponential) Probability(x float64) float64 {
	if q.Support().IsWithinInterval(x) {
		return (2 - q.q) * q.rate * smath.Expq(-q.rate*x, q.q)
	}

	return 0
}

func (q *QExponential) Distribution(x float64) float64 {
	qp := 1 / (2 - q.q)
	return 1 - smath.Expq(-(q.rate*x)/qp, qp)
}

func (q *QExponential) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		if q.q >= 1 {
			return math.Inf(1)
		}

		return 1 / (q.rate * (1 - q.q))
	}

	qp := 1 / (2 - q.q)
	return (-qp * smath.Logq(p, qp)) / q.rate
}

func (q *QExponential) Mean() float64 {
	if q.q < 3./2 {
		return 1 / (q.rate * (3 - 2*q.q))
	}

	return math.NaN()
}

func (q *QExponential) Median() float64 {
	qp := 1 / (2 - q.q)
	return (-qp * smath.Logq(.5, qp)) / q.rate
}

func (q *QExponential) Mode() float64 {
	return 0
}

func (q *QExponential) Variance() float64 {
	if q.q < 4./3 {
		return (q.q - 2) / (math.Pow(2*q.q-3, 2) * (3*q.q - 4) * (q.rate * q.rate))
	}

	return math.NaN()
}

func (q *QExponential) Skewness() float64 {
	if q.q < 5./4 {
		return (2 / (5 - 4*q.q)) * math.Sqrt((3*q.q-4)/(q.q-2))
	}

	return math.NaN()
}

func (q *QExponential) ExKurtosis() float64 {
	if q.q < 6./5 {
		return 6 * ((-4*(q.q*q.q*q.q) + 17*(q.q*q.q) - 20*q.q + 6) / ((q.q - 2) * (4*q.q - 5) * (5*q.q - 6)))
	}

	return math.NaN()
}

func (q *QExponential) Rand() float64 {
	var rnd func() float64
	if q.src != nil {
		rnd = rand.New(q.src).Float64
	} else {
		rnd = rand.Float64
	}

	return q.Inverse(rnd())
}
