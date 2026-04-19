package main

import (
    "github.com/gorilla/websocket"
)

type Client struct{
	hub *Hub
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) sendMessage(h *Hub){

}