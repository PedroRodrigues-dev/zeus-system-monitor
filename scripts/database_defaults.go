package scripts

import (
	"zeus/configs"
	"zeus/models"
)

func DatabaseDefaults() {
	// Default values
	cfg := []models.Config{
		{Name: "CpuPercentLimit", Value: "90"},
		{Name: "CpuOverloadMessage", Value: "CPU is in overload"},
		{Name: "MemoryPercentLimit", Value: "90"},
		{Name: "MemoryOverloadMessage", Value: "Memory is in overload"},
		{Name: "DiskPercentLimit", Value: "90"},
		{Name: "DiskOverloadMessage", Value: "Disk is in overload"},
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
