package laundry

import (
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(ctx *gin.Context) ([]models.Outlet, error)
	Create(data models.Outlet) error
	Update(data models.Outlet, param string) error
	Delete(Param string) error
}

type Usecase interface {
	GetAll(ctx *gin.Context) ([]models.Outlet, error)
	Create(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}
