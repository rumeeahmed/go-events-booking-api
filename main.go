package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-events-booking-api/db"
	"go-events-booking-api/routes"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
