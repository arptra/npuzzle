package apiserver

import (
	handler "N-puzzle-GO/api"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ApiServerStart() {
	router := gin.Default()
	handler.RequestHandler(router)
	if err := router.Run(); err != nil {
		fmt.Errorf("%s", err)
	}
}
