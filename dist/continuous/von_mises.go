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

// von Mises distribution
// https://en.wikipedia.org/wiki/Von_Mises_distribution
type VonMises struct {
	mean, concentration float64        // μ, κ
	support             stats.Interval // allows flexibility with Support as long as x ∈ [any interval of length 2π].
	src                 rand.Source
}

func NewVonMises(mean, concentration float64, support stats.Interval) (*VonMises, error) {
	return NewVonMisesWithSource(mean, concentration, support, nil)
}

func NewVonMisesWithSource(mean, concentration float64, support stats.Interval, src rand.Source) (*VonMises, error) {
	if support.Lower >= support.Upper {
		return nil, err.Error("lower cannot be greater than or equal upper", err.EINVAL)
	}

	if !support.IsEqualLength(defaultLength) {
		return nil, err.Error("length not equals 2π", err.EINVAL)
	}

	if !support.IsWithinInterval(mean) {
		return nil, err.Error("mean is not within support range", err.EINVAL)
	}

	if concentration <= 0 {
		return nil, err.Error("concentration should be greater than 0", err.EINVAL)
	}

	return &VonMises{mean, concentration, support, src}, nil
}

// κ ∈ (0,∞)
func (vm *VonMises) Parameters() stats.Limits {
	return stats.Limits{
		"κ": stats.Interval{0, math.Inf(1), true, true},
	}
}

// x ∈ (any interval of length 2π]
func (vm *VonMises) Support() stats.Interval {
	return vm.support
}

func (vm *VonMises) Probability(x float64) float64 {
	sup := vm.Support()
	if x <= vm.mean+sup.Lower || x > vm.mean+sup.Upper {
		return 0
	}

	return math.Exp(vm.concentration*math.Cos(x-vm.mean)) / (2 * math.Pi * specfunc.Bessel_I0(vm.concentration))

}

func (vm *VonMises) Distribution(x float64) float64 {
	if vm.Support().IsWithinInterval(x) {
		sup := vm.Support()
		val := smath.Bessel_i0_inc(x-vm.mean, vm.concentration) - smath.Bessel_i0_inc(sup.Lower-vm.mean, vm.concentration)
		if val < 0 {
			val++
		}

		return val
	}

	return 0
}

func (vm *VonMises) CircularMean() float64 {
	return vm.mean
}

func (vm *VonMises) Median() float64 {
	return vm.mean
}

func (vm *VonMises) Mode() float64 {
	return vm.mean
}

func (vm *VonMises) Entropy() float64 {
	return math.Log((2*math.Pi)*specfunc.Bessel_I0(vm.concentration)) + vm.concentration*(1-specfunc.Bessel_I1(vm.concentration)/specfunc.Bessel_I0(vm.concentration))
}

func (vm *VonMises) CircularVariance() float64 {
	return 1 - specfunc.Bessel_I1(vm.concentration)/specfunc.Bessel_I0(vm.concentration)
}

func (vm *VonMises) Rand() float64 {
	var rnd func() float64
	if vm.src == nil {
		rnd = rand.Float64
	} else {
		rnd = rand.New(vm.src).Float64
	}

	sup := vm.Support()
	var f float64
	// Step 0
	tau := 1 + math.Sqrt(1+4*math.Pow(vm.concentration, 2))
	rho := (tau - math.Sqrt(2*tau)) / (2 * vm.concentration)
	r := (1 + (rho * rho)) / (2 * rho)

	for {
		// Step 1
		u1 := 2*rnd() - 1 // In range [-1,1]
		z := math.Cos(math.Pi * u1)
		f = (1 + r*z) / (r + z)
		c := vm.concentration * (r - f)

		// Step 2
		u2 := rnd() // in range [0,1]
		if c*(2-c)-u2 > 0 {
			break
		}

		// Step 3
		if math.Log(c/u2)+1-c >= 0 {
			break
		}
	}
	// Step 4
	u3 := 2*rnd() - 1 // In range [-1,1]

	// wrap it around the given support interval
	return smath.WrapRange(gsl.Sign(u3)*math.Acos(f)+vm.mean, sup.Lower, sup.Upper, false)
}
