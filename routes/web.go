package routes

import (
	"go-simple-MVC/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(routes *gin.Engine, db *gorm.DB) {

	inDB := &controllers.InDB{DB: db}

	routes.GET("/persons", inDB.GetAllPerson)
	routes.GET("/person/:id", inDB.GetDetailPerson)
	routes.POST("/person", inDB.CreatePerson)
	routes.PUT("/person", inDB.UpdatePerson)
	routes.DELETE("/person/:id", inDB.DeletePerson)

	routes.GET("/wallet", inDB.GetAllWallet)
	routes.GET("/wallet/:id", inDB.GetDetailWallet)
	routes.POST("/wallet", inDB.CreateWallet)
	routes.PUT("/wallet", inDB.UpdateWallet)
	routes.DELETE("/wallet/:id", inDB.DeleteWallet)
}