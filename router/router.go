package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"uts-golang-laundry/laundry/handlers"
	"uts-golang-laundry/laundry/repository"
	"uts-golang-laundry/laundry/usecase"
)

type HandlerEndPoint struct {
	DB *gorm.DB
	R  *gin.Engine
}

func (H HandlerEndPoint) Routers() {
	Repo := repository.NewRepo(H.DB)
	Usecase := usecase.NewUseCase(Repo)

	Repo2 := repository.NewRepo2(H.DB)
	Usecase2 := usecase.NewUseCase2(Repo2)

	Repo3 := repository.NewRepo3(H.DB)
	Usecase3 := usecase.NewUseCase3(Repo3)

	Repo4 := repository.NewRepo4(H.DB)
	Usecase4 := usecase.NewUseCase4(Repo4)

	Repo5 := repository.NewRepo5(H.DB)
	Usecase5 := usecase.NewUseCase5(Repo5)

	Repo6 := repository.NewRepo6(H.DB)
	Usecase6 := usecase.NewUseCase6(Repo6)

	v1 := H.R.Group("Api")
	handlers.OutletRouters(Usecase, v1)
	handlers.CustomerRouters(Usecase2, v1)
	handlers.TypeRouters(Usecase3, v1)
	handlers.UserRouters(Usecase4, v1)
	handlers.PriceRouters(Usecase5, v1)
	handlers.TransactionRouters(Usecase6, v1)

}
