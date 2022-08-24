package controllers

import (
	"log"
	"net/http"
	"time"
	"zeus/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// CORS
	CheckOrigin: func(r *http.Request) bool { return true },
}

var Messages [3]models.WebsocketMessage
var WebsocketStarted bool = false

func Websocket(c *gin.Context) {
	// Upgrades the HTTP server connection to the WebSocket protocol
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade failed: ", err)
		return
	}
	defer conn.Close()

	// Informs that the websocket has been activated
	WebsocketStarted = true

	for {
		// Real-time percents
		for i := range Messages {
			conn.WriteJSON(Messages[i])
		}

		// Avoid overload by pausing the goroutine for 0.1 seconds
		time.Sleep(time.Duration(1e+9))
	}
}
