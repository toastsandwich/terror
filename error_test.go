package terror

import (
	"errors"
	"fmt"
	"testing"
)

var (
	ErrExampleOne = errors.New("error: One")
)

func Fa() error {
	return ErrExampleOne
}

func Fc() error {
	return Fb()
}

func Fb() error {
	return Fa()
}

func TestTracedError(t *testing.T) {
	err := New(Fc())
	fmt.Println(err)
	fmt.Println(err.Trace())
	if !errors.Is(err, ErrExampleOne) {
		t.Fatalf("Error does not match expected=%v got=%v", ErrExampleOne, err)
	}
}
