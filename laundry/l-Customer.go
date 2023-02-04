package laundry

import (
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
)

type Repository2 interface {
	GetAll(ctx *gin.Context) ([]models.Customer, error)
	Create(data2 models.Customer) error
	Update(data2 models.Customer, param string) error
	Delete(Param string) error

	
	
}

type Usecase2 interface {
	GetAll(ctx *gin.Context) ([]models.Customer, error)
	Create(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}
