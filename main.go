package main

import (
	config "beelogiq/notes/configs"
	"beelogiq/notes/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Nice ascii art
	welcome := `
███╗░░██╗░█████╗░████████╗███████╗░██████╗
████╗░██║██╔══██╗╚══██╔══╝██╔════╝██╔════╝
██╔██╗██║██║░░██║░░░██║░░░█████╗░░╚█████╗░
██║╚████║██║░░██║░░░██║░░░██╔══╝░░░╚═══██╗
██║░╚███║╚█████╔╝░░░██║░░░███████╗██████╔╝
╚═╝░░╚══╝░╚════╝░░░░╚═╝░░░╚══════╝╚═════╝░

	`

	fmt.Println(welcome)
	// Init Db
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":8080"))
}
