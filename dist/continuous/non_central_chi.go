package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	smath "github.com/jtejido/stats/math"
	"math"
)

// NonCentralChi distribution
// https://en.wikipedia.org/wiki/Noncentral_chi_distribution
type NonCentralChi struct {
	dof    int     // degrees of freedom
	lambda float64 // λ, non-centrality
}

func NewNonCentralChi(dof int, lambda float64) (*NonCentralChi, error) {
	if dof <= 0 || lambda <= 0 {
		return nil, err.Invalid()
	}

	return &NonCentralChi{dof, lambda}, nil
}

// k ∈ (0,∞)
// λ ∈ (0,∞)
func (n *NonCentralChi) Parameters() stats.Limits {
	return stats.Limits{
		"k": stats.Interval{0, math.Inf(1), true, true},
		"λ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ [0,∞)
func (n *NonCentralChi) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), false, true}
}

func (n *NonCentralChi) Probability(x float64) float64 {
	if n.Support().IsWithinInterval(x) {
		num := math.Exp(-((x*x)+(n.lambda*n.lambda))/2) * math.Pow(x, float64(n.dof)) * n.lambda
		denom := math.Pow(n.lambda*x, float64(n.dof)/2)
		return (num / denom) * specfunc.Bessel_Inu((float64(n.dof)/2)-1, n.lambda*x)
	}

	return 0
}

func (n *NonCentralChi) Distribution(x float64) float64 {
	if n.Support().IsWithinInterval(x) {
		return 1 - smath.MarcumQ(float64(n.dof)/2, n.lambda, x)
	}

	return 0
}

func (n *NonCentralChi) Mean() float64 {
	return math.Sqrt(math.Pi/2) * smath.AssociatedLaguerre(1./2, (float64(n.dof)/2)-1, -(n.lambda*n.lambda)/2)
}

func (n *NonCentralChi) Variance() float64 {
	mean := n.Mean()
	return float64(n.dof) + math.Pow(math.Pi, 2.) - (mean * mean)
}

func (n *NonCentralChi) ExKurtosis() float64 {
	return math.Pow(float64(n.dof)+(n.lambda*n.lambda), 2) + 2*(float64(n.dof)+(n.lambda*n.lambda))
}

func (n *NonCentralChi) Skewness() float64 {
	return 3 * math.Sqrt(math.Pi/2) * smath.AssociatedLaguerre(3./2, (float64(n.dof)/2)-1, -(n.lambda*n.lambda)/2)
}
