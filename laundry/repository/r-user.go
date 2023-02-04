package repository

import (
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repo4 struct {
	DB *gorm.DB
}

func NewRepo4(Db *gorm.DB) *Repo4 {
	return &Repo4{
		DB: Db,
	}
}

func (rp Repo4) GetAll(ctx *gin.Context) ([]models.User, error) {
	var users []models.User
	rp.DB.Find(&users)

	return users, nil
}

func (rp Repo4) Create(data models.User) error {
	err := rp.DB.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repo4) Update(data models.User, param string) error {
	err := rp.DB.First(&models.User{}, "id_user=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id_user = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repo4) Delete(param string) error {

	err := rp.DB.First(&models.User{}, "id_user=?", param).Error
	if err != nil {
		return err
	}

	err = rp.DB.Delete(&models.User{}, rp.DB.Where("id_user = ?", param)).Error
	if err != nil {
		return err
	}

	return nil
}
