package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/flow/model"
	flowReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/flow/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/flow/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FlowApi struct {
}

var flowService = service.ServiceGroupApp.FlowService

// CreateFlow 创建流程管理
// @Tags Flow
// @Summary 创建流程管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Flow true "创建流程管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /flow/createFlow [post]
func (flowApi *FlowApi) CreateFlow(c *gin.Context) {
	var flow model.Flow
	err := c.ShouldBindJSON(&flow)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := flowService.CreateFlow(&flow); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFlow 删除流程管理
// @Tags Flow
// @Summary 删除流程管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Flow true "删除流程管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /flow/deleteFlow [delete]
func (flowApi *FlowApi) DeleteFlow(c *gin.Context) {
	ID := c.Query("ID")
	if err := flowService.DeleteFlow(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFlowByIds 批量删除流程管理
// @Tags Flow
// @Summary 批量删除流程管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /flow/deleteFlowByIds [delete]
func (flowApi *FlowApi) DeleteFlowByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := flowService.DeleteFlowByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFlow 更新流程管理
// @Tags Flow
// @Summary 更新流程管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Flow true "更新流程管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /flow/updateFlow [put]
func (flowApi *FlowApi) UpdateFlow(c *gin.Context) {
	var flow model.Flow
	err := c.ShouldBindJSON(&flow)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := flowService.UpdateFlow(flow); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFlow 用id查询流程管理
// @Tags Flow
// @Summary 用id查询流程管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Graph true "创建或更改流程图"
// @Success 200 {object} response.Response{data=object{reflow=model.Flow},msg=string} "查询成功"
// @Router /flow/saveGraph [post]
func (flowApi *FlowApi) SaveGraph(c *gin.Context) {
	var graph model.Graph
	e := c.ShouldBindJSON(&graph)
	if e != nil {
		response.FailWithMessage(e.Error(), c)
		return
	}
	if err := flowService.SaveGraph(graph); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.Ok(c)
	}
}

// FindFlow 用id查询流程管理
// @Tags Flow
// @Summary 用id查询流程管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Flow true "用id查询流程管理"
// @Success 200 {object} response.Response{data=object{reflow=model.Flow},msg=string} "查询成功"
// @Router /flow/findFlow [get]
func (flowApi *FlowApi) FindFlow(c *gin.Context) {
	ID := c.Query("ID")
	uuid := c.Query("uuid")
	if reflow, err := flowService.GetFlow(ID, uuid); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reflow": reflow}, c)
	}
}

// FindFlow 用id查询流程管理
// @Tags Flow
// @Summary 用id查询流程管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param uuid query string true "用uuid查询流程管理"
// @Success 200 {object} response.Response{data=object{reflow=model.Flow},msg=string} "查询成功"
// @Router /flow/findGraph [get]
func (flowApi *FlowApi) FindGraph(c *gin.Context) {
	uuid := c.Query("uuid")
	if reflow, err := flowService.FindGraph(uuid); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reflow": reflow}, c)
	}
}

// GetFlowList 分页获取流程管理列表
// @Tags Flow
// @Summary 分页获取流程管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query flowReq.FlowSearch true "分页获取流程管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /flow/getFlowList [get]
func (flowApi *FlowApi) GetFlowList(c *gin.Context) {
	var pageInfo flowReq.FlowSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := flowService.GetFlowInfoList(pageInfo); err != nil {
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
