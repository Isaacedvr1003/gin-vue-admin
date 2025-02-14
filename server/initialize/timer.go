package initialize

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/systask/service"

func Timer() {
	service.ServiceGroupApp.SysTaskService.Initialize()
}
