package continuous

import (
	"github.com/jtejido/ggsl/specfunc"
	"github.com/jtejido/stats"
	"math"
	"math/rand"
)

// Irwin-Hall distribution
// https://en.wikipedia.org/wiki/Irwin%E2%80%93Hall_distribution
type IrwinHall struct {
	n   uint
	src rand.Source
}

func NewIrwinHall(n uint) (*IrwinHall, error) {
	return NewIrwinHallWithSource(n, nil)
}

func NewIrwinHallWithSource(n uint, src rand.Source) (*IrwinHall, error) {
	return &IrwinHall{n, src}, nil
}

// n ∈ [0,∞)
func (ih *IrwinHall) Parameters() stats.Limits {
	return stats.Limits{
		"n": stats.Interval{0, math.Inf(1), false, true},
	}
}

// x ∈ [0,n]
func (ih *IrwinHall) Support() stats.Interval {
	return stats.Interval{0, float64(ih.n), false, false}
}

func (ih *IrwinHall) Probability(x float64) float64 {
	if ih.Support().IsWithinInterval(x) {
		if ih.n == 1 {
			return 1.0
		}

		if ih.n == 2 {
			if x <= 1.0 {
				return x
			}

			return 2.0 - x
		}

		sum := 0.
		for k := 0; k <= int(math.Abs(x)); k++ {
			pow := math.Pow(x-float64(k), float64(ih.n)-1)
			sum += math.Pow(-1., float64(k)) * specfunc.Choose(uint(ih.n), uint(k)) * pow
		}

		return (1 / specfunc.Fact(uint(ih.n)-1)) * sum
	}

	return 0
}

func (ih *IrwinHall) Distribution(x float64) float64 {
	if ih.Support().IsWithinInterval(x) {
		if ih.n == 1 {
			return x
		}

		if ih.n == 2 {
			if x <= 1.0 {
				return 0.5 * x * x
			}

			temp := 2.0 - x
			return 1.0 - 0.5*temp*temp
		}

		sum := 0.
		for k := 0; k <= int(math.Abs(x)); k++ {
			pow := math.Pow(x-float64(k), float64(ih.n))
			sum += math.Pow(-1., float64(k)) * specfunc.Choose(uint(ih.n), uint(k)) * pow
		}

		return (1 / specfunc.Fact(uint(ih.n))) * sum
	} else if x >= float64(ih.n) {
		return 1.0
	}

	return 0
}

func (ih *IrwinHall) Mean() float64 {
	return float64(ih.n) / 2.
}

func (ih *IrwinHall) Median() float64 {
	return float64(ih.n) / 2.
}

func (ih *IrwinHall) Mode() float64 {
	if ih.n == 1 {
		var u float64

		if ih.src != nil {
			u = rand.New(ih.src).Float64()
		} else {
			u = rand.Float64()
		}

		return u
	}

	return float64(ih.n) / 2.
}

func (ih *IrwinHall) Variance() float64 {
	return float64(ih.n) / 12.
}

func (ih *IrwinHall) Skewness() float64 {
	return 0
}

func (ih *IrwinHall) ExKurtosis() float64 {
	return -(6 / (5. * float64(ih.n)))
}

func (ih *IrwinHall) Rand() float64 {
	u := rand.Float64

	if ih.src != nil {
		r := rand.New(ih.src)
		u = r.Float64
	}

	s := u()
	var c float64
	for i := uint(1); i < ih.n; i++ {
		temp := u()
		t := s + temp
		if math.Abs(s) > math.Abs(temp) {
			c += (s - t) + temp
		} else {
			c += (temp - t) + s
		}
		s = t
	}

	return s + c

}
