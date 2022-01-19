package math

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/ggsl/specfunc"
	gomath "math"
)

const (
	machEp float64 = 1.0 / (1 << 53)
	maxLog float64 = 1024 * gsl.Ln2
	minLog float64 = -1075 * gsl.Ln2
)

// Internal Epsilon function
func ulp(x float64) float64 {
	if gomath.IsInf(x, 1) || gomath.IsInf(x, -1) {
		return gomath.Inf(1)
	}

	return gomath.Abs(x - gomath.Float64frombits(gomath.Float64bits(x)^1))
}

// Laguerre for a floating point representation of n
func Laguerre(n, x float64) float64 {
	return specfunc.Hyperg_1F1(-n, 1, x)
}

// https://en.wikipedia.org/wiki/Laguerre_polynomials#Explicit_examples_and_properties_of_the_generalized_Laguerre_polynomials
// This is defined in terms of Kummer's function of the second kind.
func AssociatedLaguerre(n, a, x float64) float64 {
	return (gomath.Pow(-1, n) / specfunc.Fact(uint(n))) * specfunc.Hyperg_U(-n, a+1, x)
}

// Evaluate the incomplete modified Bessel function of the first kind and zeroth order I_0(theta,kappa).
// This is equivalent to the left tail area of the von Mises distribution.
// -pi<=theta<=pi . Values of theta outside that region will be wrapped into that region. Kappa is concentration param (0<kappa<Inf)
//
// G. W. Hill, "ALGORITHM 518 incomplete Bessel function i0: The
//    von Mises distribution [S14]," ACM Transactions on Mathematical
//    Software, vol. 3, no. 3, pp. 279-284, Sep. 1977.
func Bessel_i0_inc(theta, kappa float64) float64 {

	A1 := 28.0
	A2 := 0.5
	A3 := 100.0
	A4 := 5.0
	CK := 50.0
	C1 := 50.1

	Z := kappa
	var U, res float64
	//This part differs from the paper, because implementing it as in the paper
	//would cause problems for values of theta that are just below pi.
	//This is because (pi-eps(pi))+pi is numerically equal to 2*pi even though
	//pi-eps(pi) is less than pi.
	if theta == gomath.Pi || (theta > -gomath.Pi && theta+gomath.Pi <= 2*gomath.Pi) {
		U = theta + gomath.Pi
	} else {
		U = gomath.Mod(theta+gomath.Pi, 2*gomath.Pi)
	}
	if U < 0 {
		U = U + 2*gomath.Pi
	}

	Y := U - gomath.Pi
	if Z > CK {
		//For large kappa, compute the normal approximation and left tail
		C := 24 * Z
		V := C - C1
		R := gomath.Sqrt((54/(347/V+26-C) - 6 + C) / 12)
		Z = gomath.Sin(Y*0.5) * R
		S := Z * Z * 2
		V = V - S + 3
		Y = (C - S - S - 16) / 3
		Y = ((S+1.75)*S+83.5)/V - Y
		res = gomath.Erf(Z-S/(Y*Y)*Z)*0.5 + 0.5
	} else if Z <= 0 {
		res = (U * 0.5) / gomath.Pi
	} else {
		//For small kappa, sum the IP terms by backwards recursion.
		IP := int(gomath.Trunc(Z*A2 - A3/(Z+A4) + A1))
		P := float64(IP)
		S := gomath.Sin(Y)
		C := gomath.Cos(Y)
		Y = P * Y
		SN := gomath.Sin(Y)
		CN := gomath.Cos(Y)
		R := 0.0
		V := 0.0
		Z := 2 / Z
		var N int = 2
		for N <= IP {
			P = P - 1
			Y = SN
			SN = SN*C - CN*S
			CN = CN*C + Y*S
			R = 1 / (P*Z + R)
			V = (SN/P + V) * R
			N++
		}
		res = (U*0.5 + V) / gomath.Pi
	}

	res = gomath.Max(res, 0)
	return gomath.Min(res, 1)
}

// Helpers to differentiate the two normalized inc from ggsl.
var (
	UpperIncompleteGamma = specfunc.Gamma_inc_Q
	LowerIncompleteGamma = specfunc.Gamma_inc_P
)

// https://en.wikipedia.org/wiki/Tsallis_statistics#q-exponential
func Expq(x, q float64) float64 {
	if q == 1 {
		return gomath.Exp(x)
	}

	if 1+(1-q)*x > 0 {
		return gomath.Pow(1+(1-q)*x, 1/(1-q))
	}

	return 0
}

// https://en.wikipedia.org/wiki/Tsallis_statistics#q-logarithm
func Logq(x, q float64) float64 {
	if x > 0 {
		if q == 1 {
			return gomath.Log(x)
		}

		num := gomath.Pow(x, 1-q) - 1
		denom := 1 - q
		return num / denom
	}

	return gomath.NaN()
}

func Log1pexp(x float64) float64 {
	if x < 20.0 {
		return gomath.Log1p(gomath.Exp(x))
	}

	if x < 35.0 {
		return x + gomath.Exp(-x)
	}

	return x
}

func Log1mexp(x float64) float64 {
	if x < -gsl.Ln2 {
		return gomath.Log1p(-gomath.Exp(x))
	}

	return gomath.Log(-gomath.Expm1(x))
}

// https://en.wikipedia.org/wiki/Harmonic_number#Harmonic_numbers_for_real_and_complex_values
// Interpolating function related to Digamma function.
func Harmonic(x float64) float64 {
	if gomath.IsInf(x, 1) || x == 0 || x == 1 {
		return x
	}
	if x >= 1 && x <= 25 && x == gomath.Trunc(x) {
		res := 1.0
		for ; x > 1; x-- {
			res += 1 / x
		}
		return res
	}

	return gsl.Euler + specfunc.Psi(x+1)

}

func WrapRange(val, min, max float64, mirror bool) float64 {
	if max <= min {
		panic("The maximum bound must be less than the minimum bound")
	}

	if !mirror {
		a := (val - min)
		m := (max - min)

		//Deal with the case where the value is in the primary interval.
		//Treating this as a special case gets rid of finite precision issues
		//that can make values very close to the upper bound of the interval
		//wrap to the bottom.
		if val >= min && val < max {
			return val
		}

		return a - m*gomath.Floor(a/m) + min
	}

	spread := max - min

	x := val - min - spread/2
	x = (gomath.Pi / spread) * x
	x = gomath.Asin(gomath.Sin(x))

	// Scale everything back to the original size
	x = (spread / gomath.Pi) * x

	// Shift the origin back to where it should be.
	return x + min + spread/2

}
