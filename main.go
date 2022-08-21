package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"zeus/api"
	"zeus/configs"
	"zeus/core"
	"zeus/scripts"
)

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

	// Configure Rabbit
	configs.StartRabbit(scripts.ENV.RabbitUser, scripts.ENV.RabbitPassword, scripts.ENV.RabbitHost, scripts.ENV.RabbitPort)

	// Start monitors
	infinit := make(chan bool)
	for i := 0; i < 3; i++ {
		go core.Monitors(i)
	}

	// System up and down notifications
	log.Println("Initialized")
	fmt.Println("[Initialized]")
	notificateInterrupt()

	// Start API
	api.Server()

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
