package server

import (
	"github.com/Payload-solutions/iot-microservice/controllers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) *gin.Engine {
	repoLayer := controllers.New()
	router.Group("/").
		GET("iot-values", repoLayer.PhSoilValues).
		POST("post-phvalues", repoLayer.PhEnvironValues).
		POST("post-environ", repoLayer.PostEnvironValues).
		GET("get-environ", repoLayer.GetEnvironValues).
		GET("last-environ", repoLayer.ReadTheLast).
		GET("last-one", repoLayer.ReadTheLastOne).
		GET("last-one-ph", repoLayer.ReadTheLastOnePh)
	return router
}
