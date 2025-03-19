package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gofrs/uuid/v5"
	"gorm.io/datatypes"
)

type TransitionExample struct {
	global.GVA_MODEL
	FlowUUID uuid.UUID      `json:"flowUUID" form:"flowUUID" gorm:"column:flow_uuid;comment:流程uuid;"`              //流程uuid
	WorkUUID uuid.UUID      `json:"workUUID" form:"workUUID" gorm:"column:work_uuid;comment:工作uuid;"`              //工作uuid
	WorkType string         `json:"workType" form:"workType" gorm:"column:work_type;comment:工作类型;"`                //工作类型
	Operator uint           `json:"operator" form:"operator" gorm:"column:operator;comment:操作人;"`                  //操作人
	Attr     datatypes.JSON `json:"attr" form:"attr" gorm:"column:attr;comment:属性;type:text;"swaggertype:"object"` //属性
}

// 流转示例 结构体  TransitionExample
type TransitionExampleStep struct {
	global.GVA_MODEL
	FlowUUID uuid.UUID      `json:"flowUUID" form:"flowUUID" gorm:"column:flow_uuid;comment:流程uuid;"`              //流程uuid
	WorkUUID uuid.UUID      `json:"workUUID" form:"workUUID" gorm:"column:work_uuid;comment:工作uuid;"`              //工作uuid
	WorkType string         `json:"workType" form:"workType" gorm:"column:work_type;comment:工作类型;"`                //工作类型
	NodeID   string         `json:"nodeID" form:"nodeID" gorm:"column:node_id;comment:节点ID;"`                      //节点ID
	Operator uint           `json:"operator" form:"operator" gorm:"column:operator;comment:操作人;"`                  //操作人
	Attr     datatypes.JSON `json:"attr" form:"attr" gorm:"column:attr;comment:属性;type:text;"swaggertype:"object"` //属性
	Logic    string         `json:"logic" form:"logic" gorm:"column:logic;comment:逻辑;"`                            //逻辑
	Status   string         `json:"status" form:"status" gorm:"column:status;comment:状态:active,done,shape;"`       //状态 more对应节点shape
}

type NextFlow struct {
	PID    uint       `gorm:"primarykey" json:"PID"`                                     // 主键ID
	NodeID string     `json:"source" form:"source" gorm:"column:source;comment:目标节点;"`   // 目标节点
	Node   Node       `json:"node" form:"node" gorm:"column->;"`                         // 节点
	EdgeID string     `json:"edgeID" form:"edgeID" gorm:"column:edge_id;comment:过程边ID;"` // 边ID
	Edge   Edge       `json:"edge" form:"edge" gorm:"column->;"`                         // 边
	Next   []NextFlow `json:"next" form:"next" gorm:"column:next;comment:下一步;"`          // 下一步节点
}

// TableName 流转示例 TransitionExample自定义表名 gva_flow_transition_example
func (TransitionExample) TableName() string {
	return "gva_flow_transition_example"
}

// TableName 流转示例 TransitionExample自定义表名 gva_flow_transition_example
func (TransitionExampleStep) TableName() string {
	return "gva_flow_transition_example_step"
}
