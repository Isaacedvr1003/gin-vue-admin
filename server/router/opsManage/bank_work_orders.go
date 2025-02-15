package opsManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BankWorkOrdersRouter struct {}

// InitBankWorkOrdersRouter 初始化 工单管理 路由信息
func (s *BankWorkOrdersRouter) InitBankWorkOrdersRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	bankWorkOrdersRouter := Router.Group("bankWorkOrders").Use(middleware.OperationRecord())
	bankWorkOrdersRouterWithoutRecord := Router.Group("bankWorkOrders")
	bankWorkOrdersRouterWithoutAuth := PublicRouter.Group("bankWorkOrders")
	{
		bankWorkOrdersRouter.POST("createBankWorkOrders", bankWorkOrdersApi.CreateBankWorkOrders)   // 新建工单管理
		bankWorkOrdersRouter.DELETE("deleteBankWorkOrders", bankWorkOrdersApi.DeleteBankWorkOrders) // 删除工单管理
		bankWorkOrdersRouter.DELETE("deleteBankWorkOrdersByIds", bankWorkOrdersApi.DeleteBankWorkOrdersByIds) // 批量删除工单管理
		bankWorkOrdersRouter.PUT("updateBankWorkOrders", bankWorkOrdersApi.UpdateBankWorkOrders)    // 更新工单管理
	}
	{
		bankWorkOrdersRouterWithoutRecord.GET("findBankWorkOrders", bankWorkOrdersApi.FindBankWorkOrders)        // 根据ID获取工单管理
		bankWorkOrdersRouterWithoutRecord.GET("getBankWorkOrdersList", bankWorkOrdersApi.GetBankWorkOrdersList)  // 获取工单管理列表
	}
	{
	    bankWorkOrdersRouterWithoutAuth.GET("getBankWorkOrdersPublic", bankWorkOrdersApi.GetBankWorkOrdersPublic)  // 工单管理开放接口
	}
}
