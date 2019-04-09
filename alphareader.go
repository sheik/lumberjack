package lumberjack

import "io"

// alphaReader is borrowed and modified from:
// https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185

type alphaReader struct {
	src string
	cur int
}

func newAlphaReader(src string) *alphaReader {
	return &alphaReader{src: src}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	x := len(a.src) - a.cur
	n, bound, m := 0, 0, 0
	if x >= len(p) {
		bound = len(p)
	} else if x <= len(p) {
		bound = x
	}

	buf := make([]byte, bound)
	for m < bound {
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[n] = char
			n++
		}
		m++
		a.cur++
	}
	copy(p, buf)
	return n, nil
}
