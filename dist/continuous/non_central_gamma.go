package continuous

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"github.com/jtejido/stats/err"
	"math"
)

type NonCentralGamma struct {
	shape, scale, lambda float64 // k, θ, λ (noncentrality)
}

func NewNonCentralGamma(shape, scale, lambda float64) (*NonCentralGamma, error) {
	if shape <= 0 || scale <= 0 || lambda <= 0 {
		return nil, err.Invalid()
	}

	return &NonCentralGamma{shape, scale, lambda}, nil
}

// k ∈ (0,∞)
// θ ∈ (0,∞)
func (g *NonCentralGamma) Parameters() stats.Limits {
	return stats.Limits{
		"k": stats.Interval{0, math.Inf(1), true, true},
		"θ": stats.Interval{0, math.Inf(1), true, true},
		"λ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (0,∞)
func (g *NonCentralGamma) Support() stats.Interval {
	return stats.Interval{0, math.Inf(1), true, true}
}

// I. R. C. de Oliveria and D. F. Ferreira, "Computing the noncentral
//     gamma distribution, its inverse and the noncentrality parameter,"
//     Computational Statistics, vol. 28, no. 4, pp.1663-1680, 01 Aug 2013.
// L. Knüsel and B. Bablok, "Computation of the noncentral gamma
//     distribution," SIAM Journal on Scientific Computing, vol. 17, no. 5,
//     pp.1224-1231, Sep. 1996.
func (g *NonCentralGamma) Probability(x float64) float64 {
	if g.Support().IsWithinInterval(x) {
		maxIter := 5000

		x = x / g.scale
		m := math.Ceil(g.lambda)
		a := g.shape + m
		gx := math.Pow(x, a) * math.Exp(-x) / specfunc.Gamma(a+1) / g.scale
		var gxp float64
		if x == 0 {
			gxp = gx * a / x
		}

		gxr := gxp
		pp := math.Exp(-g.lambda) * math.Pow(g.lambda, m) / specfunc.Gamma(m+1)
		pr := pp
		remain := 1 - pp
		ii := 1.0
		gg := pp * gxp

		for {
			gxp = gxp * x / (a + ii - 1)
			pp = pp * g.lambda / (m + ii)
			gg = gg + pp*gxp
			er := gg * remain
			remain = remain - pp
			if ii > m {
				if er < gsl.Float64Eps || int(ii) > maxIter {
					break
				}
			} else {
				gxr = gxr * (a - ii) / x
				pr = pr * (m - ii + 1) / g.lambda
				gg = gg + pr*gxr
				remain = remain - pr
				if remain < gsl.Float64Eps || int(ii) > maxIter {
					break
				}
			}

			ii++
		}

		return gg
	}

	return 0

}

// I. R. C. de Oliveria and D. F. Ferreira, "Computing the noncentral
//     gamma distribution, its inverse and the noncentrality parameter,"
//     Computational Statistics, vol. 28, no. 4, pp.1663-1680, 01 Aug 2013.
// L. Knüsel and B. Bablok, "Computation of the noncentral gamma
//     distribution," SIAM Journal on Scientific Computing, vol. 17, no. 5,
//     pp.1224-1231, Sep. 1996.
func (g *NonCentralGamma) Distribution(x float64) float64 {
	if g.Support().IsWithinInterval(x) {
		maxIter := 5000

		x = x / g.scale
		m := math.Ceil(g.lambda)
		a := g.shape + m
		gammap := specfunc.Gamma_inc_P(a, x)
		gammar := gammap
		gxr := math.Pow(x, a) * math.Exp(-x) / specfunc.Gamma(a+1) / g.scale
		var gxp float64
		if x != 0 {
			gxp = gxr * a / x
		}

		// lpowMExpL := Pow(g.lambda, m) * Exp(-g.lambda)
		// fac := Fact(m)
		// lpowMExpL / fac
		// Poisson pmf
		pp := math.Exp(-g.lambda) * math.Pow(g.lambda, m) / specfunc.Gamma(m+1)
		pr := pp
		remain := 1 - pp
		ii := 1.0
		cdf := pp * gammap

		for {
			gxp = gxp * x / (a + ii - 1)
			gammap = gammap - gxp
			pp = pp * g.lambda / (m + ii)
			cdf = cdf + pp*gammap
			er := remain * gammap
			remain = remain - pp
			if ii > m {
				if er <= gsl.Float64Eps || int(ii) > maxIter {
					break
				}
			} else {
				gxr = gxr * (a - ii) / x
				gammar = gammar + gxr
				pr = pr * (m - ii + 1) / g.lambda
				cdf = cdf + pr*gammar
				remain = remain - pr
				if remain <= gsl.Float64Eps || int(ii) > maxIter {
					break
				}

			}
			ii++
		}

		return cdf

	}

	return 0
}

// I. R. C. de Oliveria and D. F. Ferreira, "Computing the noncentral
//     gamma distribution, its inverse and the noncentrality parameter,"
//     Computational Statistics, vol. 28, no. 4, pp.1663-1680, 01 Aug 2013.
// L. Knüsel and B. Bablok, "Computation of the noncentral gamma
//     distribution," SIAM Journal on Scientific Computing, vol. 17, no. 5,
//     pp.1224-1231, Sep. 1996.
func (ncg *NonCentralGamma) Inverse(p float64) float64 {
	p = p / ncg.scale
	maxitr := 5000
	d := ncg.lambda * 2
	k := math.Ceil(ncg.lambda)
	a := ncg.shape + k
	n := &Normal{0, 1, nil, nil}
	z := n.Inverse(p)
	x0 := ((a + 4*d) * math.Pow(z+math.Pow(math.Pow(a+2*d, 2)/(a+4*d)-1, .5), 2)) / (a + 2*d)
	xn := x0
	it := 1
	for {
		x := xn
		var gd Gamma
		gd.shape = a
		gd.rate = 1

		gamac := gd.Distribution(x)
		gamad := gamac
		gxd := math.Exp(a*math.Log(x) - x - specfunc.Lngamma(a+1))
		gxc := gxd * a / x
		hxc := gxd * a / x
		hxd := hxc
		ppoic := math.Exp(-ncg.lambda) * math.Pow(ncg.lambda, k) / specfunc.Gamma(k+1)
		ppoid := ppoic
		remain := 1 - ppoic
		cdf := ppoic * gamac
		g := ppoic * hxc
		i := 1.0
		for {
			gxc = gxc * x / (a + i - 1)
			gamac = gamac - gxc
			hxc = hxc * x / (a + i - 1)
			ppoic = ppoic * (ncg.lambda) / (k + i)
			cdf = cdf + ppoic*gamac
			g = g + ppoic*hxc
			error := remain * gamac
			remain = remain - ppoic
			if i > k {
				if error <= gsl.Float64Eps || int(i) > maxitr {
					break
				}
			} else {
				gxd = gxd * (a + 1 - i) / x
				gamad = gamad + gxd
				hxd = hxd * (a - i) / x
				ppoid = ppoid * (k - i + 1) / ncg.lambda
				cdf = cdf + ppoid*gamad
				g = g + ppoid*hxd
				remain = remain - ppoid
				if remain <= gsl.Float64Eps || int(i) > maxitr {
					break
				}
			}
			i++
		}
		if x-(cdf-p)/g <= 0 {
			xn = x / 2
		} else {
			xn = x - (cdf-p)/g
		}
		if math.Abs(xn-x) <= x*gsl.Float64Eps || it > maxitr {
			break
		}
		it++
	}

	return xn
}
