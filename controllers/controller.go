package controllers

import (
	"net/http"

	. "github.com/Payload-solutions/iot-microservice/database"
	"github.com/Payload-solutions/iot-microservice/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JsonResponseStruct struct {
	StatusCode int64  `json:"status_code"`
	Body       string `json:"body"`
	Success    bool   `json:"success"`
}

type StructRepo struct {
	Db *gorm.DB
}

func New() *StructRepo {
	db := NewInstance()
	db.AutoMigrate(&models.SoilValues{})
	return &StructRepo{
		Db: db,
	}
}

func (repo *StructRepo) GetSoilValues(c *gin.Context) {
	var reading []models.SoilValues
	err := models.GetReading(repo.Db, &reading)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{
				"error": err,
			})
		return
	}
	c.JSON(http.StatusOK, reading)
}

func ControllerSample(c *gin.Context) {
	c.JSON(200, JsonResponseStruct{
		StatusCode: 200,
		Body:       "sample",
	})
}
