package lumberjack

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Tag string
type Tags []string
type TagMap map[string]bool

// Basic logger type, contains a list
// of destinations
type Lumberjack struct {
	outputFiles []io.Writer
}

// NewLogger takes a list of io.Writers an returns a new
// logging object that can be used to lo tags
func NewLogger(files ...io.Writer) Lumberjack {
	return Lumberjack{outputFiles: files}
}

// Log a message with given tags
func (logger *Lumberjack) Log(message string, tags... string) {
	for _, log := range logger.outputFiles {
		fmt.Fprintf(log, ":%s::%s\n", strings.Join(tags, ":"), message)
	}
}

// LumberjackScanner provides a way to open
// a log file and go through the lines, matching
// only on specific tags. It is based on the
// bufio.Scanner type
type LumberjackScanner struct {
	scanner *bufio.Scanner
	text    string
	tagMap  TagMap
}

// NewLumberjackScanner returns a new scanner that will
// search for "tags"
func NewLumberjackScanner(r io.Reader, tags... string) LumberjackScanner {
	s := LumberjackScanner{bufio.NewScanner(r), "", make(map[string]bool)}
	for _, tag := range tags {
		s.tagMap[tag] = true
	}
	return s
}

// Running Scan() will cause the scanner to read
// from the input until a line matching our tag
// format is found and filtered out
func (logger *LumberjackScanner) Scan() bool {
	for logger.scanner.Scan() {
		text := logger.scanner.Text()
		parts := strings.Split(text, "::")

		// continue if the line doesn't fit format
		if len(parts) < 2 {
			continue
		}

		tags := strings.Split(parts[0], ":")

		for _, tag := range tags {
			_, ok := logger.tagMap[tag]
			if ok {
				logger.text = strings.Join(parts[1:], "")
				return true
			}
		}
	}
	return false
}

// Text() returns the last line found by the scanner
func (logger *LumberjackScanner) Text() string {
	return logger.text
}
