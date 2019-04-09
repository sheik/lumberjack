package main

import (
	"flag"
	"fmt"
	"github.com/sheik/lumberjack"
	"os"
	"strings"
)

func getFile(args []string) (*os.File, error) {
	var file *os.File
	var err error

	file = os.Stdin

	if len(flag.Args()) > 0 {
		file, err = os.Open(flag.Args()[0])
		if err != nil {
			return nil, err
		}
	}

	return file, nil
}

func main() {
	tags := flag.String("tags", "", "comma separated list of tags")
	flag.Parse()

	file, err := getFile(flag.Args())

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}

	scanner := lumberjack.NewLumberjackScanner(file, strings.Split(*tags, ",")...)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	os.Exit(0)
}
