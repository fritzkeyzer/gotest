package main

import (
	"fmt"
	"net"
	"time"
)


func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
    msg := "hello client: " + time.Now().Format("2006-01-02 15:04:05.999 -0700 MST") + "..."

    _,err := conn.WriteToUDP([]byte(msg), addr)
    if err != nil {
        fmt.Printf("Couldn't send response %v", err)
    }
}


func main() {
    p := make([]byte, 4096)
    addr := net.UDPAddr{
        Port: 1234,
        IP: net.ParseIP("0.0.0.0"),
    }
    ser, err := net.ListenUDP("udp", &addr)
    if err != nil {
        fmt.Printf("Some error %v\n", err)
        return
    }
    for {
        n, remoteaddr, err := ser.ReadFromUDP(p)
        
        if err !=  nil {
            fmt.Printf("Some error  %v", err)
            continue
        }
        
        fmt.Printf("RX[%v]: '%s'\n", remoteaddr, p[0:n])

        go sendResponse(ser, remoteaddr)
    }
}