package usecase

import (
	"log"
	"net/http"
	"uts-golang-laundry/laundry"
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
)

func NewUseCase3(ltype laundry.Repository3) *Usecase3 {
	return &Usecase3{ltype}
}

type Usecase3 struct {
	Repotype laundry.Repository3
}

func (us Usecase3) GetAll(ctx *gin.Context) ([]models.LaundryType, error) {
	result, err := us.Repotype.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (us Usecase3) Create(ctx *gin.Context) error {
	var data2 models.LaundryType
	err := ctx.ShouldBindJSON(&data2)

	if err != nil {
		return err
	}

	err = us.Repotype.Create(data2)

	if err != nil {
		return err
	}

	return nil
}

func (us Usecase3) Update(ctx *gin.Context) error {
	var data2 models.LaundryType
	id := ctx.Param("id_type")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}
	ctx.ShouldBindJSON(&data2)
	us.Repotype.Update(data2, id)
	return nil
}

func (us Usecase3) Delete(ctx *gin.Context) error {
	id := ctx.Param("id_type")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := us.Repotype.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
