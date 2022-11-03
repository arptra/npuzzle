package apiserver

import (
	handler "N-puzzle-GO/api"
	"github.com/gin-gonic/gin"
)

var firstState map[string][]int

func ApiServerStart() {
	router := gin.Default()
	handler.RequestHandler(router)
	router.Run()
}
