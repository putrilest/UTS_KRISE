package laundry

import (
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
)

type Repository4 interface {
	GetAll(ctx *gin.Context) ([]models.User, error)
	Create(data4 models.User) error
	Update(data4 models.User, param string) error
	Delete(Param string) error
}

type Usecase4 interface {
	GetAll(ctx *gin.Context) ([]models.User, error)
	Create(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}
