package api

import (
	"fmt"
	"log"
	"os"
	"zeus/api/controllers"

	"github.com/gin-gonic/gin"
)

func Server() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/configs", controllers.FindAllConfigs)
	r.GET("/configs/:name", controllers.FindOneConfig)
	r.PUT("/configs/:name", controllers.UpdateConfig)

	// Sets the port by args or default
	port := ":8080"
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}

	// Try to start the server and warn in case of error
	err := r.Run(port)
	if err != nil {
		log.Printf("%s\n", err)
		fmt.Printf("[%s]\n", err)
		log.Println("Finalized")
		fmt.Println("[Finalized]")
		os.Exit(0)
	}
}
