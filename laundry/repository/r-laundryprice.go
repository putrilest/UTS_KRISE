package repository

import (
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repo5 struct {
	DB *gorm.DB
}

func NewRepo5(Db *gorm.DB) *Repo5 {
	return &Repo5{
		DB: Db,
	}
}

func (rp Repo5) GetAll(ctx *gin.Context) ([]models.LaundryPrices, error) {
	var prices []models.LaundryPrices
	rp.DB.Find(&prices)

	return prices, nil
}

func (rp Repo5) Create(data models.LaundryPrices) error {
	err := rp.DB.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repo5) Update(data models.LaundryPrices, param string) error {
	err := rp.DB.First(&models.LaundryPrices{}, "id_price=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id_price = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repo5) Delete(param string) error {

	err := rp.DB.First(&models.LaundryPrices{}, "id_price=?", param).Error
	if err != nil {
		return err
	}

	err = rp.DB.Delete(&models.LaundryPrices{}, rp.DB.Where("id_price = ?", param)).Error
	if err != nil {
		return err
	}

	return nil
}
