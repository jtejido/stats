package math

import gomath "math"

func finalResult(w, y float64) float64 {
	f0 := w - y
	f1 := 1 + y
	f00 := f0 * f0
	f11 := f1 * f1
	f0y := f0 * y
	return w - 4*f0*(6*f1*(f11+f0y)+f00*y)/(f11*(24*f11+36*f0y)+f00*(6*y*y+8*f1*y+f0y))
}

func lambertWSeries(p float64) float64 {
	q := []float64{
		-1,
		+1,
		-0.333333333333333333,
		+0.152777777777777778,
		-0.0796296296296296296,
		+0.0445023148148148148,
		-0.0259847148736037625,
		+0.0156356325323339212,
		-0.00961689202429943171,
		+0.00601454325295611786,
		-0.00381129803489199923,
		+0.00244087799114398267,
		-0.00157693034468678425,
		+0.00102626332050760715,
		-0.000672061631156136204,
		+0.000442473061814620910,
		-0.000292677224729627445,
		+0.000194387276054539318,
		-0.000129574266852748819,
		+0.0000866503580520812717,
		-0.0000581136075044138168,
	}
	ap := gomath.Abs(p)
	if ap < 0.01159 {
		return -1 + p*(1+p*(q[2]+p*(q[3]+p*(q[4]+p*(q[5]+p*q[6])))))
	} else if ap < 0.0766 {

		return -1 + p*(1+p*(q[2]+p*(q[3]+p*(q[4]+p*(q[5]+p*(q[6]+p*(q[7]+p*(q[8]+p*(q[9]+p*q[10])))))))))
	} else {

		return -1 + p*(1+p*(q[2]+p*(q[3]+p*(q[4]+p*(q[5]+p*(q[6]+p*(q[7]+p*(q[8]+p*(q[9]+p*(q[10]+p*(q[11]+p*(q[12]+p*(q[13]+p*(q[14]+p*(q[15]+p*(q[16]+p*(q[17]+p*(q[18]+p*(q[19]+p*q[20])))))))))))))))))))
	}
}

func lambertW0ZeroSeries(z float64) float64 {
	return z * (1 - z*(1-z*(1.5-z*(2.6666666666666666667-z*(5.2083333333333333333-z*(10.8-z*(23.343055555555555556-z*(52.012698412698412698-z*(118.62522321428571429-z*(275.57319223985890653-z*(649.78717234347442681-z*(1551.1605194805194805-z*(3741.4497029592385495-z*(9104.5002411580189358-z*(22324.308512706601434-z*(55103.621972903835338-z*136808.86090394293563))))))))))))))))
}

func LambertWm1(z float64) float64 {
	e := [64]float64{}
	g := [64]float64{}
	a := [12]float64{}
	b := [12]float64{}

	if e[0] == 0 {
		e1 := 1 / gomath.E
		ej := e1
		e[0] = gomath.E
		g[0] = -e1
		j := 0
		for jj := 1; jj < 64; jj++ {
			ej *= e1
			e[jj] = e[j] * gomath.E
			g[jj] = -(float64(jj) + 1.) * ej
			j = jj
		}
		a[0] = gomath.Sqrt(gomath.E)
		b[0] = 0.5
		j = 0
		for jj := 1; jj < 12; jj++ {
			a[jj] = gomath.Sqrt(a[j])
			b[jj] = b[j] * 0.5
			j = jj
		}
	}
	if z >= 0 {
		// (lambertwm1) Argument out of range.
		return gomath.NaN()
	}
	if z < -0.35 {
		p2 := 2 * (gomath.E*z + 1)
		if p2 > 0 {
			return lambertWSeries(-gomath.Sqrt(p2))
		}
		if p2 == 0 {
			return -1
		}
		// (lambertwm1) Argument out of range
		return gomath.NaN()
	}
	n := 2
	if g[n-1] > z {
		goto line1
	}
	for j := 1; j <= 5; j++ {
		n *= 2
		if g[n-1] > z {
			goto line2
		}
	}
	// (lambertwm1) Argument too small
	return gomath.NaN()
line2:
	{
		nh := n / 2
		for j := 1; j <= 5; j++ {
			nh /= 2
			if nh <= 0 {
				break
			}
			if g[n-nh-1] > z {
				n -= nh
			}
		}
	}
line1:
	n--
	jmax := 11
	if n >= 8 {
		jmax = 8
	} else if n >= 3 {
		jmax = 9
	} else if n >= 2 {
		jmax = 10
	}
	w := -float64(n)
	y := z * e[n-1]
	for j := 0; j < jmax; j++ {
		wj := w - b[j]
		yj := y * a[j]
		if wj < yj {
			w = wj
			y = yj
		}
	}
	return finalResult(w, y)
}

func LambertW0(z float64) float64 {
	e := [66]float64{}
	g := [65]float64{}
	a := [12]float64{}
	b := [12]float64{}

	if e[0] == 0 {
		e1 := 1. / gomath.E
		ej := 1.
		e[0] = gomath.E
		e[1] = 1.
		g[0] = 0.
		j := 1
		for jj := 2; jj < 66; jj++ {
			ej *= gomath.E
			e[jj] = e[j] * e1
			g[j] = float64(j) * ej
			j = jj
		}
		a[0] = gomath.Sqrt(e1)
		b[0] = 0.5
		j = 0
		for jj := 1; jj < 12; jj++ {
			a[jj] = gomath.Sqrt(a[j])
			b[jj] = b[j] * 0.5
			j = jj
		}
	}
	if gomath.Abs(z) < 0.05 {
		return lambertW0ZeroSeries(z)
	}
	if z < -0.35 {
		p2 := 2 * (gomath.E*z + 1)
		if p2 > 0 {
			return lambertWSeries(gomath.Sqrt(p2))
		}
		if p2 == 0 {
			return -1
		}
		// (lambertw0) Argument out of range.
		return gomath.NaN()
	}
	var n int
	for n = 0; n <= 2; n++ {
		if g[n] > z {
			goto line1
		}
	}
	n = 2
	for j := 1; j <= 5; j++ {
		n *= 2
		if g[n] > z {
			goto line2
		}
	}
	// (lambertw0) Argument too large
	return gomath.NaN()
line2:
	{
		nh := n / 2
		for j := 1; j <= 5; j++ {
			nh /= 2
			if nh <= 0 {
				break
			}
			if g[n-nh] > z {
				n -= nh
			}
		}
	}
line1:
	n--
	jmax := 8
	if z <= -0.36 {
		jmax = 12
	} else if z <= -0.3 {
		jmax = 11
	} else if n <= 0 {
		jmax = 10
	} else if n <= 1 {
		jmax = 9
	}
	y := z * e[n+1]
	w := float64(n)
	for j := 0; j < jmax; j++ {
		wj := w + b[j]
		yj := y * a[j]
		if wj < yj {
			w = wj
			y = yj
		}
	}
	return finalResult(w, y)
}

// Toshio Fukushima, "Precise and fast computation of Lambert W-functions without
// transcendental function evaluations", J. Comp. Appl. Math. 244 (2013) 77-89.
func LambertW(branch int, x float64) float64 {
	switch branch {
	case -1:
		return LambertWm1(x)
	case 0:
		return LambertW0(x)
	default:
		return gomath.NaN()
	}
}
