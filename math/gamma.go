package math

import (
	"github.com/jtejido/ggsl/specfunc"
	gomath "math"
)

// Lower Incomplete Gamma
func Ligamma(a, z float64) float64 {
	return gomath.Pow(z, a) * specfunc.Hyperg_1F1(a, a+1, -z) / a
}

// Inverse of the upper incomplete Gamma function
func InverseRegularizedUpperIncompleteGamma(a, p float64) float64 {
	if gomath.IsNaN(a) || gomath.IsNaN(p) {
		return gomath.NaN()
	}

	if p < 0 || p > 1 || a <= 0 {
		panic("out of bounds")
	}

	x0 := gomath.MaxFloat64
	yl := 0.0
	x1 := 0.0
	yh := 1.0
	dithresh := 5.0 * machEp

	if p == 0 {
		return gomath.Inf(1)
	}

	if p == 1 {
		return 0.0
	}

	//  IgamC(a, x) - p = 0
	d := 1.0 / (9.0 * a)
	y := 1.0 - d - Ndtri(p)*gomath.Sqrt(d)
	x := a * y * y * y

	lgm := specfunc.Lngamma(a)

	for i := 0; i < 10; i++ {
		if x > x0 || x < x1 {
			break
		}

		y = UpperIncompleteGamma(a, x)

		if y < yl || y > yh {
			break
		}

		if y < p {
			x0 = x
			yl = y
		} else {
			x1 = x
			yh = y
		}

		// Compute the derivative of the function at this point
		d = (a-1)*gomath.Log(x) - x - lgm
		if d < -maxLog {
			break
		}
		d = -gomath.Exp(d)

		// Compute the step to the next approximation of x
		d = (y - p) / d
		if gomath.Abs(d/x) < machEp {
			return x
		}
		x = x - d
	}

	d = 0.0625
	if x0 == gomath.MaxFloat64 {
		if x <= 0 {
			x = 1
		}
		for x0 == gomath.MaxFloat64 {
			x = (1 + d) * x
			y = UpperIncompleteGamma(a, x)
			if y < p {
				x0 = x
				yl = y
				break
			}
			d = d + d
		}
	}

	d = 0.5
	dir := 0
	for i := 0; i < 400; i++ {
		x = x1 + d*(x0-x1)
		y = UpperIncompleteGamma(a, x)

		lgm = (x0 - x1) / (x1 + x0)
		if gomath.Abs(lgm) < dithresh {
			break
		}

		lgm = (y - p) / p
		if gomath.Abs(lgm) < dithresh {
			break
		}

		if x <= 0 {
			break
		}

		if y >= p {
			x1 = x
			yh = y
			if dir < 0 {
				dir = 0
				d = 0.5
			} else if dir > 1 {
				d = 0.5*d + 0.5
			} else {
				d = (p - yl) / (yh - yl)
			}
			dir++
		} else {
			x0 = x
			yl = y
			if dir > 0 {
				dir = 0
				d = 0.5
			} else if dir < -1 {
				d = 0.5 * d
			} else {
				d = (p - yl) / (yh - yl)
			}
			dir--
		}
	}

	return x
}

