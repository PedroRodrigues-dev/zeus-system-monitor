package core

import (
	"zeus/controllers"
	"zeus/models"
	"zeus/scripts"
)

func MonitiorWebsocketNotificator(percent int, option int) {
	// Create json response for websocket
	var message models.MonitorMessage
	message.Percent = percent
	message.MonitorName, message.Alert, message.IsOverload = defineMonitorValues(option, percent)
	if !message.IsOverload {
		message.Alert = ""
	}

	// Send message to websocket
	controllers.MonitorMessages[option] = message
}

func defineMonitorValues(option int, percent int) (string, string, bool) {
	switch {
	case option == 0:
		return "cpu", scripts.ENV.CpuOverloadMessage, percent >= scripts.ENV.CpuPercentLimit
	case option == 1:
		return "memory", scripts.ENV.MemoryOverloadMessage, percent >= scripts.ENV.MemoryPercentLimit
	case option == 2:
		return "disk", scripts.ENV.DiskOverloadMessage, percent >= scripts.ENV.DiskPercentLimit
	default:
		return "", "", false
	}
}
