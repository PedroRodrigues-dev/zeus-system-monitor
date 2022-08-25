package models

type WebsocketRealTime struct {
	MonitorName string
	Percent     int
}

type WebsocketForWarning struct {
	MonitorName    string
	WarningMessage string
}
