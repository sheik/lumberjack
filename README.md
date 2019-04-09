# lumberjack
Lumberjack is a tag logging library. It allows simple logging configuation which allows
for logging of messages with tags. It then provides an interface for reading through a log
based on tags.

To install and get tcat:

    go install ./...

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


See LICENSE for licensing information

(C) 2019 Jeff Aigner
