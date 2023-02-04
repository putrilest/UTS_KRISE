package repository

import (
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repo3 struct {
	DB *gorm.DB
}

func NewRepo3(Db *gorm.DB) *Repo3 {
	return &Repo3{
		DB: Db,
	}
}

func (rp Repo3) GetAll(ctx *gin.Context) ([]models.LaundryType, error) {
	var types []models.LaundryType
	rp.DB.Find(&types)

	return types, nil
}

func (rp Repo3) Create(data models.LaundryType) error {
	err := rp.DB.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repo3) Update(data models.LaundryType, param string) error {
	err := rp.DB.First(&models.LaundryType{}, "id_type=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id_type = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repo3) Delete(param string) error {

	err := rp.DB.First(&models.LaundryType{}, "id_type=?", param).Error
	if err != nil {
		return err
	}

	err = rp.DB.Delete(&models.LaundryType{}, rp.DB.Where("id_type = ?", param)).Error
	if err != nil {
		return err
	}

	return nil
}
