package handlers

import (
	"net/http"
	"uts-golang-laundry/laundry"

	"github.com/gin-gonic/gin"
)

func UserRouters(use laundry.Usecase4, r *gin.RouterGroup) {
	uc := Usecaseuser{
		use,
	}

	v2 := r.Group("user")
	v2.GET("", uc.Getuser)
	v2.POST("", uc.Createuser)
	v2.PUT(":id_user", uc.Putuser)
	v2.DELETE(":id_user", uc.Deleteuser)

}

type Usecaseuser struct {
	UserUsecase laundry.Usecase4
}

func (use Usecaseuser) Getuser(ctx *gin.Context) {
	result, err := use.UserUsecase.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (use Usecaseuser) Createuser(ctx *gin.Context) {
	err := use.UserUsecase.Create(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (use Usecaseuser) Putuser(ctx *gin.Context) {
	err := use.UserUsecase.Update(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (use Usecaseuser) Deleteuser(ctx *gin.Context) {
	err := use.UserUsecase.Delete(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
