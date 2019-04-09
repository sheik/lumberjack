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

	defer logfile.Close()

	logger := lumberjack.NewLogger(os.Stdout, logfile)

	logger.Log("The network is down", "network", "debug")
}