package flow

import (
	gvaGlobal "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/flow/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/flow/router"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
	"github.com/gin-gonic/gin"
)

type FlowPlugin struct {
}

func CreateFlowPlug() *FlowPlugin {

	gvaGlobal.GVA_DB.AutoMigrate(
		model.Flow{},
		model.Graph{},
		model.Edge{},
		model.Node{},
		model.NodeData{},
		model.EdgeData{},
		model.Source{},
		model.Target{},
		model.TransitionExample{},
		model.TransitionExampleStep{},
	) // 此处可以把插件依赖的数据库结构体自动创建表 需要填写对应的结构体

	// 下方会自动注册菜单 第一个参数为菜单一级路由信息一般为定义好的组名 第二个参数为真实使用的web页面路由信息
	// 具体值请根据实际情况修改
	utils.RegisterMenus(
		system.SysBaseMenu{
			Path:      "flowGroup",
			Name:      "flowGroup",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      1000,
			Meta: system.Meta{
				Title: "流程管理",
				Icon:  "school",
			},
		},
		system.SysBaseMenu{
			Path:      "flow",
			Name:      "flow",
			Hidden:    false,
			Component: "plugin/flow/view/flow.vue",
			Sort:      0,
			Meta: system.Meta{
				Title: "流程列表",
				Icon:  "school",
			},
		},
		system.SysBaseMenu{
			Path:      "flowSetting/:uuid",
			Name:      "flowSetting",
			Hidden:    true,
			Component: "plugin/flow/view/flowSetting.vue",
			Sort:      1,
			Meta: system.Meta{
				ActiveName: "flow",
				Title:      "流程绘制",
				Icon:       "school",
			},
		},
		system.SysBaseMenu{
			Path:      "flowTest/:workUUID",
			Name:      "flowTest",
			Hidden:    true,
			Component: "plugin/flow/view/flowTest.vue",
			Sort:      2,
			Meta: system.Meta{
				ActiveName: "flowTestList",
				Title:      "流转详情",
				Icon:       "finished",
			},
		},
		system.SysBaseMenu{
			Path:      "flowTestList",
			Name:      "flowTestList",
			Hidden:    false,
			Component: "plugin/flow/view/flowTestList.vue",
			Sort:      3,
			Meta: system.Meta{
				Title: "示例列表",
				Icon:  "list",
			},
		},
	)

	// 下方会自动注册api 以下格式为示例格式，请按照实际情况修改
	utils.RegisterApis(
		system.SysApi{
			Path:        "/flow/createFlow",
			Description: "创建流程",
			ApiGroup:    "流程",
			Method:      "POST",
		},
		system.SysApi{
			Path:        "/flow/deleteFlow",
			Description: "删除流程",
			ApiGroup:    "流程",
			Method:      "DELETE",
		},
		system.SysApi{
			Path:        "/flow/deleteFlowByIds",
			Description: "批量删除流程",
			ApiGroup:    "流程",
			Method:      "DELETE",
		},
		system.SysApi{
			Path:        "/flow/updateFlow",
			Description: "更新流程",
			ApiGroup:    "流程",
			Method:      "PUT",
		},
		system.SysApi{
			Path:        "/flow/findFlow",
			Description: "获取单个流程",
			ApiGroup:    "流程",
			Method:      "GET",
		},
		system.SysApi{
			Path:        "/flow/getFlowList",
			Description: "获取流程列表",
			ApiGroup:    "流程",
			Method:      "GET",
		},
		system.SysApi{
			Path:        "/flow/findGraph",
			Description: "获取流程图",
			ApiGroup:    "流程",
			Method:      "GET",
		},
		system.SysApi{
			Path:        "/flow/saveGraph",
			Description: "保存更改流程图",
			ApiGroup:    "流程",
			Method:      "POST",
		},

		system.SysApi{
			Path:        "/flow/startTransitionExample",
			Description: "开始示例流程业务",
			ApiGroup:    "流程",
			Method:      "POST",
		},

		system.SysApi{
			Path:        "/flow/getTransitionExampleStep",
			Description: "查找业务流程所有步骤",
			ApiGroup:    "流程",
			Method:      "GET",
		},

		system.SysApi{
			Path:        "/flow/getTransitionExampleList",
			Description: "示例业务流程",
			ApiGroup:    "流程",
			Method:      "GET",
		},
		system.SysApi{
			Path:        "/flow/transitionExampleNext",
			Description: "示例业务流程流转（下一步）",
			ApiGroup:    "流程",
			Method:      "POST",
		},
	)

	return &FlowPlugin{}
}

func (*FlowPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitFlowRouter(group)
}

func (*FlowPlugin) RouterPath() string {
	return "flow"
}
