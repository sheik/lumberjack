package lumberjack

// alphaReader is borrowed and modified from:
// https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185

import (
	"io"
	"testing"
)

func TestAlphaReader(t *testing.T) {
	output := ""
	reader := newAlphaReader("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		output += string(p[:n])
	}

	if output != "HelloItsamwhereisthesun" {
		t.Error("Invalid AlphaReader")
	}
}
