package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
)

type NonCentralBeta struct {
	alpha, beta, lambda float64 // α, β, λ (noncentrality)
}

func NewNonCentralBeta(alpha, beta, lambda float64) (*NonCentralBeta, error) {
	if alpha <= 0 || beta <= 0 || lambda <= 0 {
		return nil, err.Invalid()
	}

	return &NonCentralBeta{alpha, beta, lambda}, nil
}

// α ∈ (0,∞)
// β ∈ (0,∞)
// λ ∈ (0,∞)
func (n *NonCentralBeta) Parameters() stats.Limits {
	return stats.Limits{
		"α": stats.Interval{0, math.Inf(1), true, true},
		"β": stats.Interval{0, math.Inf(1), true, true},
		"λ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,1)
func (n *NonCentralBeta) Support() stats.Interval {
	return stats.Interval{0, 1, true, true}
}

func (n *NonCentralBeta) Probability(x float64) float64 {
	if n.Support().IsWithinInterval(x) {
		num := math.Exp(-n.lambda/2) * math.Pow(1-x, -1+n.beta) * math.Pow(x, -1+n.alpha) * specfunc.Hyperg_1F1(n.alpha+n.beta, n.alpha, (x*n.lambda)/2)
		denom := specfunc.Beta(n.alpha, n.beta)
		return num / denom
	}

	return 0
}

func (n *NonCentralBeta) Distribution(x float64) float64 {
	if n.Support().IsWithinInterval(x) {
		var j int
		var sum, c float64
		lastSum := math.Inf(-1)
		for math.Abs(sum-lastSum) > gsl.Float64Eps {
			lastSum = sum
			d := math.Exp(-n.lambda/2) * (math.Pow(n.lambda/2, float64(j)) / specfunc.Fact(uint(j))) * specfunc.Beta_inc(n.alpha+float64(j), n.beta, x)
			y := d - c
			t := sum + y
			c = (t - sum) - y
			sum = t
			j++
		}

		return sum
	}

	return 0
}
