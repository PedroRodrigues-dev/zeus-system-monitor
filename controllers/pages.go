package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MonitorsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "monitors.html", gin.H{})
}
