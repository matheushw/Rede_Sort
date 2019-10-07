package main

import (
    "net"
    "fmt"
    "tcp"
    "os"
    // "strings"
    // "encoding/binary"
    // "utils"
)

func main() {

    // server, err:= net.Listen("tcp", ":8000")
    // if server == nil {
    //     fmt.Printf("Couldn't start listening: %v\n", err)
    //     panic(err)
    // }

    // conns := tcp.ClientConns(server)

    // for {
    //     go tcp.HandleConn(<-conns)
    // }

    server, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error listetning: ", err)
		os.Exit(1)
	}
    defer server.Close()
    
    
    fmt.Println("Server started! Waiting for connections...")
    // utils.RemoveLine("dummyfile.txt", 4)
    
    
	// for {
    //     connection, err := server.Accept()
	// 	if err != nil {
    //         fmt.Println("Error: ", err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Println("Client connected")
	// 	go tcp.SendFile(connection)
    // }

    connection, err := server.Accept()

    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }

    fmt.Println("Client connected")

    tf, _ := os.Open("lista.txt")
    tcp.SendChunk(connection, tf, 5)

}