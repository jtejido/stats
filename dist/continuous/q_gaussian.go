package continuous

import (
	integ "github.com/jtejido/ggsl/integration"
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// Q-Gaussian distribution
// https://en.wikipedia.org/wiki/Q-Gaussian_distribution
type QGaussian struct {
	mean, scale, q float64 // μ, b, q
	src            rand.Source
}

func NewQGaussian(mean, scale, q float64) (*QGaussian, error) {
	return NewQGaussianWithSource(mean, scale, q, nil)
}

func NewQGaussianWithSource(mean, scale, q float64, src rand.Source) (*QGaussian, error) {
	if scale <= 0 || q >= 3 {
		return nil, err.Invalid()
	}

	r := new(QGaussian)
	r.mean = mean
	r.scale = scale
	r.q = q
	r.src = src

	return r, nil
}

func (q *QGaussian) String() string {
	return "QGaussian: Parameters - " + q.Parameters().String() + ", Support(x) - " + q.Support().String()
}

// μ  ∈ (-∞,∞)
// b  ∈ (0,∞)
// q  ∈ (-∞,3)
func (q *QGaussian) Parameters() stats.Limits {
	return stats.Limits{
		"μ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"b": stats.Interval{0, math.Inf(1), true, true},
		"q": stats.Interval{math.Inf(-1), 3, true, true},
	}
}

// x  ∈ (-∞,∞) for 1 <= q < 3
// x  ∈ [-1/sqrt(b(1-q)),1/sqrt(b(1-q))] for q < 1
func (q *QGaussian) Support() stats.Interval {
	if 1 <= q.q && q.q < 3 {
		return stats.Interval{math.Inf(-1), math.Inf(1), true, true}
	}

	return stats.Interval{-(1 / math.Sqrt(q.scale*(1-q.q))), 1 / math.Sqrt(q.scale*(1-q.q)), false, false}

}

func (q *QGaussian) Probability(x float64) float64 {
	if q.q == 1 {
		return math.Exp(-(math.Pow(-x+q.mean, 2) / (2 * (q.scale * q.scale)))) / (math.Sqrt(2*math.Pi) * q.scale)
	} else if 1 < q.q && q.q < 3 {
		num := math.Sqrt(-1+q.q) * math.Pow(1+(((-1+q.q)*math.Pow(-x+q.mean, 2))/(2*(q.scale*q.scale))), 1/(1-q.q)) * specfunc.Gamma(1/(-1+q.q))
		denom := math.Sqrt(2*math.Pi) * q.scale * specfunc.Gamma(-((-3 + q.q) / (2 * (-1 + q.q))))
		return num / denom
	} else if q.q < 1 {
		rge := (math.Sqrt((1-q.q)/(q.scale*q.scale)) * (x - q.mean)) / math.Sqrt(2)
		if -1 <= rge && rge <= 1 {
			num := math.Sqrt(1-q.q) * math.Pow(1+(((-1+q.q)*math.Pow(-x+q.mean, 2))/(2*(q.scale*q.scale))), 1/(1-q.q)) * specfunc.Gamma((3./2)+(1/(1-q.q)))
			denom := math.Sqrt(2*math.Pi) * q.scale * specfunc.Gamma(1+(1/(1-q.q)))
			return num / denom
		}

	}

	return 0
}

type integrand struct {
	pdf func(float64) float64
}

func (i *integrand) Evaluate(x float64) float64 {
	return i.pdf(x)
}

// We don't have closed-form for this. The limit value should be changeable
func (q *QGaussian) Distribution(x float64) float64 {
	f := &integrand{pdf: q.Probability}
	workspace, _ := integ.NewWorkspace(30)
	var cdf, abserr float64
	integ.Qagil(f, x, abserr, 1e-12, 30, workspace, &cdf, &abserr)

	return cdf
}

func (q *QGaussian) Mean() float64 {
	if q.q < 2 {
		return q.mean
	}

	return math.NaN()
}

func (q *QGaussian) Median() float64 {
	return q.mean
}

func (q *QGaussian) Variance() float64 {
	if q.q < 5./3 {
		return (2 * (q.scale * q.scale)) / (5 - 3*q.q)
	} else if 5./3 <= q.q && q.q < 2 {
		return math.Inf(1)
	}

	return math.NaN()
}

func (q *QGaussian) Skewness() float64 {
	if q.q < 3./2 {
		return 0
	}

	return math.NaN()
}

func (q *QGaussian) ExKurtosis() float64 {
	if q.q < 7./5 {
		return 6 * ((q.q - 1) / (7 - 5*q.q))
	}

	return math.NaN()
}

// see also, W. Thistleton, J.A. Marsh, K. Nelson and C. Tsallis
// Generalized Box–Muller method for generating q-Gaussian random deviates
// IEEE Transactions on Information Theory 53, 4805 (2007)
func (q *QGaussian) Rand() float64 {
	var rnd func() float64
	if q.src != nil {
		rnd = rand.New(q.src).Float64
	} else {
		rnd = rand.Float64
	}

	qGen := (1 + q.q) / (3 - q.q)
	u1 := rnd()
	u2 := rnd()

	z := math.Sqrt(-2*smath.Logq(u1, qGen)) * math.Cos(2*math.Pi*u2)
	return q.mean + (z / math.Sqrt(q.scale*(3-q.q)))
}
