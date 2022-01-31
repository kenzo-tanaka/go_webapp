package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// clientはチャットを行っている1人のユーザー
type client struct {
	socket *websocket.Conn // clientのためのWebsocket
	send   chan []byte     // sendはメッセージが送られるためのチャネル
	room   *room           // clientが参加しているチャットルーム
}

func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
