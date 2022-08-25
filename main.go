package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"zeus/configs"
	"zeus/controllers"
	"zeus/core"
	"zeus/scripts"

	_ "zeus/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title     		Zeus System Monitor API
// @version         v2.0.0
// @description     Real-time server resource monitor and overload notification

// @contact.name   	Pedro Rodrigues
// @contact.email  	pedroras004@gmail.com

// @license.name  	MIT
// @license.url   	https://mit-license.org
func main() {
	// System logo
	fmt.Println(` ------------------------------------------	
| |        \|        \|  \  |  \ /      \  |
|  \$$$$$$$$| $$$$$$$$| $$  | $$|  $$$$$$\ |
|     /  $$ | $$__    | $$  | $$| $$___\$$ |
|    /  $$  | $$  \   | $$  | $$ \$$    \  |
|   /  $$   | $$$$$   | $$  | $$ _\$$$$$$\ |
|  /  $$___ | $$_____ | $$__/ $$|  \__| $$ |
| |  $$    \| $$     \ \$$    $$ \$$    $$ | 
|  \$$$$$$$$ \$$$$$$$$  \$$$$$$   \$$$$$$  |	
 ------------------------------------------`)

	// Define system logs
	file, _ := os.OpenFile("system.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)

	// Configure database
	configs.ConnectDatabase()
	scripts.DatabaseDefaults()
	scripts.DefineEnviroment()

	// Declare webserver
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Start monitors
	infinit := make(chan bool)
	for i := 0; i < 3; i++ {
		go core.Monitors(i)
	}

	// Swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Websocekt
	r.GET("/realtime", controllers.RealTimeWebsocket)
	r.GET("/warning", controllers.WarningWebsocket)

	// Api
	r.GET("/configs", controllers.FindAllConfigs)
	r.GET("/configs/:name", controllers.FindOneConfig)
	r.PUT("/configs/:name", controllers.UpdateConfig)

	// Sets the port by args or default
	port := ":8080"
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}

	// System up and down notifications
	log.Println("Initialized")
	fmt.Println("[Initialized]")
	notificateInterrupt()

	// Try to start the server and warn in case of error
	err := r.Run(port)
	if err != nil {
		log.Printf("%s\n", err)
		fmt.Printf("[%s]\n", err)
		log.Println("Finalized")
		fmt.Println("[Finalized]")
		os.Exit(0)
	}
	<-infinit
}

// Notifies if the application is finished
func notificateInterrupt() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Finalized")
		fmt.Println("\n[Finalized]")
		os.Exit(0)
	}()
}
