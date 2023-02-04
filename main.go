package main

import (
	"uts-golang-laundry/connection"
	"uts-golang-laundry/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := connection.ConnectToDb()
	rh := router.HandlerEndPoint{
		DB: db,
		R:  r,
	}

	rh.Routers()
	r.Run(":8001")
}
