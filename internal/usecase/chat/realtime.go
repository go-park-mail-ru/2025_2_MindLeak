package usecase

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	repository "github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/chat"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/chat/dto"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	hub         *Hub
	conn        *websocket.Conn
	userID      uuid.UUID
	currentRoom uuid.UUID
	mu          sync.Mutex
	send        chan dto.ServerEvent
}

type Hub struct {
	clients    map[uuid.UUID]*Client
	rooms      map[uuid.UUID]map[uuid.UUID]*Client
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

type RealtimeUsecase interface {
	Register(conn *websocket.Conn, userID uuid.UUID, rooms []uuid.UUID)
	BroadcastToRoom(roomID uuid.UUID, event dto.ServerEvent)
}

type realtimeUsecase struct {
	hub      *Hub
	chatRepo repository.ChatRepository
}

func NewRealtimeUsecase(chatRepo repository.ChatRepository) RealtimeUsecase {
	h := &Hub{
		clients:    make(map[uuid.UUID]*Client),
		rooms:      make(map[uuid.UUID]map[uuid.UUID]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	rt := &realtimeUsecase{hub: h, chatRepo: chatRepo}
	go h.run()
	return rt
}

func (rt *realtimeUsecase) Register(conn *websocket.Conn, userID uuid.UUID, rooms []uuid.UUID) {
	c := &Client{
		hub:    rt.hub,
		conn:   conn,
		userID: userID,
		send:   make(chan dto.ServerEvent, 256),
	}
	rt.hub.register <- c

	rt.hub.mu.Lock()
	for _, rid := range rooms {
		if rt.hub.rooms[rid] == nil {
			rt.hub.rooms[rid] = make(map[uuid.UUID]*Client)
		}
		rt.hub.rooms[rid][userID] = c
	}
	rt.hub.mu.Unlock()

	go c.writePump()
	go c.readPump(rt)
}

func (rt *realtimeUsecase) BroadcastToRoom(roomID uuid.UUID, event dto.ServerEvent) {
	event.Ts = time.Now()
	rt.hub.mu.RLock()
	defer rt.hub.mu.RUnlock()
	if clients, ok := rt.hub.rooms[roomID]; ok {
		for _, client := range clients {
			select {
			case client.send <- event:
			default:
				close(client.send)
			}
		}
	}
}

// ========== Горутины ==========

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case event, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, _ := c.conn.NextWriter(websocket.TextMessage)
			json.NewEncoder(w).Encode(event)
			w.Close()
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			c.conn.WriteMessage(websocket.PingMessage, nil)
		}
	}
}

func (c *Client) readPump(rt *realtimeUsecase) {
	defer func() {
		rt.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		var msg dto.ClientMessage
		if err := c.conn.ReadJSON(&msg); err != nil {
			break
		}

		switch msg.Type {
		case "select_room":
			var p struct {
				RoomID uuid.UUID `json:"room_id"`
			}
			if json.Unmarshal(msg.Payload, &p) == nil {
				c.mu.Lock()
				c.currentRoom = p.RoomID
				c.mu.Unlock()
			}

		case "send_message":
			var p struct {
				Text string `json:"text"`
			}
			if json.Unmarshal(msg.Payload, &p) != nil || p.Text == "" {
				continue
			}
			c.mu.Lock()
			roomID := c.currentRoom
			c.mu.Unlock()
			if roomID == uuid.Nil {
				continue
			}
			ok, _ := rt.chatRepo.IsUserInRoom(context.Background(), roomID, c.userID)
			if !ok {
				continue
			}

			message := &models.Message{
				RoomID:    roomID,
				UserID:    c.userID,
				Text:      p.Text,
				CreatedAt: time.Now(),
			}
			rt.chatRepo.CreateMessage(context.Background(), message)
			rt.BroadcastToRoom(roomID, dto.ServerEvent{Type: "message", Data: message})
		}
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.userID] = client
			h.mu.Unlock()
		case client := <-h.unregister:
			h.mu.Lock()
			delete(h.clients, client.userID)
			for _, room := range h.rooms {
				delete(room, client.userID)
			}
			h.mu.Unlock()
			close(client.send)
		}
	}
}
