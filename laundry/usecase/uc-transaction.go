package usecase

import (
	"log"
	"net/http"
	"uts-golang-laundry/laundry"
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
)

func NewUseCase6(trs laundry.Repository6) *Usecase6 {
	return &Usecase6{trs}
}

type Usecase6 struct {
	Repotrs laundry.Repository6
}

func (us Usecase6) GetAll(ctx *gin.Context) ([]models.Transaction, error) {
	result, err := us.Repotrs.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (us Usecase6) Create(ctx *gin.Context) error {
	var data2 models.Transaction
	err := ctx.ShouldBindJSON(&data2)

	if err != nil {
		return err
	}

	err = us.Repotrs.Create(data2)

	if err != nil {
		return err
	}

	return nil
}

func (us Usecase6) Update(ctx *gin.Context) error {
	var data2 models.Transaction
	id := ctx.Param("id_transaction")

	if id == "" {
		log.Fatal("Id not found")
	}
	ctx.ShouldBindJSON(&data2)
	us.Repotrs.Update(data2, id)
	return nil
}

func (us Usecase6) Delete(ctx *gin.Context) error {
	id := ctx.Param("id_transaction")

	if id == "" {
		log.Fatal("Id not found")
	}

	err := us.Repotrs.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
