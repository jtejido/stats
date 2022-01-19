package test

import (
	"fmt"
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/ggsl/sys"
	"math"
	"os"
	"strconv"
	"testing"
)

var (
	tests   = 0
	passed  = 0
	failed  = 0
	verbose uint64
)

func initialise() {
	p := os.Getenv("STAT_TEST_VERBOSE")
	var err error

	/* 0 = show failures only (we always want to see these) */

	if p == "" { /* environment variable is not set */
		return
	}

	verbose, err = strconv.ParseUint(p, 10, 64)

	if err != nil {
		verbose = 0
	}
}

func update(s int) {
	tests++
	if s == 0 {
		passed++
	} else {
		failed++
	}
}

func Test(t *testing.T, status int, test_description string) {
	if tests == 0 {
		initialise()
	}

	update(status)

	if status != 0 || verbose != 0 {
		if status == 0 {
			fmt.Printf("PASS: %s", test_description)
		} else {
			s := fmt.Sprintf("FAIL: %s", test_description)
			if verbose == 0 {
				s += fmt.Sprintf(" [%v]", tests)
			}

			t.Errorf("%s", s)
		}

		fmt.Printf("\n")
	}
}
