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

type Lumberjack struct {
	outputFiles []io.Writer
}

func NewLogger(files ...io.Writer) Lumberjack {
	return Lumberjack{outputFiles: files}
}

func (logger *Lumberjack) log(message string, tags Tags) {
	for _, log := range logger.outputFiles {
		fmt.Fprintf(log, "%s::%s\n", strings.Join(tags, ":"), message)
	}
}

type LumberjackScanner struct {
	scanner *bufio.Scanner
	text    string
	tagMap  TagMap
}

func NewLumberjackScanner(r io.Reader, tags Tags) LumberjackScanner {
	s := LumberjackScanner{bufio.NewScanner(r), "", make(map[string]bool)}
	for _, tag := range tags {
		s.tagMap[tag] = true
	}
	return s
}

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

func (logger *LumberjackScanner) Text() string {
	return logger.text
}

