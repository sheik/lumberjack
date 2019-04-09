package main

import (
	"github.com/sheik/lumberjack"
	"os"
)

func main() {
	logfile, err := os.OpenFile("output.log", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		panic(err)
	}

	logger := lumberjack.NewLogger(logfile)

	logger.Log("The network is down", lumberjack.Tags{"network", "debug"})
}