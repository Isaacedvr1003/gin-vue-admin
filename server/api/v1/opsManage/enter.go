package opsManage

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	OpsMattersTrackingApi
	BankWorkOrdersApi
}

var (
	opsMattersTrackingService = service.ServiceGroupApp.OpsManageServiceGroup.OpsMattersTrackingService
	bankWorkOrdersService     = service.ServiceGroupApp.OpsManageServiceGroup.BankWorkOrdersService
)