// Inverse of the lower incomplete Gamma function
func InverseRegularizedLowerIncompleteGamma(a, y0 float64) float64 {

	if gomath.IsNaN(a) || gomath.IsNaN(y0) {
		return gomath.NaN()
	}

	if y0 < 0 || y0 > 1 || a <= 0 {
		panic("out of bounds")
	}

	xUpper := gomath.MaxFloat64
	xLower := 0.
	yUpper := 1.
	yLower := 0.
	dithresh := 5.0 * machEp

	if y0 == 0. {
		return 0
	}

	if y0 == 1 {
		return gomath.Inf(1)
	}

	y0 = 1 - y0

	// Initial Guess
	d := 1 / (9 * a)
	y := 1 - d - (0.98 * gomath.Sqrt2 * gomath.Erfinv((2.0*y0)-1.0) * gomath.Sqrt(d))
	x := a * y * y * y
	lgm := specfunc.Lngamma(a)

	for i := 0; i < 20; i++ {
		if x < xLower || x > xUpper {
			d = 0.0625
			break
		}

		y = 1 - LowerIncompleteGamma(a, x)
		if y < yLower || y > yUpper {
			d = 0.0625
			break
		}

		if y < y0 {
			xUpper = x
			yLower = y
		} else {
			xLower = x
			yUpper = y
		}

		d = ((a - 1) * gomath.Log(x)) - x - lgm
		if d < -709.78271289338399 {
			d = 0.0625
			break
		}

		d = -gomath.Exp(d)
		d = (y - y0) / d
		if gomath.Abs(d/x) < machEp {
			return x
		}

		if (d > (x / 4)) && (y0 < 0.05) {
			// Naive heuristics for cases near the singularity
			d = x / 10
		}

		x -= d
	}

	if xUpper == gomath.MaxFloat64 {
		if x <= 0 {
			x = 1
		}

		for xUpper == gomath.MaxFloat64 {
			x = (1 + d) * x
			y = 1 - LowerIncompleteGamma(a, x)
			if y < y0 {
				xUpper = x
				yLower = y
				break
			}

			d = d + d
		}
	}

	dir := 0
	d = 0.5
	for i := 0; i < 400; i++ {
		x = xLower + (d * (xUpper - xLower))
		y = 1 - LowerIncompleteGamma(a, x)
		lgm = (xUpper - xLower) / (xLower + xUpper)
		if gomath.Abs(lgm) < dithresh {
			return x
		}

		lgm = (y - y0) / y0
		if gomath.Abs(lgm) < dithresh {
			return x
		}

		if x <= 0 {
			return 0
		}

		if y >= y0 {
			xLower = x
			yUpper = y
			if dir < 0 {
				dir = 0
				d = 0.5
			} else {
				if dir > 1 {
					d = (0.5 * d) + 0.5
				} else {
					d = (y0 - yLower) / (yUpper - yLower)
				}
			}

			dir = dir + 1
		} else {
			xUpper = x
			yLower = y
			if dir > 0 {
				dir = 0
				d = 0.5
			} else {
				if dir < -1 {
					d = 0.5 * d
				} else {
					d = (y0 - yLower) / (yUpper - yLower)
				}
			}

			dir = dir - 1
		}
	}

	return x
}

