package main

import (
	"encoding/json"
	"log"

	"github.com/MatejaMaric/testing-websockets/pkg/messages"
	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://localhost/"
	url := "ws://localhost:12345/time"

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(ws)

	for {
		var msg messages.Msg
		if err := dec.Decode(&msg); err != nil {
			log.Fatal(err)
		}
		log.Println("Received: ", msg.Value)
	}
}
