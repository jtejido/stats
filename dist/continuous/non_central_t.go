package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
)

// Noncentral t-distribution
// https://en.wikipedia.org/wiki/Noncentral_t-distribution
type NonCentralT struct {
	dof, lambda float64 // v (degrees of freedom), μ (non-centrality)
}

func NewNonCentralT(dof, lambda float64) (*NonCentralT, error) {
	if dof <= 0 {
		return nil, err.Invalid()
	}

	return &NonCentralT{dof, lambda}, nil
}

// ν ∈ (0,∞)
// μ ∈ (-∞,∞)
func (n *NonCentralT) ParameterLimits() stats.Limits {
	return stats.Limits{
		"ν": stats.Interval{0, math.Inf(1), true, true},
		"μ": stats.Interval{math.Inf(-1), math.Inf(1), true, true},
	}
}

// x ∈ (-∞,∞)
func (n *NonCentralT) SupportLimits() stats.Interval {
	return stats.Interval{math.Inf(-1), math.Inf(1), true, true}
}

func (n *NonCentralT) Probability(x float64) float64 {
	ν := n.dof
	μ := n.lambda

	p1 := math.Pow(ν, ν/2.) * math.Gamma(ν+1) * math.Exp(-1*(μ*μ)/2.) / math.Pow(2., ν) / math.Pow(ν+(x*x), ν/2) / math.Gamma(ν/2.)

	f1 := ν/2. + 1.
	f2 := 3. / 2.
	f3 := (μ * μ) * (x * x) / 2 / (ν + (x * x))
	ip1 := math.Sqrt(2.) * μ * x * specfunc.Hyperg_1F1(f1, f2, f3) / (ν + (x * x)) / math.Gamma((ν+1.)/2.)

	f1 = (ν + 1.) / 2.
	f2 = 1. / 2.
	ip2 := specfunc.Hyperg_1F1(f1, f2, f3) / math.Sqrt(ν+(x*x)) / math.Gamma(ν/2.+1)

	return p1 * (ip1 + ip2)

	// Yet to be benchmarked (dt process in R)
	// if x != 0 {
	// 	A := lowerTail(x*math.Sqrt(1+2/n.dof), n.dof+2, n.lambda)
	// 	B := lowerTail(x, n.dof, n.lambda)
	// 	return (n.dof / x) * (A - B)
	// } else {
	// 	A := specfunc.Gamma((n.dof + 1) / 2)
	// 	B := math.Sqrt(math.Pi*n.dof) * specfunc.Gamma(n.dof/2)
	// 	return (A / B) * math.Exp(-(n.lambda * n.lambda) / 2)
	// }
}

func (n *NonCentralT) Distribution(x float64) float64 {
	alnrpi := 0.57236494292470008707
	errmax := 1.0e-10
	itrmax := 100
	r2pi := 0.79788456080286535588

	del := n.lambda
	negdel := false

	if x < 0.0 {
		del = -n.lambda
		negdel = true
	}

	// Initialize twin series.
	en := 1.0
	t := x * x / (x*x + n.dof)
	if math.IsNaN(x) {
		t = 1
	}
	value := 0.

	if t <= 0.0 {
		// upper tail of normal cumulative function
		value += 0.5 * math.Erfc(del/math.Sqrt2)

		if negdel {
			value = 1.0 - value
		}
		return value
	}

	lambda := del * del
	p := 0.5 * math.Exp(-0.5*lambda)
	q := r2pi * p * del
	s := 0.5 - p
	a := 0.5
	b := 0.5 * n.dof
	rxb := math.Pow(1.0-t, b)
	albeta := alnrpi + specfunc.Lngamma(b) - specfunc.Lngamma(a+b)
	xodd := specfunc.Beta_inc(a, b, t)
	godd := 2.0 * rxb * math.Exp(a*math.Log(t)-albeta)
	xeven := 1.0 - rxb
	geven := b * t * rxb
	value = p*xodd + q*xeven

	// Repeat until convergence.
	for {
		a = a + 1.0
		xodd = xodd - godd
		xeven = xeven - geven
		godd = godd * t * (a + b - 1.0) / a
		geven = geven * t * (a + b - 0.5) / (a + 0.5)
		p = p * lambda / (2.0 * en)
		q = q * lambda / (2.0*en + 1.0)
		s = s - p
		en++
		value = value + p*xodd + q*xeven
		errbd := 2.0 * s * (xodd - godd)

		if errbd <= errmax || math.IsNaN(errbd) || en >= float64(itrmax) {
			break
		}
	}

	// upper tail of normal cumulative function
	value = value + (0.5 * math.Erfc(del/math.Sqrt2))

	if negdel {
		value = 1.0 - value
	}

	return value
}

func (n *NonCentralT) Mean() float64 {
	if n.dof == 1 {
		return math.NaN()
	}

	return n.lambda * math.Sqrt(n.dof/2) * specfunc.Gamma((n.dof-1.)/2) / specfunc.Gamma(n.dof/2)
}

func (n *NonCentralT) Mode() float64 {
	if n.lambda == 0 {
		ratio := specfunc.Gamma((n.dof+2)/2) / specfunc.Gamma((n.dof+3)/3)
		return math.Sqrt(2/n.dof) * ratio * n.lambda
	} else if math.IsInf(n.lambda, 1) {
		return math.Sqrt(n.dof/(n.dof+1)) * n.lambda
	} else {
		// not implemented yet
		stats.NotImplementedError()
		return math.NaN()
	}
}

func (n *NonCentralT) Variance() float64 {
	if n.dof > 2 {
		a := (n.dof * (1 + n.lambda*n.lambda)) / (n.dof - 2)
		b := (n.lambda * n.lambda * n.dof) / 2
		c := math.Gamma((n.dof-1)/2) / math.Gamma(n.dof/2)
		return a - b*c*c
	}

	return math.NaN()
}

func (n *NonCentralT) Skewness() float64 {
	if n.dof > 3 {
		na := (n.dof * (-3 + (n.lambda * n.lambda) + 2*n.dof)) / ((-3 + n.dof) * (-2 + n.dof))
		nb := ((1 + (n.lambda * n.lambda)) * n.dof) / (-2 + n.dof)
		nc := ((n.lambda * n.lambda) * n.dof * math.Pow(specfunc.Gamma(.5*(-1+n.dof)), 2)) / (2 * math.Pow(math.Gamma(n.dof/2), 2))
		num := n.lambda * math.Sqrt(n.dof) * specfunc.Gamma(.5*(-1+n.dof)) * (na - 2*(nb-nc))
		denom := math.Sqrt(2) * math.Pow(nb-nc, 3./2) * specfunc.Gamma(n.dof/2)

		return num / denom
	}

	return math.NaN()
}
