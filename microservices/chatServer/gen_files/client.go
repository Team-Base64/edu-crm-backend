// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	//send chan []byte
	send chan *MessageWebsocket
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {

		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println(err)
			}
			break
		}
		var req MessageWebsocket
		err = json.Unmarshal(message, &req)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(req)
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		//c.hub.Broadcast <- message
		//c.hub.Broadcast <- &MessageWebsocket{Text: req.Text, ChatID: req.ChatID}
		c.hub.chats[req.ChatID] = c
		c.hub.clientChats[c] = append(c.hub.clientChats[c], req.ChatID)
		c.hub.Broadcast <- &req
		c.hub.MessagesToTGBot <- &req
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println(err)
				return
			}
			req, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
				return
			}

			w.Write(req)
			c.hub.chats[message.ChatID] = c
			c.hub.clientChats[c] = append(c.hub.clientChats[c], message.ChatID)

			// // Add queued chat messages to the current websocket message.
			// n := len(c.send)
			// for i := 0; i < n; i++ {
			// 	w.Write(newline)
			// 	w.Write(<-c.send)
			// }

			if err := w.Close(); err != nil {
				log.Println(err)
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client := &Client{hub: hub, conn: conn, send: make(chan *MessageWebsocket)}
	client.hub.register <- client
	//chats:= GetAllUserChats
	curChats := []int32{1, 2}
	for _, ch := range curChats {
		hub.chats[ch] = client
		hub.clientChats[client] = append(hub.clientChats[client], ch)
	}

	//connect := &Connection{ws: conn, send: make(chan MessageWebsocket, 256)}
	//conn := &connection{send: make(chan models.Message, 256), ws: ws}

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.

	go client.writePump()
	go client.readPump()
}
