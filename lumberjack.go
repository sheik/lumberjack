// Copyright 2019 Jeff Aigner. All rights reserved.
// Use of this code is governed by the MIT license
// found in the LICENSE file

// Package lumberjack is a logger that allows logging
// and filtering based on tags.
//
// It is intended to be used to help with debugging of
// systems by providing an easy way to get different
// views of your system log
package lumberjack

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type tagMap map[string]bool

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
func (logger *Lumberjack) Log(message string, tags ...string) {
	for _, log := range logger.outputFiles {
		fmt.Fprintf(log, ":%s::%s\n", strings.Join(tags, ":"), message)
	}
}

// LumberjackScanner provides a way to open
// a log file and go through the lines, matching
// only on specific tags. It is based on the
// bufio.Scanner type. Use NewLumberjackScanner() to
// create a new LumberjackScanner object
type LumberjackScanner struct {
	scanner *bufio.Scanner
	text    string
	tagMap  tagMap
}

// NewLumberjackScanner returns a new scanner that will
// search for "tags"
func NewLumberjackScanner(r io.Reader, tags ...string) LumberjackScanner {
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
