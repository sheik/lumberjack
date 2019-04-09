package main

import (
	"flag"
	"fmt"
	"github.com/sheik/lumberjack"
	"os"
	"strings"
)

func main() {
	tags := flag.String("tags", "", "comma separated list of tags")
	flag.Parse()

	var file *os.File
	var err error

	file = os.Stdin

	if len(flag.Args()) > 0 {
		file, err = os.Open(flag.Args()[0])
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	scanner := lumberjack.NewLumberjackScanner(file, strings.Split(*tags, ",")...)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
