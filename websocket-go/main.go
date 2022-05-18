package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		upgrade := websocket.Upgrader{} // use default options
		ws, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
			return
		}
		// create client, run processes, and add to hub
		cl := NewClient(ws)
		go cl.ReceiveMessages()
		go cl.SendMessages()
		//s.lobby.add <- cl

		cl.send <- []byte("hello new client")
	})

	err := http.ListenAndServe("localhost:5000", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}