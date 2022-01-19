package continuous

import (
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
	"math/rand"
)

// Inverse Gaussian distribution
// https://en.wikipedia.org/wiki/Inverse_Gaussian_distribution
type InverseGaussian struct {
	mean, shape float64 // μ (mean), λ (shape)
	src         rand.Source
}

func NewInverseGaussian(mean, shape float64) (*InverseGaussian, error) {
	return NewInverseGaussianWithSource(mean, shape, nil)
}

func NewInverseGaussianWithSource(mean, shape float64, src rand.Source) (*InverseGaussian, error) {
	if mean <= 0 || shape <= 0 {
		return nil, err.Invalid()
	}

	return &InverseGaussian{mean, shape, src}, nil
}

// μ ∈ (0,∞)
// λ ∈ (0,∞)
func (ig *InverseGaussian) Parameters() stats.Limits {
	return stats.Limits{
		"μ": stats.Interval{0, math.Inf(1), true, true},
		"λ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (ig *InverseGaussian) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

func (ig *InverseGaussian) Probability(x float64) float64 {
	if ig.Support().IsWithinInterval(x) {
		return math.Sqrt(ig.shape/(2*math.Pi*(x*x*x))) * math.Exp(-(ig.shape*math.Pow(x-ig.mean, 2))/(2*(ig.mean*ig.mean)*x))
	}

	return 0
}

func (ig *InverseGaussian) Distribution(x float64) float64 {
	if ig.Support().IsWithinInterval(x) {
		x1 := math.Sqrt((ig.shape / x) * ((x / ig.mean) - 1))
		x2 := -math.Sqrt((ig.shape / x) * ((x / ig.mean) + 1))
		sn := &Normal{0, 1, nil, nil}
		g1 := sn.Distribution(x1)
		g2 := sn.Distribution(x2)

		return g1 + math.Exp((2*ig.shape)/ig.mean)*g2
	}

	return 0
}

func (ig *InverseGaussian) Entropy() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (ig *InverseGaussian) ExKurtosis() float64 {
	return 15 * ig.mean / ig.shape
}

func (ig *InverseGaussian) Skewness() float64 {
	return 3 * math.Pow(ig.mean/ig.shape, .5)
}

func (ig *InverseGaussian) Inverse(p float64) float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (ig *InverseGaussian) Mean() float64 {
	return ig.mean
}

func (ig *InverseGaussian) Median() float64 {
	stats.NotImplementedError()
	return math.NaN()
}

func (ig *InverseGaussian) Mode() float64 {
	aux := 1.5 * ig.mean / ig.shape
	mode := 1 + aux*aux
	mode = math.Sqrt(mode)
	mode -= aux
	return ig.mean * mode
}

func (ig *InverseGaussian) Variance() float64 {
	return (ig.mean * ig.mean * ig.mean) / ig.shape
}

func (ig *InverseGaussian) Rand() float64 {
	nd := &Normal{0, 1, ig.src, nil}
	ud := &Uniform{0, 1, ig.src}
	x := nd.Rand()
	u := ud.Rand()
	x *= x
	mupX := ig.mean * x
	y := 4*ig.shape + mupX
	y = math.Sqrt(y * mupX)
	y -= mupX
	y *= -0.5 / ig.shape
	y++
	if u*(1+y) > 1.0 {
		y = 1.0 / y
	}

	return ig.mean * y
}
