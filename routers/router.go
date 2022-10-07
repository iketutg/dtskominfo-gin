package routers

import (
	"dtskominfo-gin/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrderAll)
	router.PUT("/orders/{orderID}", controllers.UpdateOrder)
	router.DELETE("/orders/{orderID}", controllers.DeleteOrder)
	return router
}
