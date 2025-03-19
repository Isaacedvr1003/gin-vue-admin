import service from '@/utils/request'

// @Tags Flow
// @Summary 创建流程管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Flow true "创建流程管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /flow/createFlow [post]
export const createFlow = (data) => {
  return service({
    url: '/flow/createFlow',
    method: 'post',
    data
  })
}

// @Tags Flow
// @Summary 删除流程管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Flow true "删除流程管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /flow/deleteFlow [delete]
export const deleteFlow = (params) => {
  return service({
    url: '/flow/deleteFlow',
    method: 'delete',
    params
  })
}

// @Tags Flow
// @Summary 批量删除流程管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除流程管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /flow/deleteFlow [delete]
export const deleteFlowByIds = (params) => {
  return service({
    url: '/flow/deleteFlowByIds',
    method: 'delete',
    params
  })
}

// @Tags Flow
// @Summary 更新流程管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Flow true "更新流程管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /flow/updateFlow [put]
export const updateFlow = (data) => {
  return service({
    url: '/flow/updateFlow',
    method: 'put',
    data
  })
}

// @Tags Flow
// @Summary 用id查询流程管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Flow true "用id查询流程管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /flow/findFlow [get]
export const findFlow = (params) => {
  return service({
    url: '/flow/findFlow',
    method: 'get',
    params
  })
}

// @Tags Flow
// @Summary 分页获取流程管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取流程管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /flow/getFlowList [get]
export const getFlowList = (params) => {
  return service({
    url: '/flow/getFlowList',
    method: 'get',
    params
  })
}


export const findGraph = (params) => {
  return service({
    url: '/flow/findGraph',
    method: 'get',
    params
  })
}


export const saveGraph = (data) => {
  return service({
    url: '/flow/saveGraph',
    method: 'post',
    data
  })
}



