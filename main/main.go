package main

import (
    "net"
    "fmt"
    "tcp"
    // "utils"
)

func main() {

    server, err:= net.Listen("tcp", ":8000")
    if server == nil {
        fmt.Printf("Couldn't start listening: %v\n", err)
        panic(err)
    }

    conns := tcp.ClientConns(server)

    for {
        go tcp.HandleConn(<-conns)
    }
}