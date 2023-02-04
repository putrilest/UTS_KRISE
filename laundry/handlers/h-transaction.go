package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"uts-golang-laundry/laundry"
)

func TransactionRouters(tran laundry.Usecase6, r *gin.RouterGroup) {
	uc := Usecasetransaction {
		tran, 
	}

	v2:= r.Group("transaction")
	v2.GET("",uc.Gettransaction)
	v2.POST("",uc.Createtransaction)
	v2.PUT(":id_transaction",uc.Puttransaction)
	v2.DELETE(":id_transaction",uc.Deletetransaction)

}

type Usecasetransaction struct {
	TransUsecase laundry.Usecase6
}

func (tran Usecasetransaction) Gettransaction(ctx *gin.Context) {
	result, err := tran.TransUsecase.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (tran Usecasetransaction) Createtransaction(ctx *gin.Context) {
	err := tran.TransUsecase.Create(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (tran Usecasetransaction) Puttransaction(ctx *gin.Context) {
	err := tran.TransUsecase.Update(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (tran Usecasetransaction) Deletetransaction(ctx *gin.Context) {
	err := tran.TransUsecase.Delete(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
