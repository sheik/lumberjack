package main

import (
	"github.com/sheik/lumberjack"
	"os"
)

func consoleLogger() {
	logger := lumberjack.NewLogger(os.Stdout)

	logger.Log("Simple log message, with no tags")
	logger.Log("Another message, this time with tags", "tag1", "tag2")
	logger.Log("Another message, this time with only one tag", "tag1")
}

func logfileLogger(filename string) {
	logfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		panic(err)
	}

	defer logfile.Close()

	logger := lumberjack.NewLogger(logfile)

	logger.Log("Simple test message")
	logger.Log("Another message, this time with tags", "tag1", "tag2")
	logger.Log("Another message, this time with only one tag", "tag1")
}

func main() {
	consoleLogger()
	logfileLogger("output.log")
}
