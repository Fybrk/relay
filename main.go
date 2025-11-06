package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

type RelayServer struct {
	clients map[string]*websocket.Conn
	mu      sync.RWMutex
}

type Message struct {
	Type     string          `json:"type"`
	DeviceID string          `json:"device_id,omitempty"`
	Target   string          `json:"target,omitempty"`
	Data     json.RawMessage `json:"data,omitempty"`
}

func NewRelayServer() *RelayServer {
	return &RelayServer{
		clients: make(map[string]*websocket.Conn),
	}
}

func (r *RelayServer) HandleConnection(ws *websocket.Conn) {
	defer ws.Close()
	
	for {
		var msg Message
		if err := websocket.JSON.Receive(ws, &msg); err != nil {
			break
		}
		
		switch msg.Type {
		case "register":
			r.mu.Lock()
			r.clients[msg.DeviceID] = ws
			r.mu.Unlock()
			log.Printf("Device registered: %s", msg.DeviceID)
			
		case "relay":
			r.mu.RLock()
			targetConn := r.clients[msg.Target]
			r.mu.RUnlock()
			
			if targetConn != nil {
				websocket.JSON.Send(targetConn, msg)
			}
		}
	}
	
	// Clean up on disconnect
	r.mu.Lock()
	for deviceID, conn := range r.clients {
		if conn == ws {
			delete(r.clients, deviceID)
			log.Printf("Device disconnected: %s", deviceID)
			break
		}
	}
	r.mu.Unlock()
}

func main() {
	relay := NewRelayServer()
	
	http.Handle("/relay", websocket.Handler(relay.HandleConnection))
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	
	log.Println("Fybrk relay server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
