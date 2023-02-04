package laundry

import (
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
)


type Repository5 interface {
	GetAll(ctx *gin.Context) ([]models.LaundryPrices, error)
	Create(data5 models.LaundryPrices) error
	Update(data5 models.LaundryPrices, param string) error
	Delete(Param string) error
}

type Usecase5 interface {
	GetAll(ctx *gin.Context) ([]models.LaundryPrices, error)
	Create(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}