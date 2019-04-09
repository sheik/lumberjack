package main

import (
	"os"
	"testing"
)

func TestGetFile(t *testing.T) {
	file, err := getFile([]string{})

	if err != nil {
		t.Error("could not get os.Stdin from getFile()")
	}

	if file != os.Stdin {
		t.Error("did not receive os.Stdin from getFile()")
	}
}
