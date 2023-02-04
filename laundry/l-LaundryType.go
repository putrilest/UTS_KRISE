package laundry

import (
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
)

type Repository3 interface {
	GetAll(ctx *gin.Context) ([]models.LaundryType, error)
	Create(data3 models.LaundryType) error
	Update(data3 models.LaundryType, param string) error
	Delete(Param string) error
}

type Usecase3 interface {
	GetAll(ctx *gin.Context) ([]models.LaundryType, error)
	Create(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}