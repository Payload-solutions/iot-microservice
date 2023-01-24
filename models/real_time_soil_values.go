package models

import "gorm.io/gorm"

type RealEnvironmentValues struct {
	SoilId             int64   `json:"soil_id"`
	TemperatureSoil    float64 `json:"temperature_soil"`
	TemperatureEnviron float64 `json:"temperature_environ"`
	PH                 float64 `json:"ph"`
	Humidity           float64 `json:"humidity"`
	Dates              string  `json:"dates"`
	SoilState          int     `json:"soil_state"`
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
