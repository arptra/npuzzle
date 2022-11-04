package api

import (
	"github.com/gin-gonic/gin"
)

func RequestHandler(router *gin.Engine) {
	router.PUT("/api/v1/state/", PutState)
	router.GET("/api/v1/path/", GetPath)
	//router.GET("/api/v1/data/receive/user/:id", GetUser)
	//router.DELETE("/api/v1/data/delete/users/:id", DeleteUser)
}
