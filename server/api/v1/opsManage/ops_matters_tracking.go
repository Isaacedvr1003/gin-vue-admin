package opsManage

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/opsManage"
    opsManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/opsManage/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type OpsMattersTrackingApi struct {}



// CreateOpsMattersTracking 创建事项管理
// @Tags OpsMattersTracking
// @Summary 创建事项管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body opsManage.OpsMattersTracking true "创建事项管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /opsMattersTracking/createOpsMattersTracking [post]
func (opsMattersTrackingApi *OpsMattersTrackingApi) CreateOpsMattersTracking(c *gin.Context) {
	var opsMattersTracking opsManage.OpsMattersTracking
	err := c.ShouldBindJSON(&opsMattersTracking)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = opsMattersTrackingService.CreateOpsMattersTracking(&opsMattersTracking)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteOpsMattersTracking 删除事项管理
// @Tags OpsMattersTracking
// @Summary 删除事项管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body opsManage.OpsMattersTracking true "删除事项管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /opsMattersTracking/deleteOpsMattersTracking [delete]
func (opsMattersTrackingApi *OpsMattersTrackingApi) DeleteOpsMattersTracking(c *gin.Context) {
	id := c.Query("id")
	err := opsMattersTrackingService.DeleteOpsMattersTracking(id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteOpsMattersTrackingByIds 批量删除事项管理
// @Tags OpsMattersTracking
// @Summary 批量删除事项管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /opsMattersTracking/deleteOpsMattersTrackingByIds [delete]
func (opsMattersTrackingApi *OpsMattersTrackingApi) DeleteOpsMattersTrackingByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := opsMattersTrackingService.DeleteOpsMattersTrackingByIds(ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateOpsMattersTracking 更新事项管理
// @Tags OpsMattersTracking
// @Summary 更新事项管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body opsManage.OpsMattersTracking true "更新事项管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /opsMattersTracking/updateOpsMattersTracking [put]
func (opsMattersTrackingApi *OpsMattersTrackingApi) UpdateOpsMattersTracking(c *gin.Context) {
	var opsMattersTracking opsManage.OpsMattersTracking
	err := c.ShouldBindJSON(&opsMattersTracking)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = opsMattersTrackingService.UpdateOpsMattersTracking(opsMattersTracking)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindOpsMattersTracking 用id查询事项管理
// @Tags OpsMattersTracking
// @Summary 用id查询事项管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询事项管理"
// @Success 200 {object} response.Response{data=opsManage.OpsMattersTracking,msg=string} "查询成功"
// @Router /opsMattersTracking/findOpsMattersTracking [get]
func (opsMattersTrackingApi *OpsMattersTrackingApi) FindOpsMattersTracking(c *gin.Context) {
	id := c.Query("id")
	reopsMattersTracking, err := opsMattersTrackingService.GetOpsMattersTracking(id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reopsMattersTracking, c)
}
// GetOpsMattersTrackingList 分页获取事项管理列表
// @Tags OpsMattersTracking
// @Summary 分页获取事项管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query opsManageReq.OpsMattersTrackingSearch true "分页获取事项管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /opsMattersTracking/getOpsMattersTrackingList [get]
func (opsMattersTrackingApi *OpsMattersTrackingApi) GetOpsMattersTrackingList(c *gin.Context) {
	var pageInfo opsManageReq.OpsMattersTrackingSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := opsMattersTrackingService.GetOpsMattersTrackingInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetOpsMattersTrackingPublic 不需要鉴权的事项管理接口
// @Tags OpsMattersTracking
// @Summary 不需要鉴权的事项管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /opsMattersTracking/getOpsMattersTrackingPublic [get]
func (opsMattersTrackingApi *OpsMattersTrackingApi) GetOpsMattersTrackingPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    opsMattersTrackingService.GetOpsMattersTrackingPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的事项管理接口信息",
    }, "获取成功", c)
}
