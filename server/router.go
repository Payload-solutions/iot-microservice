package server

import (
	"github.com/Payload-solutions/iot-microservice/controllers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) *gin.Engine {
	repoLayer := controllers.New()
	router.Group("/").
		GET("iot-values", repoLayer.GetSoilValues)
	return router
}
