package router

import (
	"SimpleGolang/controllers"
	"github.com/gin-gonic/gin"
)

func MainRouter(route *gin.Engine) {
	// Simple Application Router
	route.GET("/", controllers.MainFile)
	route.GET("/GetAllProduct", controllers.GetAllProduct)
	route.GET("/GetProduct/:ProductID", controllers.GetProductID)
	route.POST("/MoreProduct", controllers.TambahProduct)
}
