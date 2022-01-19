package err

import (
	"log"
	// "os"
)

type (
	ErrorHandlerType  = func(reason, file string, line, gsl_errno int)
	StreamHandlerType = func(label, file string, line int, reason string)
)

var (
	errorHandler ErrorHandlerType = nil
)

type StatsError interface {
	Status() int
	Error() string // should embed 'error' interface, but running into errors compiling on Windows machines.
}

type statsError struct {
	status  int
	message string
}

func New(status int, text string) *statsError {
	return &statsError{status, text}
}

func (err *statsError) Status() int {
	return err.status
}

func (err *statsError) Error() string {
	return err.message
}

func HandleError(reason, file string, line, gsl_errno int) {
	if errorHandler != nil {
		errorHandler(reason, file, line, gsl_errno)
		return
	}

	StreamPrintf("ERROR", file, line, reason)
	log.Printf("Default Stats error handler invoked.\n")
	// SIGABRT = 6?
	// os.Exit(6)?
	panic(reason)
}

func SetErrorHandler(new_handler ErrorHandlerType) {
	errorHandler = new_handler
}

func SetErrorHandlerOff() {
	errorHandler = NoErrorHandler
}

func NoErrorHandler(reason, file string, line int, gsl_errno int) {
	/* do nothing */
	return
}
