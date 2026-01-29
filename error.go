package terror

import (
	"fmt"
	"runtime"
	"strings"
)

var Depth = 32

// Optional:  Use init to configure depth of the trace stack
func Init(depth int) {
	Depth = depth
}

// create trace stack
func stack() []uintptr {
	pc := make([]uintptr, Depth)
	n := runtime.Callers(3, pc[:])
	return pc[:n]
}

// TracedError is wrapper for err which helps err keep track of the err
type TracedError struct {
	Err     error
	Message string
	stack   []uintptr
}

// Creates new TraceError
func New(err error) *TracedError {
	if err == nil {
		return nil
	}

	terr := &TracedError{
		Err:   err,
		stack: stack(),
	}
	return terr
}

// Newf creates error with format, use fmt.Errorf inside
func Newf(format string, args ...any) *TracedError {
	return &TracedError{
		Err:   fmt.Errorf(format, args...),
		stack: stack(),
	}
}

// Wraps an error with message
func Wrap(err error, msg string) *TracedError {
	if err == nil {
		return nil
	}

	terr := &TracedError{
		Err:     err,
		Message: msg,
		stack:   stack(),
	}
	return terr
}

// Wraps an error with formated message
func Wrapf(err error, format string, args ...any) *TracedError {
	if err != nil {
		return nil
	}

	terr := &TracedError{
		Err:     err,
		Message: fmt.Sprintf(format, args...),
		stack:   stack(),
	}
	return terr
}

// return error with message if exisits
func (t *TracedError) Error() string {
	if t == nil {
		return ""
	}
	b := strings.Builder{}
	if t.Message != "" {
		fmt.Fprintf(&b, "%s: ", t.Message)
	}
	fmt.Fprint(&b, t.Err)
	return b.String()
}

// Unwrap the error that is wraped :)
func (t *TracedError) Unwrap() error {
	return t.Err
}

// Trace return trace for the error
func (t *TracedError) Trace() string {
	if t == nil {
		return ""
	}
	b := strings.Builder{}
	frames := runtime.CallersFrames(t.stack)
	for {
		f, nxt := frames.Next()
		fmt.Fprintf(&b, "%s\n\t%s:%d\n", f.Function, f.File, f.Line)
		if !nxt {
			return b.String()
		}
	}
}
