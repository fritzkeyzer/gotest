package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
    var conn net.Conn
    defer func() {
        conn.Close()
    }()

    p :=  make([]byte, 2048)
    conn, err := net.Dial("udp", "0.0.0.0:1234")
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }
    fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
    _, err = bufio.NewReader(conn).Read(p)
    if err == nil {
        fmt.Printf("%s\n", p)
    } else {
        fmt.Printf("Some error %v\n", err)
    }

}