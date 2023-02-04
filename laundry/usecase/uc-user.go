package usecase

import (
	"log"
	"net/http"
	"uts-golang-laundry/laundry"
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
)

func NewUseCase4(usr laundry.Repository4) *Usecase4 {
	return &Usecase4{usr}
}

type Usecase4 struct {
	Repousr laundry.Repository4
}

func (us Usecase4) GetAll(ctx *gin.Context) ([]models.User, error) {
	result, err := us.Repousr.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (us Usecase4) Create(ctx *gin.Context) error {
	var data2 models.User  
	err := ctx.ShouldBindJSON(&data2)

	if err != nil {
		return err
	}

	err = us.Repousr.Create(data2)

	if err != nil {
		return err
	}

	return nil
}

func (us Usecase4) Update(ctx *gin.Context) error {
	var data2 models.User
	id := ctx.Param("id_user")

	if id == "" {
		log.Fatal("Id not found")
	}
	ctx.ShouldBindJSON(&data2)
	us.Repousr.Update(data2, id)
	return nil
}

func (us Usecase4) Delete(ctx *gin.Context) error {
	id := ctx.Param("id_user")

	if id == "" {
		log.Fatal("Id not found")
	}

	err := us.Repousr.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
