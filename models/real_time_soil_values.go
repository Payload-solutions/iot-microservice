package models

import (
	"time"

	"gorm.io/gorm"
)

type RealEnvironmentValues struct {
	SoilId             int64     `gorm:"primary_key;auto_increment;not_null" json:"soil_id"`
	SoilDates          time.Time `gorm:"type:timestamp;default:current_timestamp" json:"soil_dates"`
	TemperatureSoil    float64   `json:"temperature_soil"`
	TemperatureEnviron float64   `json:"temperature_environ"`
	PH                 float64   `json:"ph"`
	Humidity           float64   `json:"humidity"`
	SoilState          int       `gorm:"type:int;default:1" json:"soil_state"`
}

func CreateEnviron(db *gorm.DB, reader *RealEnvironmentValues) (err error) {
	err = db.Create(reader).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadEnviron(db *gorm.DB, reader *[]RealEnvironmentValues) (err error) {
	err = db.Find(reader).Error
	if err != nil {
		return err
	}
	return nil
}
