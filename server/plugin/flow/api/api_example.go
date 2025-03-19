package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/flow/model"
	flowReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/flow/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// StarFlow 开始流程
// @Tags Flow
// @Summary 开始流程
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Flow true "开始流程"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /flow/startTransitionExample [post]
func (flowApi *FlowApi) StartTransitionExample(c *gin.Context) {
	var flow model.TransitionExample
	err := c.ShouldBindJSON(&flow)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	operation := utils.GetUserID(c)

	if workUUID, err := flowService.StartTransitionExample(flow, operation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithData(gin.H{
			"workUUID": workUUID,
		}, c)
	}
}

// StarFlow 查询流转详情
// @Tags Flow
// @Summary 查询流转详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query {} true "查询流转详情"
// @Success 200 {object} response.Response{data=object{retransitionExample=flow.TransitionExample},msg=string} "查询流转详情"
// @Router /flow/getTransitionExampleStep [get]
func (flowApi *FlowApi) GetTransitionExampleStep(c *gin.Context) {
	workUUID := c.Query("workUUID")
	if step, err := flowService.GetTransitionExampleStep(workUUID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"step": step}, c)
	}
}

// StarFlow 下一步
// @Tags Flow
// @Summary 下一步
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query {} true "下一步"
// @Success 200 {object} response.Response{data=object{retransitionExample=flow.TransitionExample},msg=string} "查询流转详情"
// @Router /flow/transitionExampleNext [post]
func (flowApi *FlowApi) TransitionExampleNext(c *gin.Context) {
	var next flowReq.TransitionExampleNext
	err := c.ShouldBindJSON(&next)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	operator := utils.GetUserID(c)

	if err := flowService.TransitionExampleNext(nil, next, operator); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.Ok(c)
	}
}

// StarFlow 分页获取流转示例列表
// @Tags StarFlow
// @Summary 分页获取流转示例列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query flowReq.TransitionExampleSearch true "分页获取流转示例列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /flow/getTransitionExampleList [get]
func (flowApi *FlowApi) GetTransitionExampleList(c *gin.Context) {
	var pageInfo flowReq.TransitionExampleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := flowService.GetTransitionExampleInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
