package apiserver

import (
	//handler "N-puzzle-GO/api"
	"github.com/gin-gonic/gin"
)

func ApiServerStart(config *configs.Config) {
	router := gin.Default()
	//handler.RequestHandler(router)
	router.Run()
}
