package controllers

import (
	"dtskominfo-gin/apimodels"
	"dtskominfo-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {

	var req apimodels.Request
	user, password, ok := ctx.Request.BasicAuth()
	if ok && user == "iketutg" && password == "1234" {
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
	ctx.JSON(http.StatusNonAuthoritativeInfo, nil)
	// ctx.Abort()
	// ctx.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
}

func GetOrderAll(ctx *gin.Context) {
}
func UpdateOrder(ctx *gin.Context) {

}
func DeleteOrder(ctx *gin.Context) {

}
