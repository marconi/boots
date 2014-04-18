# Boots

[![Build Status](https://travis-ci.org/marconi/boots.svg?branch=master)](https://travis-ci.org/marconi/boots)

A supposedly [STOMP](http://stomp.github.io/) client.

### Installation

    $ go get github.com/marconi/boots

### Usage

    package main

    import (
        "log"

        "github.com/marconi/boots"
        "github.com/marconi/boots/client"
    )

    func main() {
        c := client.NewClient(
            "localhost:61613",
            "stomp", // vhost
            "admin", // login
            "admin", // passcode
        )
        if err := c.Connect(); err != nil {
            log.Fatal(err)
        }

        param := client.SendParam{
            Dest:      "/queue/hello-queue",
            Body:      "Hello!",
            Ctype:     "text/plain",
        }
        if err := c.Send(param); err != nil {
            log.Fatal(err)
        }
    }

### Notes

Only tested against RabbitMQ 3.2.4 and only supports STOMP 1.2

### License

http://marconi.mit-license.org/
