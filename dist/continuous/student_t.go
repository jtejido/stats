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

// Student's t-distribution
// https://en.wikipedia.org/wiki/Student%27s_t-distribution
type StudentT struct {
	dof float64 // ν
	src rand.Source
}

func NewStudentT(dof float64) (*StudentT, error) {
	return NewStudentTWithSource(dof, nil)
}

func NewStudentTWithSource(dof float64, src rand.Source) (*StudentT, error) {
	if dof <= 0 {
		return nil, err.Invalid()
	}

	return &StudentT{dof, src}, nil
}

// ν ∈ (0,∞)
func (st *StudentT) Parameters() stats.Limits {
	return stats.Limits{
		"ν": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (-∞,∞)
func (st *StudentT) Support() stats.Interval {
	return stats.Interval{math.Inf(-1), math.Inf(1), true, true}
}

func (st *StudentT) Probability(x float64) float64 {
	return specfunc.Gamma((st.dof+1)/2) * math.Pow(1+((x*x)/st.dof), -(st.dof+1)/2) / (math.Sqrt(st.dof*math.Pi) * specfunc.Gamma(st.dof/2))
}

func (st *StudentT) Distribution(x float64) float64 {
	if x == 0 {
		return .5
	}

	Ix := specfunc.Beta_inc(st.dof/2, .5, st.dof/((x*x)+st.dof))

	if x < 0 {
		return .5 * Ix
	}

	return 1 - .5*Ix

}

func (st *StudentT) Inverse(p float64) float64 {
	if p <= 0 {
		return math.Inf(-1)
	}

	if p >= 1 {
		return math.Inf(1)
	}

	if p > 0.25 && p < 0.75 {

		if p == 0.5 {
			return 0
		}

		u := 1.0 - 2.0*p
		z := smath.InverseRegularizedIncompleteBeta(0.5, 0.5*st.dof, math.Abs(u))
		t := math.Sqrt(st.dof * z / (1.0 - z))

		if p < 0.5 {
			t = -t
		}

		return t
	}

	rflg := -1.

	if p >= 0.5 {
		p = 1.0 - p
		rflg = 1.
	}

	z := smath.InverseRegularizedIncompleteBeta(0.5*st.dof, 0.5, 2.0*p)

	if (gsl.MaxFloat64 * z) < st.dof {
		return rflg * gsl.MaxFloat64
	}

	t := math.Sqrt(st.dof/z - st.dof)

	return rflg * t

}

func (st *StudentT) Mean() float64 {
	if st.dof > 1 {
		return 0
	}

	return math.NaN()
}

func (st *StudentT) Median() float64 {
	return 0
}

func (st *StudentT) Mode() float64 {
	return 0
}

func (st *StudentT) ExKurtosis() float64 {
	if st.dof > 4 {
		return 6 / (st.dof - 4)
	}

	if st.dof > 1 && st.dof <= 4 {
		return math.Inf(1)
	}

	return math.NaN()
}

func (st *StudentT) Skewness() float64 {
	if st.dof > 3 {
		return 0
	}

	return math.NaN()
}

func (st *StudentT) Entropy() float64 {
	h := st.dof / 2
	h1 := h + 1 //2
	return h1*(specfunc.Psi(h1)-specfunc.Psi(h)) + math.Log(st.dof)/2 + specfunc.Lnbeta(h, .5)
}

func (st *StudentT) Variance() float64 {
	if st.dof > 2 {
		return st.dof / (st.dof - 2.)
	}

	if st.dof > 1 && st.dof <= 2 {
		return math.Inf(1)
	}

	return math.NaN()
}

func (st *StudentT) Rand() float64 {
	var rnd func() float64
	if st.src == nil {
		rnd = rand.Float64
	} else {
		rnd = rand.New(st.src).Float64
	}

	b := .461585657 // math.Sqrt(2*math.Exp(-1/2) - 1)
	alpha := st.dof
	for {
		// Step 1
		u := rnd()

		if u < b/2 {
			x := 4*u - b
			// Step 2
			v := rnd()
			if v <= 1-math.Abs(x)/2 {
				return x
			}

			uAlpha := math.Pow(1+(x*x)/alpha, -(alpha+1)/2)

			if v <= uAlpha {
				return x
			}
			continue
		}

		if u < 0.5 {
			// Step 3
			temp := 4*u - 1 - b
			x := (math.Abs(temp) + b) * gsl.Sign(temp)
			v := rnd()

			// Step 4
			if v <= 1-math.Abs(x)/2 {
				return x
			}

			if v >= (1+(b*b))/(1+(x*x)) {
				continue
			}

			uAlpha := math.Pow(1+(x*x)/alpha, -(alpha+1)/2)
			if v <= uAlpha {
				return x
			}
			continue
		}

		if u < 0.75 {
			// Step 5
			temp := 8*u - 5
			x := 2 / ((math.Abs(temp) + 1) * gsl.Sign(temp))
			u1 := rnd()
			v := math.Pow(x, -2) * u1

			// Step 4 again
			if v <= 1-math.Abs(x)/2 {
				return x
			}

			if v >= (1+(b*b))/(1+(x*x)) {
				continue
			}

			uAlpha := math.Pow(1+(x*x)/alpha, -(alpha+1)/2)
			if v <= uAlpha {
				return x
			}
			continue
		}

		// Step 6
		x := 2 / (8*u - 7)
		v := rnd()

		uAlpha := math.Pow(1+(x*x)/alpha, -(alpha+1)/2)
		if v < (x*x)*uAlpha {
			return x
		}
	}
}
