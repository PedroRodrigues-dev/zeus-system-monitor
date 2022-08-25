package scripts

import (
	"zeus/configs"
	"zeus/models"
)

func DatabaseDefaults() {
	// Default values
	cfg := []models.Config{
		{Name: "TimeNotificationLimit", Value: "300"},
		{Name: "CpuOverloadCounterLimit", Value: "30"},
		{Name: "CpuPercentLimit", Value: "90"},
		{Name: "CpuOverloadMessage", Value: "CPU with use above %d"},
		{Name: "MemoryOverloadCounterLimit", Value: "30"},
		{Name: "MemoryPercentLimit", Value: "90"},
		{Name: "MemoryOverloadMessage", Value: "Memory with use above %d"},
		{Name: "DiskOverloadCounterLimit", Value: "30"},
		{Name: "DiskPercentLimit", Value: "90"},
		{Name: "DiskOverloadMessage", Value: "Disk with use above %d"},
	}

	for i := range cfg {
		createOrIgnore(&models.Config{Name: cfg[i].Name, Value: cfg[i].Value})
	}
}

// Check field already exists in database
func createOrIgnore(config *models.Config) {
	cfg := models.Config{}
	configs.DB.Where("name = ?", config.Name).Find(&cfg)
	if cfg.Name == "" {
		configs.DB.Create(config)
	}
}
