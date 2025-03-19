// 自动生成模板Flow
package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gofrs/uuid/v5"
)

// 流程管理 结构体  Flow
type Flow struct {
	global.GVA_MODEL
	Name string    `json:"name" form:"name" gorm:"column:name;comment:name;" binding:"required"` //流程名称
	UUID uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:uuid;"`                    //uuid
	Desc string    `json:"desc" form:"desc" gorm:"column:desc;comment:流程介绍;"`                    //流程介绍
}

// TableName 流程管理 Flow自定义表名 ai_flow
func (Flow) TableName() string {
	return "gva_flow"
}