func InverseRegularizedIncompleteBeta(aa, bb, yy0 float64) float64 {
	var a, b, y0, d, y, x, x0, x1, lgm, yp, di, dithresh, yl, yh, xt float64
	var i, rflg, dir, nflg int

	i = 0
	if yy0 <= 0 {
		return 0
	}
	if yy0 >= 1.0 {
		return 1
	}
	x0 = 0.0
	yl = 0.0
	x1 = 1.0
	yh = 1.0
	nflg = 0

	if aa <= 1.0 || bb <= 1.0 {
		dithresh = 1.0e-6
		rflg = 0
		a = aa
		b = bb
		y0 = yy0
		x = a / (a + b)
		y = specfunc.Beta_inc(a, b, x)
		goto ihalve
	} else {
		dithresh = 1.0e-4
	}
	// Approximation to inverse function
	yp = -Ndtri(yy0)

	if yy0 > 0.5 {
		rflg = 1
		a = bb
		b = aa
		y0 = 1.0 - yy0
		yp = -yp
	} else {
		rflg = 0
		a = aa
		b = bb
		y0 = yy0
	}

	lgm = (yp*yp - 3.0) / 6.0
	x = 2.0 / (1.0/(2.0*a-1.0) + 1.0/(2.0*b-1.0))
	d = yp*gomath.Sqrt(x+lgm)/x - (1.0/(2.0*b-1.0)-1.0/(2.0*a-1.0))*(lgm+5.0/6.0-2.0/(3.0*x))
	d = 2.0 * d
	if d < minLog {
		// mtherr("incbi", UNDERFLOW)
		x = 0
		goto done
	}
	x = a / (a + b*gomath.Exp(d))
	y = specfunc.Beta_inc(a, b, x)
	yp = (y - y0) / y0
	if gomath.Abs(yp) < 0.2 {
		goto newt
	}

	/* Resort to interval halving if not close enough. */
ihalve:

	dir = 0
	di = 0.5
	for i = 0; i < 100; i++ {
		if i != 0 {
			x = x0 + di*(x1-x0)
			if x == 1.0 {
				x = 1.0 - machEp
			}
			if x == 0.0 {
				di = 0.5
				x = x0 + di*(x1-x0)
				if x == 0.0 {
					// mtherr("incbi", UNDERFLOW)
					goto done
				}
			}
			y = specfunc.Beta_inc(a, b, x)
			yp = (x1 - x0) / (x1 + x0)
			if gomath.Abs(yp) < dithresh {
				goto newt
			}
			yp = (y - y0) / y0
			if gomath.Abs(yp) < dithresh {
				goto newt
			}
		}
		if y < y0 {
			x0 = x
			yl = y
			if dir < 0 {
				dir = 0
				di = 0.5
			} else if dir > 3 {
				di = 1.0 - (1.0-di)*(1.0-di)
			} else if dir > 1 {
				di = 0.5*di + 0.5
			} else {
				di = (y0 - y) / (yh - yl)
			}
			dir += 1
			if x0 > 0.75 {
				if rflg == 1 {
					rflg = 0
					a = aa
					b = bb
					y0 = yy0
				} else {
					rflg = 1
					a = bb
					b = aa
					y0 = 1.0 - yy0
				}
				x = 1.0 - x
				y = specfunc.Beta_inc(a, b, x)
				x0 = 0.0
				yl = 0.0
				x1 = 1.0
				yh = 1.0
				goto ihalve
			}
		} else {
			x1 = x
			if rflg == 1 && x1 < machEp {
				x = 0.0
				goto done
			}
			yh = y
			if dir > 0 {
				dir = 0
				di = 0.5
			} else if dir < -3 {
				di = di * di
			} else if dir < -1 {
				di = 0.5 * di
			} else {
				di = (y - y0) / (yh - yl)
			}
			dir -= 1
		}
	}
	// mtherr("incbi", PLOSS)
	if x0 >= 1.0 {
		x = 1.0 - machEp
		goto done
	}
	if x <= 0.0 {
		// mtherr("incbi", UNDERFLOW)
		x = 0.0
		goto done
	}

newt:
	if nflg > 0 {
		goto done
	}
	nflg = 1
	lgm = specfunc.Lngamma(a+b) - specfunc.Lngamma(a) - specfunc.Lngamma(b)

	for i = 0; i < 8; i++ {
		/* Compute the function at this point. */
		if i != 0 {
			y = specfunc.Beta_inc(a, b, x)
		}
		if y < yl {
			x = x0
			y = yl
		} else if y > yh {
			x = x1
			y = yh
		} else if y < y0 {
			x0 = x
			yl = y
		} else {
			x1 = x
			yh = y
		}
		if x == 1.0 || x == 0.0 {
			break
		}
		/* Compute the derivative of the function at this point. */
		d = (a-1.0)*gomath.Log(x) + (b-1.0)*gomath.Log(1.0-x) + lgm
		if d < minLog {
			goto done
		}
		if d > maxLog {
			break
		}
		d = gomath.Exp(d)
		/* Compute the step to the next approximation of x. */
		d = (y - y0) / d
		xt = x - d
		if xt <= x0 {
			y = (x - x0) / (x1 - x0)
			xt = x0 + 0.5*y*(x-x0)
			if xt <= 0.0 {
				break
			}
		}
		if xt >= x1 {
			y = (x1 - x) / (x1 - x0)
			xt = x1 - 0.5*y*(x1-x)
			if xt >= 1.0 {
				break
			}
		}
		x = xt
		if gomath.Abs(d/x) < 128.0*machEp {
			goto done
		}
	}
	/* Did not converge.  */
	dithresh = 256.0 * machEp
	goto ihalve

done:

	if rflg > 0 {
		if x <= machEp {
			x = 1.0 - machEp
		} else {
			x = 1.0 - x
		}
	}
	return (x)
}
