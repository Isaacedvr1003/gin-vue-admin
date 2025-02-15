package opsManage

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/opsManage"
    opsManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/opsManage/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type BankWorkOrdersApi struct {}



// CreateBankWorkOrders 创建工单管理
// @Tags BankWorkOrders
// @Summary 创建工单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body opsManage.BankWorkOrders true "创建工单管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /bankWorkOrders/createBankWorkOrders [post]
func (bankWorkOrdersApi *BankWorkOrdersApi) CreateBankWorkOrders(c *gin.Context) {
	var bankWorkOrders opsManage.BankWorkOrders
	err := c.ShouldBindJSON(&bankWorkOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bankWorkOrdersService.CreateBankWorkOrders(&bankWorkOrders)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteBankWorkOrders 删除工单管理
// @Tags BankWorkOrders
// @Summary 删除工单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body opsManage.BankWorkOrders true "删除工单管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /bankWorkOrders/deleteBankWorkOrders [delete]
func (bankWorkOrdersApi *BankWorkOrdersApi) DeleteBankWorkOrders(c *gin.Context) {
	ID := c.Query("ID")
	err := bankWorkOrdersService.DeleteBankWorkOrders(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBankWorkOrdersByIds 批量删除工单管理
// @Tags BankWorkOrders
// @Summary 批量删除工单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /bankWorkOrders/deleteBankWorkOrdersByIds [delete]
func (bankWorkOrdersApi *BankWorkOrdersApi) DeleteBankWorkOrdersByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := bankWorkOrdersService.DeleteBankWorkOrdersByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBankWorkOrders 更新工单管理
// @Tags BankWorkOrders
// @Summary 更新工单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body opsManage.BankWorkOrders true "更新工单管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bankWorkOrders/updateBankWorkOrders [put]
func (bankWorkOrdersApi *BankWorkOrdersApi) UpdateBankWorkOrders(c *gin.Context) {
	var bankWorkOrders opsManage.BankWorkOrders
	err := c.ShouldBindJSON(&bankWorkOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bankWorkOrdersService.UpdateBankWorkOrders(bankWorkOrders)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBankWorkOrders 用id查询工单管理
// @Tags BankWorkOrders
// @Summary 用id查询工单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query string true "用id查询工单管理"
// @Success 200 {object} response.Response{data=opsManage.BankWorkOrders,msg=string} "查询成功"
// @Router /bankWorkOrders/findBankWorkOrders [get]
func (bankWorkOrdersApi *BankWorkOrdersApi) FindBankWorkOrders(c *gin.Context) {
	ID := c.Query("ID")
	rebankWorkOrders, err := bankWorkOrdersService.GetBankWorkOrders(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rebankWorkOrders, c)
}
// GetBankWorkOrdersList 分页获取工单管理列表
// @Tags BankWorkOrders
// @Summary 分页获取工单管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query opsManageReq.BankWorkOrdersSearch true "分页获取工单管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /bankWorkOrders/getBankWorkOrdersList [get]
func (bankWorkOrdersApi *BankWorkOrdersApi) GetBankWorkOrdersList(c *gin.Context) {
	var pageInfo opsManageReq.BankWorkOrdersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := bankWorkOrdersService.GetBankWorkOrdersInfoList(pageInfo)
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

// GetBankWorkOrdersPublic 不需要鉴权的工单管理接口
// @Tags BankWorkOrders
// @Summary 不需要鉴权的工单管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bankWorkOrders/getBankWorkOrdersPublic [get]
func (bankWorkOrdersApi *BankWorkOrdersApi) GetBankWorkOrdersPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    bankWorkOrdersService.GetBankWorkOrdersPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的工单管理接口信息",
    }, "获取成功", c)
}
