package opsManage

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ OpsMattersTrackingApi }

var opsMattersTrackingService = service.ServiceGroupApp.OpsManageServiceGroup.OpsMattersTrackingService
