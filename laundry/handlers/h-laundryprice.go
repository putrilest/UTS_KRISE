package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"uts-golang-laundry/laundry"
)

func PriceRouters(prc laundry.Usecase5, r *gin.RouterGroup) {
	uc := Usecaseprices {
		prc, 
	}

	v2:= r.Group("price")
	v2.GET("",uc.Getprice)
	v2.POST("",uc.Createprice)
	v2.PUT(":id_price",uc.Putprice)
	v2.DELETE(":id_price",uc.Deleteprice)

}

type Usecaseprices struct {
	PriUsecase laundry.Usecase5
}

func (prc Usecaseprices) Getprice(ctx *gin.Context) {
	result, err := prc.PriUsecase.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (prc Usecaseprices) Createprice(ctx *gin.Context) {
	err := prc.PriUsecase.Create(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (prc Usecaseprices) Putprice(ctx *gin.Context) {
	err := prc.PriUsecase.Update(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (prc Usecaseprices) Deleteprice(ctx *gin.Context) {
	err := prc.PriUsecase.Delete(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
