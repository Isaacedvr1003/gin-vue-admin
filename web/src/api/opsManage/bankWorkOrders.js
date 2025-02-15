import service from '@/utils/request'
// @Tags BankWorkOrders
// @Summary 创建工单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.BankWorkOrders true "创建工单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bankWorkOrders/createBankWorkOrders [post]
export const createBankWorkOrders = (data) => {
  return service({
    url: '/bankWorkOrders/createBankWorkOrders',
    method: 'post',
    data
  })
}

// @Tags BankWorkOrders
// @Summary 删除工单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.BankWorkOrders true "删除工单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bankWorkOrders/deleteBankWorkOrders [delete]
export const deleteBankWorkOrders = (params) => {
  return service({
    url: '/bankWorkOrders/deleteBankWorkOrders',
    method: 'delete',
    params
  })
}

// @Tags BankWorkOrders
// @Summary 批量删除工单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除工单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bankWorkOrders/deleteBankWorkOrders [delete]
export const deleteBankWorkOrdersByIds = (params) => {
  return service({
    url: '/bankWorkOrders/deleteBankWorkOrdersByIds',
    method: 'delete',
    params
  })
}

// @Tags BankWorkOrders
// @Summary 更新工单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.BankWorkOrders true "更新工单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bankWorkOrders/updateBankWorkOrders [put]
export const updateBankWorkOrders = (data) => {
  return service({
    url: '/bankWorkOrders/updateBankWorkOrders',
    method: 'put',
    data
  })
}

// @Tags BankWorkOrders
// @Summary 用id查询工单管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.BankWorkOrders true "用id查询工单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bankWorkOrders/findBankWorkOrders [get]
export const findBankWorkOrders = (params) => {
  return service({
    url: '/bankWorkOrders/findBankWorkOrders',
    method: 'get',
    params
  })
}

// @Tags BankWorkOrders
// @Summary 分页获取工单管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取工单管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bankWorkOrders/getBankWorkOrdersList [get]
export const getBankWorkOrdersList = (params) => {
  return service({
    url: '/bankWorkOrders/getBankWorkOrdersList',
    method: 'get',
    params
  })
}

// @Tags BankWorkOrders
// @Summary 不需要鉴权的工单管理接口
// @Accept application/json
// @Produce application/json
// @Param data query opsManageReq.BankWorkOrdersSearch true "分页获取工单管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bankWorkOrders/getBankWorkOrdersPublic [get]
export const getBankWorkOrdersPublic = () => {
  return service({
    url: '/bankWorkOrders/getBankWorkOrdersPublic',
    method: 'get',
  })
}
