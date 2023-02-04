package repository

import (
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repo6 struct {
	DB *gorm.DB
}

func NewRepo6(Db *gorm.DB) *Repo6 {
	return &Repo6{
		DB: Db,
	}
}

func (rp Repo6) GetAll(ctx *gin.Context) ([]models.Transaction, error) {
	var trans []models.Transaction
	rp.DB.Find(&trans)

	return trans, nil
}

func (rp Repo6) Create(data models.Transaction) error {
	err := rp.DB.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repo6) Update(data models.Transaction, param string) error {
	err := rp.DB.First(&models.Transaction{}, "id_transaction=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id_transaction = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repo6) Delete(param string) error {

	err := rp.DB.First(&models.Transaction{}, "id_transaction=?", param).Error
	if err != nil {
		return err
	}

	err = rp.DB.Delete(&models.Transaction{}, rp.DB.Where("id_transaction = ?", param)).Error
	if err != nil {
		return err
	}

	return nil
}
