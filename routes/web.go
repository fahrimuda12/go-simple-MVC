package routes

import (
	"go-simple-MVC/app/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(routes *gin.Engine, db *gorm.DB) {

	inDB := &controllers.InDB{DB: db}

	routes.GET("/wallet", inDB.GetAllWallet)
	routes.GET("/wallet/:id", inDB.GetDetailWallet)
	routes.POST("/wallet/create", inDB.CreateWallet)
	routes.PUT("/wallet/update", inDB.UpdateWallet)
	routes.DELETE("/wallet/:id/delete", inDB.DeleteWallet)
}