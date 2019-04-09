Lumberjack
==========
Lumberjack is a tag logging library. It allows simple logging configuation which allows
for logging of messages with tags. It then provides an interface for reading through a log
based on tags.

To install and get tcat:

    go install ./...
    
Example Code:
    
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

Example:

    user@host$ cat test.log 
    first:info:test::my first log
    second:debug:test::my second log
    network:info::I like network cards
    second:debug:test::my third log
    network bogus line
    test::my first log
    network::eth0 is up
    
    user@host$ cat test.log | tcat -tags "network"
    I like network cards
    eth0 is up
    
    user@host$ tcat -tags network,first test.log 
    my first log
    I like network cards
    eth0 is up



See LICENSE for licensing information

(C) 2019 Jeff Aigner