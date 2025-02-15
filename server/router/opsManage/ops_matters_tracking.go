package opsManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OpsMattersTrackingRouter struct {}

// InitOpsMattersTrackingRouter 初始化 事项管理 路由信息
func (s *OpsMattersTrackingRouter) InitOpsMattersTrackingRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	opsMattersTrackingRouter := Router.Group("opsMattersTracking").Use(middleware.OperationRecord())
	opsMattersTrackingRouterWithoutRecord := Router.Group("opsMattersTracking")
	opsMattersTrackingRouterWithoutAuth := PublicRouter.Group("opsMattersTracking")
	{
		opsMattersTrackingRouter.POST("createOpsMattersTracking", opsMattersTrackingApi.CreateOpsMattersTracking)   // 新建事项管理
		opsMattersTrackingRouter.DELETE("deleteOpsMattersTracking", opsMattersTrackingApi.DeleteOpsMattersTracking) // 删除事项管理
		opsMattersTrackingRouter.DELETE("deleteOpsMattersTrackingByIds", opsMattersTrackingApi.DeleteOpsMattersTrackingByIds) // 批量删除事项管理
		opsMattersTrackingRouter.PUT("updateOpsMattersTracking", opsMattersTrackingApi.UpdateOpsMattersTracking)    // 更新事项管理
	}
	{
		opsMattersTrackingRouterWithoutRecord.GET("findOpsMattersTracking", opsMattersTrackingApi.FindOpsMattersTracking)        // 根据ID获取事项管理
		opsMattersTrackingRouterWithoutRecord.GET("getOpsMattersTrackingList", opsMattersTrackingApi.GetOpsMattersTrackingList)  // 获取事项管理列表
	}
	{
	    opsMattersTrackingRouterWithoutAuth.GET("getOpsMattersTrackingPublic", opsMattersTrackingApi.GetOpsMattersTrackingPublic)  // 事项管理开放接口
	}
}
