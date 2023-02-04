package usecase

import (
	"log"
	"net/http"
	"uts-golang-laundry/laundry"
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
)

func NewUseCase2(cst laundry.Repository2) *Usecase2 {
	return &Usecase2{cst}
}

type Usecase2 struct {
	Repocst laundry.Repository2
}

func (us Usecase2) GetAll(ctx *gin.Context) ([]models.Customer, error) {
	result, err := us.Repocst.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (us Usecase2) Create(ctx *gin.Context) error {
	var data2 models.Customer
	err := ctx.ShouldBindJSON(&data2)

	if err != nil {
		return err
	}

	err = us.Repocst.Create(data2)

	if err != nil {
		return err
	}

	return nil
}

func (us Usecase2) Update(ctx *gin.Context) error {
	var data2 models.Customer
	id := ctx.Param("id_customer")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}
	ctx.ShouldBindJSON(&data2)
	us.Repocst.Update(data2, id)
	return nil
}

func (us Usecase2) Delete(ctx *gin.Context) error {
	id := ctx.Param("id_customer")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := us.Repocst.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
