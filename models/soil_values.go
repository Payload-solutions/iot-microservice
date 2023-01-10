package models

import "gorm.io/gorm"

type SoilValues struct {
	SoilId      int64   `json:"soil_id"`
	Temperature float64 `json:"temperature"`
	PH          float64 `json:"ph"`
	Dates       string  `json:"dates"`
	SoilState   int     `json:"soil_state"`
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
