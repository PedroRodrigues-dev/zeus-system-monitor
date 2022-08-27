package main

import (
	"embed"
	"fmt"
	"html/template"
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

//go:embed pages/*
var embeddedFiles embed.FS

// @title     		Zeus System Monitor API
// @version         v3.0.0
// @description     Real-time server resource monitor

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

	// Websocket
	r.GET("/monitor", controllers.MonitorWebsocket)

	// Api
	r.GET("/configs", controllers.FindAllConfigs)
	r.GET("/configs/:name", controllers.FindOneConfig)
	r.PUT("/configs/:name", controllers.UpdateConfig)

	// Pages
	templ := template.Must(template.New("").ParseFS(embeddedFiles, "pages/*"))
	r.SetHTMLTemplate(templ)
	r.GET("/", controllers.MonitorsPage)

	// Sets the port by args or default
	port := ":8080"
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}

	// System up and down notifications
	log.Println("Initialized")
	notificateInterrupt()

	// Try to start the server and warn in case of error
	err := r.Run(port)
	if err != nil {
		log.Printf("%s\n", err)
		log.Println("Finalized")
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
		os.Exit(0)
	}()
}
