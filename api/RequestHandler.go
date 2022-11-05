package api

import (
	"github.com/gin-gonic/gin"
)

func RequestHandler(router *gin.Engine) {
	router.PUT("/api/v1/state", PutState)
	router.GET("/api/v1/algo", GetAlgo)
	router.GET("/api/v1/path", GetPath)
	router.GET("/api/v1/stop", StopCalc)
}
