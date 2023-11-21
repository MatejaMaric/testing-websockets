package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MatejaMaric/testing-websockets/pkg/messages"
	"golang.org/x/net/websocket"
)

func EchoServer(ws *websocket.Conn) {
	enc := json.NewEncoder(ws)
	tick := time.NewTicker(time.Second)

	for range tick.C {
		msg := messages.Msg{
			Value: time.Now().String(),
		}

		if err := enc.Encode(msg); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	http.Handle("/time", websocket.Handler(EchoServer))
	if err := http.ListenAndServe(":12345", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
