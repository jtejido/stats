package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/linear"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Weibull distribution
// https://en.wikipedia.org/wiki/Weibull_distribution
type Weibull struct {
	scale, shape float64 // λ, k
	src          rand.Source
	natural      linear.RealVector
}

func NewWeibull(scale, shape float64) (*Weibull, error) {
	return NewWeibullWithSource(scale, shape, nil)
}

func NewWeibullWithSource(scale, shape float64, src rand.Source) (*Weibull, error) {
	if scale <= 0 || shape <= 0 {
		return nil, err.Invalid()
	}

	return &Weibull{scale, shape, src, nil}, nil
}

// λ ∈ (0,∞)
// k ∈ (0,∞)
func (w *Weibull) Parameters() stats.Limits {
	return stats.Limits{
		"λ": stats.Interval{0, math.Inf(1), true, true},
		"k": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [0,∞)
func (w *Weibull) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), false, false}
}

func (w *Weibull) Probability(x float64) float64 {
	if w.Support().IsWithinInterval(x) {
		return (w.shape / w.scale) * math.Pow(x/w.scale, w.shape-1) * math.Exp(-math.Pow(x/w.scale, w.shape))
	}

	return 0
}

func (w *Weibull) Distribution(x float64) float64 {
	if w.Support().IsWithinInterval(x) {
		return 1 - math.Exp(-math.Pow(x/w.scale, w.shape))
	}

	return 0
}

func (w *Weibull) Mean() float64 {
	return w.scale * specfunc.Gamma(1+1/w.shape)
}

func (w *Weibull) Median() float64 {
	return w.scale * math.Pow(gsl.Ln2, 1/w.shape)
}

func (w *Weibull) Mode() float64 {
	if w.shape > 1 {
		return w.scale * math.Pow((w.shape-1)/w.shape, 1/w.shape)
	} else if w.shape <= 1 {
		return 0
	} else {
		return math.NaN()
	}
}

func (w *Weibull) Inverse(p float64) float64 {
	if p <= 0 {
		return 0
	}

	if p >= 1 {
		return math.Inf(1)
	}

	return w.scale * math.Pow(-math.Log(1-p), 1/w.shape)
}

func (w *Weibull) Skewness() float64 {
	m1 := w.rm(1)
	m2 := w.rm(2)
	m3 := w.rm(3)
	return (m1*(2*(m1*m1)-3*m2) + m3) / math.Pow(m2-(m1*m1), 3./2)
}

func (w *Weibull) ExKurtosis() float64 {
	m1 := w.rm(1)
	m2 := w.rm(2)
	m3 := w.rm(3)
	m4 := w.rm(4)
	return (-3*(m1*m1*m1*m1) + 6*(m1*m1)*m2 - 4*m1*m3 + m4 - 3*((m2-(m1*m1))*(m2-(m1*m1)))) / ((m2 - (m1 * m1)) * (m2 - (m1 * m1)))
}

func (w *Weibull) Variance() float64 {
	m1 := w.rm(1)
	return -(m1 * m1) + w.rm(2)
}

// raw moment
func (w *Weibull) rm(k float64) float64 {
	return math.Pow(w.scale, k) * specfunc.Gamma(1+(k/w.shape))
}

func (w *Weibull) Rand() float64 {
	var rnd float64
	if w.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(w.src).Float64()
	}

	return w.Inverse(rnd)
}

func (w *Weibull) ToExponential() {
	vec, _ := linear.NewSizedArrayRealVector(1)
	vec.SetEntry(0, -1/math.Pow(w.scale, w.shape))
	w.natural = vec

	// E[x^k] = d/dn(-log(-n) - log(k)) = -1/n
	// stats = x^k = math.Pow(-1/n, k)
}

func (w *Weibull) SufficientStatistics(x float64) linear.RealVector {
	vec, _ := linear.NewArrayRealVectorFromSlice([]float64{math.Pow(x, w.shape)})
	return vec
}
