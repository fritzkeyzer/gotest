package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)


type Client struct{
	ws *websocket.Conn
	send chan[]byte
}

func NewClient(ws *websocket.Conn) *Client{
	cl := Client{
		ws: ws,
		send: make(chan []byte),
	}

	//

	return &cl
}

func (cl *Client) ReceiveMessages() {
	defer func() {
		fmt.Println("Client.RecieveMessages() goroutine stopping")
		//cl.HandlePlayerExit(nil)
		cl.ws.Close()
	}()
	for {
		// read message
		_, message, err := cl.ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		// log message received
		fmt.Println("client message received:")
		fmt.Println(string(message))
		//ConsoleLogJsonByteArray(message)
		


	}
}

func (cl *Client) SendMessages() {
	defer func() {
		fmt.Println("Client.SendMessages() goroutine stopping")
	}()
	for message := range cl.send {
		//SendJsonMessage(cl.ws, message)
		cl.ws.WriteMessage(1, message)
	}
}