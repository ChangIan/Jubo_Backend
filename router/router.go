package router

import (
	"github.com/gin-gonic/gin"

	middleware "changian.com/jubo/middleware"

	controllerV1 "changian.com/jubo/controller/v1"
)

func SetRouter() *gin.Engine {

	router := gin.Default()

	//
	router.Use(middleware.CorsMiddleware())

	// 路由组 - 版本管理 - V1
	groupV1 := router.Group("v1")
	{
		// 路由组 - Patient
		// 查詢 - Patient
		groupV1.GET("/patient", controllerV1.GetPatient)

		// 路由组 - Orders
		// 查詢 - Orders
		groupV1.GET("/orders/:id", controllerV1.GetOrders)
		// 新增 - Orders
		groupV1.POST("/orders", controllerV1.PostOrders)
		// 修改 - Orders
		groupV1.PUT("/orders", controllerV1.PutOrders)
		// 刪除 - Orders
		groupV1.DELETE("/orders/:id", controllerV1.DeleteOrders)
	}

	return router
}
