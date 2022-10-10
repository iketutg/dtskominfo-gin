package routers

import (
	"dtskominfo-gin/controllers"
	"dtskominfo-gin/midlewares"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func setLogger() {
	file, _ := os.Create("dts.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}
func GetRouter() *gin.Engine {
	setLogger()
	router := gin.New()
	//engine.Use(Logger(), Recovery())
	router.Use(gin.Recovery(), midlewares.Logger(), midlewares.GenerateIdUnix())

	groupuser := router.Group("/api/v1/users")
	groupuser.POST("/registrasi", controllers.Registrasi)
	groupuser.POST("/login", controllers.Login)

	transgroup := router.Group("/api/v1/transaction/").Use(midlewares.JWTAuth())
	transgroup.POST("/orders", controllers.CreateOrder)
	transgroup.GET("/orders", controllers.GetOrderAll)
	transgroup.PUT("/orders/{orderID}", controllers.UpdateOrder)
	transgroup.DELETE("/orders/{orderID}", controllers.DeleteOrder)
	return router
}
