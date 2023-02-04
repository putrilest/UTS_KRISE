package handlers

import (
	"net/http"
	"uts-golang-laundry/laundry"

	"github.com/gin-gonic/gin"
)

func OutletRouters(otl laundry.Usecase, r *gin.RouterGroup) {
	uc := Usecaseotl{
		otl,
	}

	v2 := r.Group("outlet")
	v2.GET("", uc.Getoutlet)
	v2.POST("", uc.Createoutlet)
	v2.PUT(":id_outlet", uc.Putoutlet)
	v2.DELETE(":id_outlet", uc.Deleteoutlet)

}

type Usecaseotl struct {
	outletUsecase laundry.Usecase
}

func (otl Usecaseotl) Getoutlet(ctx *gin.Context) {
	result, err := otl.outletUsecase.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (otl Usecaseotl) Createoutlet(ctx *gin.Context) {
	err := otl.outletUsecase.Create(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (otl Usecaseotl) Putoutlet(ctx *gin.Context) {
	err := otl.outletUsecase.Update(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (otl Usecaseotl) Deleteoutlet(ctx *gin.Context) {
	err := otl.outletUsecase.Delete(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
