package opsManage

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	OpsMattersTrackingRouter
	BankWorkOrdersRouter
}

var (
	opsMattersTrackingApi = api.ApiGroupApp.OpsManageApiGroup.OpsMattersTrackingApi
	bankWorkOrdersApi     = api.ApiGroupApp.OpsManageApiGroup.BankWorkOrdersApi
)
