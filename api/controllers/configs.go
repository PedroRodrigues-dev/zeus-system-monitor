package controllers

import (
	"net/http"
	"zeus/configs"
	"zeus/models"
	"zeus/scripts"

	"github.com/gin-gonic/gin"
)

func FindAllConfigs(c *gin.Context) {
	var cfgs []models.Config
	configs.DB.Find(&cfgs)
	c.JSON(http.StatusOK, cfgs)
}

func FindOneConfig(c *gin.Context) {
	var cfg models.Config

	if err := configs.DB.Where("name = ?", c.Param("name")).First(&cfg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, cfg)
}

func UpdateConfig(c *gin.Context) {
	// Get config if exist
	var cfg models.Config
	if err := configs.DB.Where("name = ?", c.Param("name")).First(&cfg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Config
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Name = c.Param("name")
	configs.DB.Model(&cfg).Updates(input)

	scripts.DefineEnviroment()
	c.JSON(http.StatusOK, cfg)
}
