package repository

import (
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(Db *gorm.DB) *Repo {
	return &Repo{
		DB: Db,
	}
}

func (rp Repo) GetAll(ctx *gin.Context) ([]models.Outlet, error) {
	var outlets []models.Outlet
	rp.DB.Find(&outlets)

	return outlets, nil
}

func (rp Repo) Create(data models.Outlet) error {
	err := rp.DB.Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}

func (rp Repo) Update(data models.Outlet, param string) error {
	err := rp.DB.First(&models.Outlet{}, "id_outlet=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id_outlet = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repo) Delete(param string) error {

	err := rp.DB.First(&models.Outlet{}, "id_outlet=?", param).Error
	if err != nil {
		return err
	}

	err = rp.DB.Delete(&models.Outlet{}, rp.DB.Where("id_outlet = ?", param)).Error
	if err != nil {
		return err
	}

	return nil
}
