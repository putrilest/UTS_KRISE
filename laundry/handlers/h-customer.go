package handlers

import (
	"net/http"
	"uts-golang-laundry/laundry"

	"github.com/gin-gonic/gin"
)

func CustomerRouters(cst laundry.Usecase2, r *gin.RouterGroup) {
	uc := Usecasecst{
		cst,
	}

	v2 := r.Group("customer")
	v2.GET("", uc.GetCustomer)
	v2.POST("", uc.CreateCustomer)
	v2.PUT(":id_customer", uc.PutCustomer)
	v2.DELETE(":id_customer", uc.DeleteCustomer)

}

type Usecasecst struct {
	CustomerUsecase laundry.Usecase2
}

func (cst Usecasecst) GetCustomer(ctx *gin.Context) {
	result, err := cst.CustomerUsecase.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (cst Usecasecst) CreateCustomer(ctx *gin.Context) {
	err := cst.CustomerUsecase.Create(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (cst Usecasecst) PutCustomer(ctx *gin.Context) {
	err := cst.CustomerUsecase.Update(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (cst Usecasecst) DeleteCustomer(ctx *gin.Context) {
	err := cst.CustomerUsecase.Delete(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
