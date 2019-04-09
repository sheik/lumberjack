package main

import (
	"os"
	"reflect"
	"testing"
)

type lines []string

func TestLumberjackInit(t *testing.T) {
	logger := NewLogger()
	logger.log("My first log message", Tags{"first", "info", "alert"})

	logger = NewLogger(os.Stdout)
	logger.log("My first log message", Tags{"first", "info", "alert"})
	logger.log("My second log message", Tags{"second", "info", "debug"})
}

func TestLumberjackScanner(t *testing.T) {
	r, err := os.Open("test.log")
	defer r.Close()
	if err != nil {
		t.Error(err)
	}

	var output lines
	scanner := NewLumberjackScanner(r, Tags{"first"})
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	if !reflect.DeepEqual(output, lines{"my first log"}) {
		t.Error("Unexpected output")
	}

	r2, err := os.Open("test.log")
	defer r2.Close()
	if err != nil {
		t.Error(err)
	}

	output = lines{}
	scanner = NewLumberjackScanner(r2, Tags{"network"})
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	networkLines := lines{"I like network cards", "eth0 is up"}
	if !reflect.DeepEqual(output, networkLines) {
		t.Error("Unexpected output:", output)
	}
}

func TestLumberjackScannerMultipleTags(t *testing.T) {
	r, err := os.Open("test.log")
	defer r.Close()
	if err != nil {
		t.Error(err)
	}

	var output lines
	scanner := NewLumberjackScanner(r, Tags{"first", "second"})
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	if !reflect.DeepEqual(output, lines{"my first log", "my second log", "my third log"}) {
		t.Error("Unexpected output:", output)
	}
}
