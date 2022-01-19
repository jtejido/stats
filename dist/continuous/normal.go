package continuous

import (
	"github.com/jtejido/linear"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Normal (A.K.A. Gaussian) distribution
// https://en.wikipedia.org/wiki/Normal_distribution
type Normal struct {
	location, scale float64 // μ (location), σ (scale)
	src             rand.Source
	natural         linear.RealVector
}

func NewNormal(location, scale float64) (*Normal, error) {
	return NewNormalWithSource(location, scale, nil)
}

func NewNormalWithSource(location, scale float64, src rand.Source) (*Normal, error) {
	if scale <= 0 {
		return nil, err.Invalid()
	}

	return &Normal{location, scale, src, nil}, nil
}

// μ ∈ (-∞,∞)
// σ ∈ (0,∞)
func (n *Normal) Parameters() stats.Limits {
	return stats.Limits{
		"μ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
		"σ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (-∞,∞)
func (n *Normal) Support() stats.Interval {
	return stats.Interval{math.Inf(-1), math.Inf(1), true, true}
}

func (n *Normal) Probability(x float64) float64 {
	if n.Support().IsWithinInterval(x) {
		σsqrt_2pi := n.scale * math.Sqrt(2*math.Pi)
		xmμsqrd_over_2σsqrd := ((x - n.location) * (x - n.location)) / (2. * (n.scale * n.scale))
		ℯPowNegxmμsqrd_over2σsqrd := math.Exp(-xmμsqrd_over_2σsqrd)

		return (1. / σsqrt_2pi) * ℯPowNegxmμsqrd_over2σsqrd
	}

	return 0
}

func (n *Normal) Distribution(x float64) float64 {
	if n.Support().IsWithinInterval(x) {
		return 0.5 + 0.5*math.Erf(((x-n.location)/n.scale)/math.Sqrt2)
	}

	return 0
}

func (n *Normal) Entropy() float64 {
	return 0.5*math.Log(2.0*math.Pi) + 0.5 + math.Log(n.scale)
}

func (n *Normal) ExKurtosis() float64 {
	return 0
}

func (n *Normal) Skewness() float64 {
	return 0
}

func (n *Normal) Inverse(p float64) float64 {
	if p <= 0. {
		return math.Inf(-1)
	}

	if p >= 1. {
		return math.Inf(1)
	}

	return math.Sqrt(2*(n.scale*n.scale))*math.Erfinv(2*p-1) + n.location
}

func (n *Normal) Mean() float64 {
	return n.location
}

func (n *Normal) Median() float64 {
	return n.location
}

func (n *Normal) Mode() float64 {
	return n.location
}

func (n *Normal) Variance() float64 {
	return n.scale * n.scale
}

func (n *Normal) Rand() float64 {
	return n.rand()*n.scale + n.location
}

// Ratio method (Kinderman-Monahan); see Knuth v2, 3rd ed, p130.
// J. L. Leva, ACM Trans Math Software 18 (1992) 449-453 and 454-455.
func (n *Normal) rand() float64 {
	var (
		rnd           func() float64
		s             = 0.449871 /* Constants from Leva */
		t             = -0.386595
		a             = 0.19600
		b             = 0.25472
		r1            = 0.27597
		r2            = 0.27846
		u, v, x, y, Q float64
	)

	if n.src != nil {
		u := rand.New(n.src)
		rnd = u.Float64
	} else {
		rnd = rand.Float64
	}

	/* Accept P if Q < r1 (Leva) */
	/* Reject P if Q > r2 (Leva) */
	/* Accept if v^2 <= -4 u^2 log(u) (K+M) */
	/* This test is executed 0.012 times on average. */
	for ok := true; ok; ok = (Q >= r1 && (Q > r2 || v*v > -4*u*u*math.Log(u))) {
		/* Generate a point P = (u, v) uniform in a rectangle enclosing
		   the K+M region v^2 <= - 4 u^2 log(u). */

		/* u in (0, 1] to avoid singularity at u = 0 */
		u = 1 - rnd()

		/* v is in the asymmetric interval [-0.5, 0.5).  However v = -0.5
		   is rejected during next validation.  The
		   resulting normal deviate is strictly symmetric about 0
		   (provided that v is symmetric once v = -0.5 is excluded). */
		v = rnd() - 0.5

		/* Constant 1.7156 > sqrt(8/e) (for accuracy); but not by too much (for efficiency). */
		v *= 1.7156

		/* Compute Leva's quadratic form Q */
		x = u - s
		y = math.Abs(v) - t
		Q = x*x + y*(a*y-b*x)
	}

	return (v / u) /* Return slope */
}

// This will set Moment parameters from the given Location and Scale (dual parametrization),
// then precompute Natural from the computed Moment.
// we'll be using Alternate parametrization by using precision τ instead of variance for Natural parameters.
func (n *Normal) ToExponential() {
	// The dual, expectation parameters for normal distribution are η1 = μ and η2 = μ2 + σ2.
	// vec, _ := linear.NewArrayRealVectorFromSlice([]float64{n.location, (n.location * n.location) + (n.scale * n.scale)})
	// n.Moment = vec
	// n.computeNaturalFromMoment()
}

func (n *Normal) SufficientStatistics(x float64) linear.RealVector {
	vec, _ := linear.NewArrayRealVectorFromSlice([]float64{x, x * x})
	return vec
}

func (n *Normal) computeNaturalFromMoment() {
	// m0 := n.Moment.At(0)
	// m1 := n.Moment.At(1)
	// v := m1 - m0*m0

	// vec, _ := linear.NewSizedArrayRealVector(2)
	// vec.SetEntry(0, m0)
	// vec.SetEntry(1, 1/v)
	// n.natural = vec
}
