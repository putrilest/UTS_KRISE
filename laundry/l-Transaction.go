package laundry

import (
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
)

type Repository6 interface {
	GetAll(ctx *gin.Context) ([]models.Transaction, error)
	Create(data6 models.Transaction) error
	Update(data6 models.Transaction, param string) error
	Delete(Param string) error
}

type Usecase6 interface {
	GetAll(ctx *gin.Context) ([]models.Transaction, error)
	Create(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}
