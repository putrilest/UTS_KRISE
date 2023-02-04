package usecase

import (
	"log"
	"net/http"
	"uts-golang-laundry/laundry"
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
)

func NewUseCase(otl laundry.Repository) *Usecase {
	return &Usecase{otl}
}

type Usecase struct {
	RepoOtl laundry.Repository
}

func (us Usecase) GetAll(ctx *gin.Context) ([]models.Outlet, error) {
	result, err := us.RepoOtl.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (us Usecase) Create(ctx *gin.Context) error {
	var data models.Outlet
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = us.RepoOtl.Create(data)

	if err != nil {
		return err
	}

	return nil
}

func (us Usecase) Update(ctx *gin.Context) error {
	var data models.Outlet
	id := ctx.Param("id_outlet")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}
	ctx.ShouldBindJSON(&data)
	us.RepoOtl.Update(data, id)
	return nil
}

func (us Usecase) Delete(ctx *gin.Context) error {
	id := ctx.Param("id_outlet")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := us.RepoOtl.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
