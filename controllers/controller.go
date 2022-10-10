package controllers

import (
	"dtskominfo-gin/apimodels"
	"dtskominfo-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var req apimodels.Request
	var res apimodels.Response
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := services.SaveOrder(req)
	if err != nil {
		res.Status = "Create Order Gagal"
		res.ResponseCode = "400"
	}
	ctx.JSON(http.StatusOK, res)
	return
}

func GetOrderBy(orderID string) {
	//orderId := ctx.Param("orderID")
	//log.Println("OrderID :", orderId)
	//res, _ := services.GetOrderBy(orderId)
	//ctx.JSON(http.StatusOK, res)

}

func GetOrderAll(ctx *gin.Context) {

}
func UpdateOrder(ctx *gin.Context) {

}
func DeleteOrder(ctx *gin.Context) {

}
