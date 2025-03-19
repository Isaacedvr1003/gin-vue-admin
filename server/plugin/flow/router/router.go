package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/flow/api"
	"github.com/gin-gonic/gin"
)

type FlowRouter struct {
}

// InitFlowRouter 初始化 流程管理 路由信息
func (s *FlowRouter) InitFlowRouter(Router *gin.RouterGroup) {
	flowRouter := Router.Use(middleware.OperationRecord())
	flowRouterWithoutRecord := Router

	var flowApi = api.ApiGroupApp.FlowApi
	{
		flowRouter.POST("createFlow", flowApi.CreateFlow)             // 新建流程管理
		flowRouter.DELETE("deleteFlow", flowApi.DeleteFlow)           // 删除流程管理
		flowRouter.DELETE("deleteFlowByIds", flowApi.DeleteFlowByIds) // 批量删除流程管理
		flowRouter.PUT("updateFlow", flowApi.UpdateFlow)              // 更新流程管理
		flowRouterWithoutRecord.POST("saveGraph", flowApi.SaveGraph)  // 创建/修改流程图
	}
	{
		flowRouterWithoutRecord.GET("findFlow", flowApi.FindFlow)       // 根据ID获取流程管理
		flowRouterWithoutRecord.GET("getFlowList", flowApi.GetFlowList) // 获取流程管理列表
		flowRouterWithoutRecord.GET("findGraph", flowApi.FindGraph)     // 根据uuid获取流程图
	}

	// 下方为流转示例 具体流转业务关联性较强 请自行根据业务参考示例进行开发
	{
		flowRouter.POST("startTransitionExample", flowApi.StartTransitionExample)    // 开始流程示例
		flowRouter.GET("getTransitionExampleStep", flowApi.GetTransitionExampleStep) // 获取流程所有节点
		flowRouter.GET("getTransitionExampleList", flowApi.GetTransitionExampleList) // 获取流程列表
		flowRouter.POST("transitionExampleNext", flowApi.TransitionExampleNext)      // 下一步
	}
}
