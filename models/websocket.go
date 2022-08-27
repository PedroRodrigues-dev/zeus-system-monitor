package models

type MonitorMessage struct {
	MonitorName string
	Percent     int
	IsOverload  bool
	Alert       string
}
