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

var WebsocketRealTimeStarted bool = false
var WebsocketRealTimeChannel chan models.WebsocketRealTime

var WebsocketForWarningStarted bool = false
var WebsocketForWarningChannel chan models.WebsocketForWarning

func RealTimeWebsocket(c *gin.Context) {
	// Upgrades the HTTP server connection to the WebSocket protocol
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade failed: ", err)
		return
	}
	defer conn.Close()

	// Informs that the websocket has been activated
	WebsocketRealTimeStarted = true

	// Messages channel
	WebsocketRealTimeChannel = make(chan models.WebsocketRealTime)

	for {
		// Real-time percents
		conn.WriteJSON(<-WebsocketRealTimeChannel)

		// Avoid overload by pausing the goroutine for 0.1 seconds
		time.Sleep(time.Duration(1e+9))
	}
}

func WarningWebsocket(c *gin.Context) {
	// Upgrades the HTTP server connection to the WebSocket protocol
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade failed: ", err)
		return
	}
	defer conn.Close()

	// Informs that the websocket has been activated
	WebsocketForWarningStarted = true

	// Messages channel
	WebsocketForWarningChannel = make(chan models.WebsocketForWarning)

	for {
		// Overload warnings
		conn.WriteJSON(<-WebsocketForWarningChannel)

		// Avoid overload by pausing the goroutine for 0.1 seconds
		time.Sleep(time.Duration(1e+9))
	}
}
