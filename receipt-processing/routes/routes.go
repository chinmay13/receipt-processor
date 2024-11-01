package routes

import (
	"receipt-processing/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/receipts/process", controllers.ProcessReceipt)
	router.GET("/receipts/:id/points", controllers.GetPoints)
	return router
}
