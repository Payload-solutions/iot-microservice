package models

import (
	"time"

	"gorm.io/gorm"
)

type SoilValues struct {
	SoilId    int64     `gorm:"primary_key;auto_increment;not_null" json:"soil_id"`
	PH        float64   `json:"ph"`
	SoilDates time.Time `gorm:"type:timestamp;default:current_timestamp" json:"soil_dates"`
	SoilState int       `gorm:"type:int;default:1" json:"soil_state"`
}

func CreateReading(db *gorm.DB, reading *SoilValues) (err error) {
	err = db.Create(reading).Error
	if err != nil {
		return err
	}
	return nil
}

func GetReading(db *gorm.DB, reading *[]SoilValues) (err error) {
	err = db.Find(reading).Error
	if err != nil {
		return err
	}
	return nil
}
