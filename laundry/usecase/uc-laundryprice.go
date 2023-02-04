package usecase

import (
	"log"
	"net/http"
	"uts-golang-laundry/laundry"
	"uts-golang-laundry/models"

	"github.com/gin-gonic/gin"
)

func NewUseCase5(prc laundry.Repository5) *Usecase5 {
	return &Usecase5{prc}
}

type Usecase5 struct {
	Repoprc laundry.Repository5
}

func (us Usecase5) GetAll(ctx *gin.Context) ([]models.LaundryPrices, error) {
	result, err := us.Repoprc.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (us Usecase5) Create(ctx *gin.Context) error {
	var data2 models.LaundryPrices
	err := ctx.ShouldBindJSON(&data2)

	if err != nil {
		return err
	}

	err = us.Repoprc.Create(data2)

	if err != nil {
		return err
	}

	return nil
}

func (us Usecase5) Update(ctx *gin.Context) error {
	var data2 models.LaundryPrices
	id := ctx.Param("id_price")

	if id == "" {
		log.Fatal("Id Not Found")
	}
	ctx.ShouldBindJSON(&data2)
	us.Repoprc.Update(data2, id)
	return nil
}

func (us Usecase5) Delete(ctx *gin.Context) error {
	id := ctx.Param("id_price")

	if id == "" {
		log.Fatal("Id Not Found")
	}

	err := us.Repoprc.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
