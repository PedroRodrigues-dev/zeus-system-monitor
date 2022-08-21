package models

type Enviroment struct {
	TimeNotificationLimit      int
	CpuOverloadCounterLimit    int
	CpuPercentLimit            int
	CpuOverloadMessage         string
	MemoryOverloadCounterLimit int
	MemoryPercentLimit         int
	MemoryOverloadMessage      string
	DiskOverloadCounterLimit   int
	DiskPercentLimit           int
	DiskOverloadMessage        string
	RabbitHost                 string
	RabbitPort                 int
	RabbitUser                 string
	RabbitPassword             string
	RabbitQueue                string
}
