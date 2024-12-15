package main

import (
	"go-simple-MVC/config"
	"go-simple-MVC/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//inisialiasai Gin
	db := config.DBInit()

	router := gin.Default()

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("bookabledate", walletName)
	// }

	routes.Routes(router, db)

	//membuat route
	// router.GET("/persons", inDB.GetAllPerson)
	// router.GET("/person/:id", inDB.GetDetailPerson)
	// router.POST("/person", inDB.CreatePerson)
	// router.PUT("/person", inDB.UpdatePerson)
	// router.DELETE("/person/:id", inDB.DeletePerson)
	

	//mulai server dengan port 3000
	router.Run(":3000")
}