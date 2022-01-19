package err

import (
	"runtime"
)

const (
	FAILURE = iota - 1
	SUCCESS // sanity purpose
	EDOM
	ERANGE
	EFAULT
	EINVAL
	EFAILED
	EFACTOR
	ESANITY
	ENOMEM
	EBADFUNC
	ERUNAWAY
	EMAXITER
	EZERODIV
	EBADTOL
	ETOL
	EUNDRFLW
	EOVRFLW
	ELOSS
	EROUND
	EBADLEN
	ENOTSQR
	ESING
	EDIVERGE
	EUNSUP
	EUNIMPL
	ECACHE
	ETABLE
	ENOPROG
	ENOPROGJ
	ETOLF
	ETOLX
	ETOLG
	EOF
)

const (
	failure_msg  = "General failure"
	edom_msg     = "Input domain error, e.g sqrt(-1)"
	erange_msg   = "Output range error, e.g. exp(1e100)"
	efault_msg   = "Invalid pointer"
	einval_msg   = "Invalid argument supplied by user"
	efailed_msg  = "Generic failure"
	efactor_msg  = "Factorization failed"
	esanity_msg  = "Sanity check failed - shouldn't happen"
	enomem_msg   = "Malloc failed"
	ebadfunc_msg = "Problem with user-supplied function"
	erunaway_msg = "Iterative process is out of control"
	emaxiter_msg = "Exceeded max number of iterations"
	ezerodiv_msg = "Tried to divide by zero"
	ebadtol_msg  = "User specified an invalid tolerance"
	etol_msg     = "Failed to reach the specified tolerance"
	eundrflw_msg = "Underflow"
	eovrflw_msg  = "Overflow"
	eloss_msg    = "Loss of accuracy"
	eround_msg   = "Failed because of roundoff error"
	ebadlen_msg  = "Matrix, vector lengths are not conformant"
	enotsqr_msg  = "Matrix not square"
	esing_msg    = "Apparent singularity detected"
	ediverge_msg = "Integral or series is divergent"
	eunsup_msg   = "Requested feature is not supported by the hardware"
	eunimpl_msg  = "Requested feature not (yet) implemented"
	ecache_msg   = "Cache limit exceeded"
	etable_msg   = "Table limit exceeded"
	enoprog_msg  = "Iteration is not making progress towards solution"
	enoprogj_msg = "Jacobian evaluations are not improving the solution"
	etolf_msg    = "Cannot reach the specified tolerance in F"
	etolx_msg    = "Cannot reach the specified tolerance in X"
	etolg_msg    = "Cannot reach the specified tolerance in gradient"
	eof_msg      = "End of file"
)

//  call the error handler, and return the error
func Error(reason string, errno int) StatsError {
	_, file, line, _ := runtime.Caller(1)
	HandleError(reason, file, line, errno)
	return New(errno, reason)
}

//  call the error handler, and return the given value
func StatsErrorVal(reason string, errno int, value float64) float64 {
	_, file, line, _ := runtime.Caller(1)
	HandleError(reason, file, line, errno)
	return value
}

// for void functions which still need to generate an error
func StatsErrorVoid(reason string, errno int) {
	_, file, line, _ := runtime.Caller(1)
	HandleError(reason, file, line, errno)
	return
}

// suitable for out-of-memory conditions
func StatsErrorNull(reason string, errno int) {
	StatsErrorVal(reason, errno, 0)
}

// If multiple errors, select first non-Success one
func ErrorSelect(err ...StatsError) StatsError {
	for _, e := range err {
		if e != nil {
			return e
		}
	}

	return nil
}

// helpers
func Failure() StatsError {
	return New(FAILURE, failure_msg)
}

func Success() StatsError {
	return StatsError(nil)
}

func Domain() StatsError {
	return New(EDOM, edom_msg)
}

func Range() StatsError {
	return New(ERANGE, erange_msg)
}

func Fault() StatsError {
	return New(EFAULT, efault_msg)
}

func Invalid() StatsError {
	return New(EINVAL, einval_msg)
}

func Generic() StatsError {
	return New(EFAILED, efailed_msg)
}

func Factor() StatsError {
	return New(EFACTOR, efactor_msg)
}

func Sanity() StatsError {
	return New(ESANITY, esanity_msg)
}

func NoMemory() StatsError {
	return New(ENOMEM, enomem_msg)
}

func BadFunc() StatsError {
	return New(EBADFUNC, ebadfunc_msg)
}

func RunAway() StatsError {
	return New(ERUNAWAY, erunaway_msg)
}

func MaxIteration() StatsError {
	return New(EMAXITER, emaxiter_msg)
}

func ZeroDiv() StatsError {
	return New(EZERODIV, ezerodiv_msg)
}

func BadTolerance() StatsError {
	return New(EBADTOL, ebadtol_msg)
}

func Tolerance() StatsError {
	return New(ETOL, etol_msg)
}

func Underflow() StatsError {
	return New(EUNDRFLW, eundrflw_msg)
}

func Overflow() StatsError {
	return New(EOVRFLW, eovrflw_msg)
}

func Loss() StatsError {
	return New(ELOSS, eloss_msg)
}

func Round() StatsError {
	return New(EROUND, eround_msg)
}

func BadLength() StatsError {
	return New(EBADLEN, ebadlen_msg)
}

func NotSquare() StatsError {
	return New(ENOTSQR, enotsqr_msg)
}

func Singularity() StatsError {
	return New(ESING, esing_msg)
}

func Divergent() StatsError {
	return New(EDIVERGE, ediverge_msg)
}

func NotSupported() StatsError {
	return New(EUNSUP, eunsup_msg)
}

func NotImplemented() StatsError {
	return New(EUNIMPL, eunimpl_msg)
}

func CacheExceeded() StatsError {
	return New(ECACHE, ecache_msg)
}

func TableExceeded() StatsError {
	return New(ETABLE, etable_msg)
}

func NoProgress() StatsError {
	return New(ENOPROG, enoprog_msg)
}

func NoProgressJacobian() StatsError {
	return New(ENOPROGJ, enoprogj_msg)
}

func ToleranceF() StatsError {
	return New(ETOLF, etolf_msg)
}

func ToleranceX() StatsError {
	return New(ETOLX, etolx_msg)
}

func ToleranceGradient() StatsError {
	return New(ETOLG, etolg_msg)
}

func EndOfFile() StatsError {
	return New(EOF, eof_msg)
}
