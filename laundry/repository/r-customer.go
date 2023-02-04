package repository

import (
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repo2 struct {
	DB *gorm.DB
}

func NewRepo2(Db *gorm.DB) *Repo2 {
	return &Repo2{
		DB: Db,
	}
}

func (rp Repo2) GetAll(ctx *gin.Context) ([]models.Customer, error) {
	var csts []models.Customer
	rp.DB.Find(&csts)

	return csts, nil
}

func (rp Repo2) Create(data models.Customer) error {
	err := rp.DB.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repo2) Update(data models.Customer, param string) error {
	err := rp.DB.First(&models.Customer{}, "id_customer=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id_customer = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repo2) Delete(param string) error {

	err := rp.DB.First(&models.Customer{}, "id_customer=?", param).Error
	if err != nil {
		return err
	}

	err = rp.DB.Delete(&models.Customer{}, rp.DB.Where("id_customer = ?", param)).Error
	if err != nil {
		return err
	}

	return nil

}
