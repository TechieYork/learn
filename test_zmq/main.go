package main

import (
    "fmt"
    "github.com/pebbe/zmq3"
)

func main() {
    fmt.Println("test zmq")

    server, _ := zmq3.NewSocket(zmq3.DEALER)
    defer server.Close()
}

