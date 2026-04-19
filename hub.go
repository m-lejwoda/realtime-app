package main

type Hub struct {
	clients map[*Client]bool
	broadcast chan []byte
	register chan *Client
	unregister chan *Client
}

// func (h *Hub) addClient(c *Client){
// 	h.clients[c] = true
// }

// func (h *Hub) removeClient(c *Client){
// 	h.clients[c] = false
// }

func (h *Hub) registerClient(c *Client){
	h.register <- c
}


func (h *Hub) unregisterClient(c *Client){
	h.unregister <- c
}

func (h *Hub) broadcastMessage(message []byte){
	for client := range h.clients{
		client.send <- message
	}
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
			for c := range h.clients{
				c.send <- msg
			}
		}
	}
}