package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pratyushvid3105/Go-Rest-API/db"
	"github.com/pratyushvid3105/Go-Rest-API/routes"
)

func main(){
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8083") // localhost:8083
}

