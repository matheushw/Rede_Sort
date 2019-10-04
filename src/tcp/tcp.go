package tcp

import (
	"fmt"
	"net"
    "bufio"
)

func ClientConns(listener net.Listener) chan net.Conn {
    ch := make(chan net.Conn)
    go func() {
        for {
            client, err := listener.Accept()
            if client == nil {
                fmt.Printf("Couldn't accept: %v\n", err)
                continue
            }
            fmt.Printf("[One connection open]\n")
            ch <- client
        }
    }()
    return ch
}

func HandleConn(client net.Conn) {
    b := bufio.NewReader(client)
    for {
        line, err := b.ReadBytes('\n')
        if err != nil { 
            break
        }
        client.Write(line)
        fmt.Printf("[%v sent:] ", client.RemoteAddr())
        fmt.Printf("%s\n", line)
    }
}


