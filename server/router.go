package server

import (
	"github.com/Payload-solutions/iot-microservice/controllers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) *gin.Engine {
	repoLayer := controllers.New()
	router.Group("/").
		GET("iot-values", repoLayer.GetSoilValues).
		POST("post-environ", repoLayer.PostEnvironValues).
		GET("get-environ", repoLayer.GetEnvironValues).
		GET("last-environ", repoLayer.ReadTheLast).
		GET("last-one", repoLayer.ReadTheLastOne)
	return router
}
