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

func (c *Client) readMessage(h *Hub){
	fmt.Println("readMessage")
}