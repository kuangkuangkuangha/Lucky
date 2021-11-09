package main

import (
	"github.com/gin-gonic/gin"
	"lucky/db_server"
	"lucky/server"
)

var httpServer *gin.Engine

func main() {
	defer func() {
		db_server.MySqlDb.Close()
	}()

	server.Run(httpServer)

}
