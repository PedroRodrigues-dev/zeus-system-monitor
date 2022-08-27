package scripts

import (
	"strconv"
	"zeus/configs"
	"zeus/models"
)

// Enviroment
var ENV *models.Enviroment

func DefineEnviroment() {
	// Find values in database
	config := []models.Config{}
	configs.DB.Find(&config)

	// Define Enviroment variables
	env := models.Enviroment{}
	for i := range config {
		switch {
		case config[i].Name == "CpuPercentLimit":
			env.CpuPercentLimit, _ = strconv.Atoi(config[i].Value)
		case config[i].Name == "CpuOverloadMessage":
			env.CpuOverloadMessage = config[i].Value
		case config[i].Name == "MemoryPercentLimit":
			env.MemoryPercentLimit, _ = strconv.Atoi(config[i].Value)
		case config[i].Name == "MemoryOverloadMessage":
			env.MemoryOverloadMessage = config[i].Value
		case config[i].Name == "DiskPercentLimit":
			env.DiskPercentLimit, _ = strconv.Atoi(config[i].Value)
		case config[i].Name == "DiskOverloadMessage":
			env.DiskOverloadMessage = config[i].Value
		}
	}

	ENV = &env
}
