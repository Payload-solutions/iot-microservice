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
	db.AutoMigrate(&models.SoilValues{}, &models.RealEnvironmentValues{})
	return &StructRepo{
		Db: db,
	}
}

// get the last twentty values of the table
func (repo *StructRepo) ReadTheLast(c *gin.Context) {
	var reader []models.RealEnvironmentValues
	err := models.ReadTheLastTwenttyValues(repo.Db, &reader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{
				"error": err,
			})
		return
	}

	c.JSON(http.StatusOK, reader[len(reader)-10:])
}

// get the last twentty values of the table
func (repo *StructRepo) ReadTheLastOne(c *gin.Context) {
	var reader []models.RealEnvironmentValues
	err := models.ReadTheLastTwenttyValues(repo.Db, &reader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{
				"error": err,
			})
		return
	}
	c.JSON(http.StatusOK, reader[len(reader)-1:])
}

func (repo *StructRepo) GetEnvironValues(c *gin.Context) {
	var reader []models.RealEnvironmentValues
	err := models.ReadEnviron(repo.Db, &reader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{
				"error": err,
			})
		return
	}
	c.JSON(http.StatusOK, reader)
}

func (repo *StructRepo) PostEnvironValues(c *gin.Context) {
	var reader models.RealEnvironmentValues
	c.ShouldBindJSON(&reader)

	err := models.CreateEnviron(repo.Db, &reader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error": err,
			})
		return
	}
	c.JSON(http.StatusOK, reader)
}

func ControllerSample(c *gin.Context) {
	c.JSON(200, JsonResponseStruct{
		StatusCode: 200,
		Body:       "sample",
	})
}

func (repo *StructRepo) PhEnvironValues(c *gin.Context) {
	var reader models.SoilValues
	c.ShouldBindJSON(&reader)

	err := models.CreateReading(repo.Db, &reader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error": err,
			})
		return
	}
	c.JSON(http.StatusOK, reader)
}
func (repo *StructRepo) PhSoilValues(c *gin.Context) {
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
