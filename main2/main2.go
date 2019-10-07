package main

import (
    // "tcp"
    // "utils"
    "net"
    "fmt"
    // "tcp"
    // "os"
    "encoding/binary"
    // "utils"
)

func main() {
    
    connection, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}

    // bi := make([]int,8)
    var bi int64
    
    for i:=0;i<5;i++ {
        erro := binary.Read(connection, binary.BigEndian, &bi)
        if erro != nil {
            fmt.Println("erro:", erro)
        }
        // connection.Read(bi)
        fmt.Printf("%d\n",bi)
    }
}