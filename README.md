Lumberjack
==========
Lumberjack is a tag logging library. It allows simple logging configuation which allows
for logging of messages with tags. It then provides an interface for reading through a log
based on tags.

To use the lumberjack library in your program:

    go get github.com/sheik/lumberjack
    
After you install, you can run "godoc -http :8080" and browse for lumberjack to get more detailed documentation.
Or you can visit the [lumberjack documentation](https://godoc.org/github.com/sheik/lumberjack) online at godoc.org.

To install the filter program to filter the logs you produce:
    
    go get github.com/sheik/lumberjack/filter
    
Example Code:
    
    package main

    import (
        "github.com/sheik/lumberjack"
        "os"
    )

    func main() {
        // Create our log file for output
        logfile, err := os.OpenFile("output.log", os.O_RDWR|os.O_CREATE, 0755)
        if err != nil {
            panic(err)
        }
        defer logfile.Close()

        // print to stdout and also to a log file
        logger := lumberjack.NewLogger(os.Stdout, logfile)

        // Make a log message with two tags: network, debug
        logger.Log("The network is down", "network", "debug")
    }

Example:

    user@host$ cat test.log 
    first:info:test::my first log
    second:debug:test::my second log
    network:info::I like network cards
    second:debug:test::my third log
    network bogus line
    test::my first log
    network::eth0 is up
    
    user@host$ cat test.log | filter -tags network
    I like network cards
    eth0 is up
    
    user@host$ filter -tags network,first test.log 
    my first log
    I like network cards
    eth0 is up



See LICENSE for licensing information

 ©2019 Jeff Aigner
