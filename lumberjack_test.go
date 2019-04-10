// Copyright 2019 Jeff Aigner. All rights reserved.
// Use of this code is governed by the MIT license
// found in the LICENSE file

package lumberjack

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

type lines []string

func TestLumberjackScanner(t *testing.T) {
	r, err := os.Open("test.log")
	if err != nil {
		t.Error(err)
	}
	defer r.Close()

	var output lines
	scanner := NewScanner(r, "first")
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	if !reflect.DeepEqual(output, lines{"my first log"}) {
		t.Error("Unexpected output")
	}

	r2, err := os.Open("test.log")
	if err != nil {
		t.Error(err)
	}
	defer r2.Close()

	output = lines{}
	scanner = NewScanner(r2, "network")
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	networkLines := lines{"I like network cards", "eth0 is up"}
	if !reflect.DeepEqual(output, networkLines) {
		t.Error("Unexpected output:", output)
	}
}

func TestScannerMultipleTags(t *testing.T) {
	r, err := os.Open("test.log")
	defer r.Close()
	if err != nil {
		t.Error(err)
	}

	var output lines
	scanner := NewScanner(r, "first", "second")
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	if !reflect.DeepEqual(output, lines{"my first log", "my second log", "my third log"}) {
		t.Error("Unexpected output:", output)
	}
}

// see test.log in the source to get matching output
func ExampleNewScanner() {
	file, err := os.Open("test.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := NewScanner(file, "tag1", "tag2")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// Output:
	// my second log
	// my third log
}

func ExampleNewLogger_stdout() {
	logger := NewLogger(os.Stdout)

	logger.Log("Simple log message, with no tags")
	logger.Log("Another message, this time with tags", "tag1", "tag2")
	logger.Log("Another message, this time with only one tag", "tag1")
	// Output:
	// ::Simple log message, with no tags
	// tag1:tag2::Another message, this time with tags
	// tag1::Another message, this time with only one tag
}

func ExampleNewLogger_tags() {
	logger := NewLogger(os.Stdout).Tags("myApp", "TestFunction")
	logger.Log("Running")
	// Output:
	// myApp:TestFunction::Running
}

func ExampleNewLogger_file() {
	filename := "output.log"
	logfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	defer logfile.Close()

	logger := NewLogger(logfile)

	logger.Log("Simple test message")
	logger.Log("Another message, this time with tags", "tag1", "tag2")
	logger.Log("Another message, this time with only one tag", "tag1")

	// Test code for verifying output
	f, err := os.Open("output.log")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// Output:
	// ::Simple test message
	// tag1:tag2::Another message, this time with tags
	// tag1::Another message, this time with only one tag
}
