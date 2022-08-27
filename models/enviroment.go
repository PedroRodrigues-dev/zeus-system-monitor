package models

type Enviroment struct {
	CpuPercentLimit       int
	CpuOverloadMessage    string
	MemoryPercentLimit    int
	MemoryOverloadMessage string
	DiskPercentLimit      int
	DiskOverloadMessage   string
}
