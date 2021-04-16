package main

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	err := errors.New("Hello")

	if err != nil {
		t.Logf("This is the generated error: %v", err)
		return
	}

	t.Fatal("Execution continued after testing for the error")
}
