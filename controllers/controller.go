package controllers

import "github.com/gin-gonic/gin"

type JsonResponseStruct struct {
	StatusCode int64  `json:"status_code"`
	Body       string `json:"body"`
	Success    bool   `json:"success"`
}

func ControllerSample(c *gin.Context) {
	c.JSON(200, JsonResponseStruct{
		StatusCode: 200,
		Body:       "sample",
	})
}
