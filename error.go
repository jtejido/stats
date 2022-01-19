package stats

import (
	"fmt"
	"github.com/jtejido/stats/err"
	"runtime"
)

var (
	notImplementedError = "Method Not Implemented"
)

func DomainErrorf(msg string, args ...interface{}) {
	errno := err.EDOM
	reason := fmt.Sprintf(msg, args...)
	_, file, line, _ := runtime.Caller(1)
	err.HandleError(reason, file, line, errno)
}

func LossErrorf(msg string, args ...interface{}) {
	errno := err.ELOSS
	reason := fmt.Sprintf(msg, args...)
	_, file, line, _ := runtime.Caller(1)
	err.HandleError(reason, file, line, errno)
}

func RangeErrorf(msg string, args ...interface{}) {
	errno := err.ERANGE
	reason := fmt.Sprintf(msg, args...)
	_, file, line, _ := runtime.Caller(1)
	err.HandleError(reason, file, line, errno)
}

func NotImplementedError() {
	errno := err.EUNIMPL
	reason := notImplementedError
	_, file, line, _ := runtime.Caller(1)
	err.HandleError(reason, file, line, errno)
}
