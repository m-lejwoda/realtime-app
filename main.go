package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}
func main() {
	fmt.Println("Hello world")
	manager := &HubManager{
		hubs: make([]*Hub, 0),
	}
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(manager, w,r)
	})
	fmt.Println("Server has just started - port:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		fmt.Println("Server Error", err)
	}
}	


func serveWs(manager *HubManager, w http.ResponseWriter, r *http.Request){
	fmt.Println("serveWs")
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil{
		fmt.Println("Error with upgrade")
	}
	hub := manager.getRandomHub()

	client := &Client{
		hub: hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
	hub.register <- client
	fmt.Println("hub", hub.broadcast ,hub.clients)
	// go client.writePump()
	// go client.readPump()
}