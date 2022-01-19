package continuous

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestWrappedNormalProbability(t *testing.T) {
	tol := 0.0000001
	sn, _ := NewNormal(0, 1)
	cases := []struct {
		θ, expected float64
	}{
		{.5, 0.3520653},
		{2.432, 0.02096843},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b, _ := NewWrapped(sn, 1000, DefaultCircularSupport)

			res := b.Probability(c.θ)
			run_test(t, res, c.expected, tol, "WrappedNormalProbability")

		})
	}
}

func TestWrappedNormalDistribution(t *testing.T) {
	tol := 0.0000001
	sn, _ := NewNormal(0, 1)
	cases := []struct {
		θ, expected float64
	}{
		{.5, 0.6914625},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b, _ := NewWrapped(sn, 1000, DefaultCircularSupport)

			res := b.Distribution(c.θ)
			run_test(t, res, c.expected, tol, "WrappedNormalDistribution")

		})
	}
}

func BenchmarkWrappedNormalDistribution(b *testing.B) {
	sn, _ := NewNormal(0, 1)
	bd, _ := NewWrapped(sn, 1000, DefaultCircularSupport)
	rnd := rand.New(rand.NewSource(12345))
	for n := 0; n < b.N; n++ {
		bd.Distribution(rnd.NormFloat64())
	}
}
