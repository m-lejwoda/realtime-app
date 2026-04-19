package main

import(
	"sync"
)

type HubManager struct {
	hubs []*Hub
	mu   sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		broadcast: make(chan []byte),
		register: make(chan *Client),
		unregister: make(chan *Client),
		clients: make(map[*Client]bool),
	}
}

func (m *HubManager) getRandomHub() *Hub {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, h := range m.hubs {
		if len(h.clients) < 2 {
			return h
		}
	}
	newHub := NewHub()
	go newHub.Run()
	m.hubs = append(m.hubs, newHub)
	return newHub
}

