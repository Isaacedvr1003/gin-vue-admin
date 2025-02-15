import service from '@/utils/request'
// @Tags OpsMattersTracking
// @Summary 创建事项管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.OpsMattersTracking true "创建事项管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /opsMattersTracking/createOpsMattersTracking [post]
export const createOpsMattersTracking = (data) => {
  return service({
    url: '/opsMattersTracking/createOpsMattersTracking',
    method: 'post',
    data
  })
}

// @Tags OpsMattersTracking
// @Summary 删除事项管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.OpsMattersTracking true "删除事项管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /opsMattersTracking/deleteOpsMattersTracking [delete]
export const deleteOpsMattersTracking = (params) => {
  return service({
    url: '/opsMattersTracking/deleteOpsMattersTracking',
    method: 'delete',
    params
  })
}

// @Tags OpsMattersTracking
// @Summary 批量删除事项管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除事项管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /opsMattersTracking/deleteOpsMattersTracking [delete]
export const deleteOpsMattersTrackingByIds = (params) => {
  return service({
    url: '/opsMattersTracking/deleteOpsMattersTrackingByIds',
    method: 'delete',
    params
  })
}

// @Tags OpsMattersTracking
// @Summary 更新事项管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.OpsMattersTracking true "更新事项管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /opsMattersTracking/updateOpsMattersTracking [put]
export const updateOpsMattersTracking = (data) => {
  return service({
    url: '/opsMattersTracking/updateOpsMattersTracking',
    method: 'put',
    data
  })
}

// @Tags OpsMattersTracking
// @Summary 用id查询事项管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.OpsMattersTracking true "用id查询事项管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /opsMattersTracking/findOpsMattersTracking [get]
export const findOpsMattersTracking = (params) => {
  return service({
    url: '/opsMattersTracking/findOpsMattersTracking',
    method: 'get',
    params
  })
}

// @Tags OpsMattersTracking
// @Summary 分页获取事项管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取事项管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /opsMattersTracking/getOpsMattersTrackingList [get]
export const getOpsMattersTrackingList = (params) => {
  return service({
    url: '/opsMattersTracking/getOpsMattersTrackingList',
    method: 'get',
    params
  })
}

// @Tags OpsMattersTracking
// @Summary 不需要鉴权的事项管理接口
// @Accept application/json
// @Produce application/json
// @Param data query opsManageReq.OpsMattersTrackingSearch true "分页获取事项管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /opsMattersTracking/getOpsMattersTrackingPublic [get]
export const getOpsMattersTrackingPublic = () => {
  return service({
    url: '/opsMattersTracking/getOpsMattersTrackingPublic',
    method: 'get',
  })
}
