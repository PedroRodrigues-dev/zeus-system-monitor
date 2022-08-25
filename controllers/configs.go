package controllers

import (
	"net/http"
	"zeus/configs"
	"zeus/models"
	"zeus/scripts"

	"github.com/gin-gonic/gin"
)

// FindAllConfigs		godoc
// @Summary      		Get configs array
// @Description  		Responds with the list of all configs as JSON.
// @Tags         		configs
// @Produce      		json
// @Success      		200  {array}  models.Config
// @Router       		/configs [get]
func FindAllConfigs(c *gin.Context) {
	var cfgs []models.Config
	configs.DB.Find(&cfgs)
	c.JSON(http.StatusOK, cfgs)
}

// FindOneConfig		godoc
// @Summary      		Get one config
// @Description  		Responds with one config as JSON.
// @Tags         		configs
// @Produce      		json
// @Param        		name  path      string  true  "search config by name"
// @Success      		200  {array}  models.Config
// @Router       		/configs/{name} [get]
func FindOneConfig(c *gin.Context) {
	var cfg models.Config

	if err := configs.DB.Where("name = ?", c.Param("name")).First(&cfg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, cfg)
}

// UpdateConfig			godoc
// @Summary      		Update a config
// @Description  		Responds with http status
// @Tags         		configs
// @Produce      		json
// @Param        		name  path      string  true  "update config by name"
// @Param       	 	config  body      models.Config  true  "Config JSON"
// @Success      		200
// @Router       		/configs/{name} [put]
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
