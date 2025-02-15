package opsManage

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ OpsMattersTrackingRouter }

var opsMattersTrackingApi = api.ApiGroupApp.OpsManageApiGroup.OpsMattersTrackingApi
