package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct{
	hub *Hub
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) sendMessage(h *Hub){


}

func (c *Client) addToBroadcast(h *Hub){
	for{
		_, msg, err := c.conn.ReadMessage()
		if err != nil{
			fmt.Println("err", err)
		}
		fmt.Println(msg)
		h.broadcast <- msg
	}
}

func (c *Client) sendToClient(){
	for{
		for msg := range c.send{
			fmt.Println("Sending")
			c.conn.WriteMessage(websocket.TextMessage, msg)
		}
	}
}