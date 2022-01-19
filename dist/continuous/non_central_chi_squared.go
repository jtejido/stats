package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
	"math/rand"
)

// NonCentralChiSquared distribution
// https://en.wikipedia.org/wiki/Beta_distribution
type NonCentralChiSquared struct {
	dof    int
	lambda float64 // degrees of freedom, non-centrality
	src    rand.Source
}

func NewNonCentralChiSquared(dof int, lambda float64) (*NonCentralChiSquared, error) {
	return NewNonCentralChiSquaredWithSource(dof, lambda, nil)
}

func NewNonCentralChiSquaredWithSource(dof int, lambda float64, src rand.Source) (*NonCentralChiSquared, error) {
	if dof <= 0 || lambda <= 0 {
		return nil, err.Invalid()
	}

	return &NonCentralChiSquared{dof, lambda, src}, nil
}

// k ∈ (0,∞)
// λ ∈ (0,∞)
func (n *NonCentralChiSquared) Parameters() stats.Limits {
	return stats.Limits{
		"k": stats.Interval{0, math.Inf(1), true, true},
		"λ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (n *NonCentralChiSquared) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

func (n *NonCentralChiSquared) Probability(x float64) float64 {
	if n.Support().IsWithinInterval(x) {
		return (1. / 2) * math.Exp(-(x+n.lambda)/2) * math.Pow(x/n.lambda, (float64(n.dof)/4)-(1./2)) * specfunc.Bessel_Inu((float64(n.dof)/2)-1, math.Sqrt(n.lambda*x))
	}

	return 0
}

func (n *NonCentralChiSquared) Distribution(x float64) float64 {
	if n.Support().IsWithinInterval(x) {
		return 1 - smath.MarcumQ(float64(n.dof)/2, math.Sqrt(n.lambda), math.Sqrt(x))
	}

	return 0
}

func (n *NonCentralChiSquared) Inverse(p float64) float64 {
	const (
		eps = 1e-11
	)

	/* Finding an upper and lower bound. This is Pearson's (1959) approximation, which is usually good to 4 figs or so.  */
	var b, c, ff, ux, lx, ux0, nx, pp float64
	b = (n.lambda * n.lambda) / (float64(n.dof) + 3*n.lambda)
	c = (float64(n.dof) + 3*n.lambda) / (float64(n.dof) + 2*n.lambda)
	ff = (float64(n.dof) + 2*n.lambda) / (c * c)
	qcsq := smath.InverseRegularizedLowerIncompleteGamma(ff/2.0, p) * 2.0
	ux = b + c*qcsq
	if ux < 0 {
		ux = 1
	}
	ux0 = ux

	if p > 1-gsl.Float64Eps {
		return math.Inf(1)
	}

	pp = math.Min(1-gsl.Float64Eps, p*(1+eps))
	for ux < gsl.MaxFloat64 && n.Distribution(ux) < pp {
		ux *= 2
	}

	pp = p * (1 - eps)

	lx = math.Min(ux0, gsl.MaxFloat64)
	for lx > gsl.MinFloat64 && n.Distribution(lx) > pp {
		lx *= 0.5
	}

	/* 2. interval (lx,ux)  halving : */
	for (ux-lx)/nx > 1e-13 {
		nx = 0.5 * (lx + ux)
		if n.Distribution(nx) > p {
			ux = nx
		} else {
			lx = nx
		}
	}

	return 0.5 * (ux + lx)
}

func (n *NonCentralChiSquared) Mean() float64 {
	return float64(n.dof) + n.lambda
}

func (n *NonCentralChiSquared) Variance() float64 {
	return 2 * (float64(n.dof) + (2 * n.lambda))
}

func (n *NonCentralChiSquared) Skewness() float64 {
	return (math.Pow(2., 3./2) * (float64(n.dof) + (3. * n.lambda))) / math.Pow(float64(n.dof)+(2*n.lambda), 3./2)
}

func (n *NonCentralChiSquared) ExKurtosis() float64 {
	return (12 * (float64(n.dof) + (4 * n.lambda))) / math.Pow(float64(n.dof)+(2*n.lambda), 2.)
}

func (n *NonCentralChiSquared) Rand() float64 {
	var rnd float64
	if n.src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(n.src).Float64()
	}

	return n.Inverse(rnd)
}
