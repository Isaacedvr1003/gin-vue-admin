package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/gofrs/uuid/v5"
	"gorm.io/datatypes"
	"time"
)

type TransitionExampleNext struct {
	StepID   uint           `json:"stepID" form:"stepID" gorm:"column:step_id;comment:步骤ID;"`                      //步骤ID
	NodeID   string         `json:"nodeID" form:"nodeID" gorm:"column:node_id;comment:节点ID;"`                      //节点ID
	WorkUUID uuid.UUID      `json:"workUUID" form:"workUUID" gorm:"column:work_uuid;comment:工作uuid;"`              //工作uuid
	Attr     datatypes.JSON `json:"attr" form:"attr" gorm:"column:attr;comment:属性;type:text;"swaggertype:"object"` //属性
	Logic    string         `json:"logic" form:"logic" gorm:"column:logic;comment:逻辑;"`                            //逻辑
}

type TransitionExampleSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
}
