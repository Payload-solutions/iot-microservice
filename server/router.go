package server

import (
	. "github.com/Payload-solutions/iot-microservice/controllers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) *gin.Engine {
	router.Group("/").
		GET("iot-values", ControllerSample)
	return router
}
