package router

import (
	"hacktiv8-workshop-go-concurrency/controllers"
	race_example "hacktiv8-workshop-go-concurrency/controllers/race-example"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/update-stock-race", race_example.UpdateStock)
	router.GET("/update-stock", controllers.UpdateStock)

	return router
}
