package core

import (
	"log"
	"zeus/controllers"
	"zeus/models"
)

// 0 - cpu | 1 - memory | 2 - disk
var lastRealTimeNotification = [3]models.WebsocketRealTime{}

func WarningLogging(msg string) {
	log.Println(msg)
}

func WebsocketForWarningNotificator(msg string, option int) {
	// Choose monitor
	monitorName := defineMonitorName(option)

	// Create json response for websocket
	var message models.WebsocketForWarning
	message.MonitorName = monitorName
	message.WarningMessage = msg

	// Send message to websocket channel
	controllers.WebsocketForWarningChannel <- message
}

func WebsocketRealTimeNotificator(percent int, option int) {
	// Choose monitor
	monitorName := defineMonitorName(option)

	// Create json response for websocket
	var message models.WebsocketRealTime
	message.MonitorName = monitorName
	message.Percent = percent

	// Checks if current message is different from the previous one
	if lastRealTimeNotification[option] != message {
		lastRealTimeNotification[option] = message
		// Send message to websocket channel
		controllers.WebsocketRealTimeChannel <- message
	}
}

func defineMonitorName(option int) string {
	switch {
	case option == 0:
		return "cpu"
	case option == 1:
		return "memory"
	case option == 2:
		return "disk"
	default:
		return ""
	}
}
