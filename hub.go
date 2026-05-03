package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	Date time.Time
	User string
	Content string
}
func newMessage(content string) *Message{
	message := Message{}
	message.Date = time.Now()
	message.User = "Anonymous"
	message.Content = content
	return &message
}

type Hub struct {
	clients map[*Client]bool
	broadcast chan []byte
	register chan *Client
	unregister chan *Client
}

func (h *Hub) registerClient(c *Client){
	h.register <- c
}


func (h *Hub) unregisterClient(c *Client){
	h.unregister <- c
}

func (h *Hub) Run(){
	for{
		select{
		case c := <-h.register:
			if len(h.clients) <= 2{
				h.clients[c] = true
			}
		case c := <-h.unregister:
			delete(h.clients, c)
		case msg := <-h.broadcast:
			fmt.Printf("Hub %p odbiera wiadomość. Rozsyłam do %d klientów\n", h, len(h.clients))
			fmt.Println(string(msg))
			new_msg := newMessage(string(msg))
			converted_msg, err := json.Marshal(new_msg)
			if err != nil{
				fmt.Println(err)
			}
			for c := range h.clients{
				c.send <- converted_msg
			}
		}
	}
}