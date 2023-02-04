package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"uts-golang-laundry/laundry"
)

func TypeRouters(typ laundry.Usecase3, r *gin.RouterGroup) {
	uc := Usecasetyp {
		typ, 
	}

	v2:= r.Group("types")
	v2.GET("",uc.Gettype)
	v2.POST("",uc.Createtype)
	v2.PUT(":id_type",uc.Puttype)
	v2.DELETE(":id_type",uc.Deletetype)

}

type Usecasetyp struct {
	TypeUsecase laundry.Usecase3
}

func (typ Usecasetyp) Gettype(ctx *gin.Context) {
	result, err := typ.TypeUsecase.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (typ Usecasetyp) Createtype(ctx *gin.Context) {
	err := typ.TypeUsecase.Create(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (typ Usecasetyp) Puttype(ctx *gin.Context) {
	err := typ.TypeUsecase.Update(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()	})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (typ Usecasetyp) Deletetype(ctx *gin.Context) {
	err := typ.TypeUsecase.Delete(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
