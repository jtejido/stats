package stats

import (
	"fmt"
	gslerr "github.com/jtejido/ggsl/err"
	"math"
)

func init() {
	gslerr.SetErrorHandlerOff()
}

type RandomVariate interface {
	Rand() float64
}

type Limits map[string]Interval

func (l Limits) String() string {
	var s string
	var i int
	for n, v := range l {
		i++
		// u2209 not element of...
		s += n + " \u2208 " + v.String()
		if i != len(l) {
			s += ", "
		}
	}

	return s
}

// https://en.wikipedia.org/wiki/Interval_(mathematics)
type Interval struct {
	Lower, Upper         float64
	LowerOpen, UpperOpen bool
}

// Check if x is within the interval.
func (i Interval) IsWithinInterval(x float64) bool {
	// disregard NaNs.
	if math.IsNaN(x) {
		return false
	}

	// If the lower limit is -∞, we are always within bounds.
	if !math.IsInf(i.Lower, -1) {
		if i.LowerOpen {
			if x <= i.Lower {
				return false // x must be > lower bound
			}
		}

		if x < i.Lower {
			return false // x must be >= lower bound
		}
	}

	// If the upper limit is ∞, we are always within bounds.
	if !math.IsInf(i.Upper, 1) {
		if i.UpperOpen {
			if x >= i.Upper {
				return false // x must be < upper bound
			}
		}

		if x > i.Upper {
			return false // x must be <= upper bound
		}

	}

	return true
}

// Check if length is equal the interval.
func (i Interval) IsEqualLength(len float64) bool {
	// disregard NaNs.
	if math.IsNaN(len) {
		return false
	}

	a := i.Lower
	b := i.Upper

	if i.LowerOpen {
		a -= math.SmallestNonzeroFloat64
	}

	if i.UpperOpen {
		b -= math.SmallestNonzeroFloat64
	}

	return math.Abs(a-b) == len
}

func (i Interval) String() string {
	l_temp := fmt.Sprintf("%v", i.Lower)

	if math.IsInf(i.Lower, 1) {
		l_temp = "∞"
	}

	if math.IsInf(i.Lower, -1) {
		l_temp = "-∞"
	}

	u_temp := fmt.Sprintf("%v", i.Upper)

	if math.IsInf(i.Upper, -1) {
		u_temp = "-∞"
	}

	if math.IsInf(i.Upper, 1) {
		u_temp = "∞"
	}

	str := "(" + l_temp + "," + u_temp + ")"

	if !i.LowerOpen {
		str = "[" + str[1:]
	}

	if !i.UpperOpen {
		str = str[:len(str)-1] + "]"
	}

	return str
}
