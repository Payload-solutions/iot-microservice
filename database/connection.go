package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewInstance() *gorm.DB {
	return GetConnection()
}

func GetConnection() *gorm.DB {
	URL := fmt.Sprintf("root_arturo:arturo_root@tcp(127.0.0.1:3306)/IOTSoilValues?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(
		mysql.Open(URL),
		&gorm.Config{})
	if err != nil {
		log.Fatal("error in the connection")
		return nil
	}
	fmt.Println("Databse connected successfully")
	return db
}
