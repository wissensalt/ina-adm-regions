package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wissensalt/wissensalt-go-util/db"
	"github.com/wissensalt/wissensalt-go-util/logging"
)

func main()  {
	//init logger
	logging.Init()
	logging.AppLogger.Println("Hello World")

	//init db connection
	dbConnection := new(db.ConnectionProperty)
	dbConnection.Init()
	var connectionService db.ConnectionService = dbConnection
	connectionService.ConnectMySql()

	// init router
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
