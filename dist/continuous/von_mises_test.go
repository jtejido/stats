package continuous

import (
	"github.com/jtejido/stats"
	"math"
	"strconv"
	"testing"
)

// https://keisan.casio.com/exec/system/14550813435110
func TestVonMisesProbability(t *testing.T) {
	tol := 0.0000001

	cases := []struct {
		x, μ, κ, expected float64
	}{
		{1, 2, 3, 0.1649228},
		{.874564, .98786, .98675, 0.3370634},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b, _ := NewVonMises(c.μ, c.κ, DefaultCircularSupport)

			res := b.Probability(c.x)
			run_test(t, res, c.expected, tol, "VonMisesProbability")

		})
	}
}

// https://keisan.casio.com/exec/system/14550813435110
// 0≦μ≦2π
func TestVonMisesDistribution(t *testing.T) {
	tol := 0.0000001
	support := stats.Interval{0, 2 * math.Pi, false, false}
	cases := []struct {
		x, μ, κ, expected float64
	}{
		{1, 2, 3, 0.05559866153476771038236},
		{.874564, .98786, .98675, 0.2517594817460086731542},
		{5, 1, 5.86, 0.988050025449279526715},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b, _ := NewVonMises(c.μ, c.κ, support)

			res := b.Distribution(c.x)
			run_test(t, res, c.expected, tol, "VonMisesDistribution")

		})
	}
}
