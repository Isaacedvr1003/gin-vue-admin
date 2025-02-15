
// 自动生成模板OpsMattersTracking
package opsManage
import (
	"time"
)

// 事项管理 结构体  OpsMattersTracking
type OpsMattersTracking struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:事项ID;size:20;"`  //事项ID
    Description  *string `json:"description" form:"description" gorm:"column:description;comment:事项简述;" binding:"required"`  //事项简述
    Priority  *string `json:"priority" form:"priority" gorm:"column:priority;comment:事项级别;size:16;" binding:"required"`  //事项级别
    Status  *string `json:"status" form:"status" gorm:"column:status;comment:事项状态;size:16;" binding:"required"`  //事项状态
    ResolutionSteps  *string `json:"resolutionSteps" form:"resolutionSteps" gorm:"column:resolution_steps;comment:解决过程;type:text;" binding:"required"`  //解决过程
    AssignedTo  *string `json:"assignedTo" form:"assignedTo" gorm:"column:assigned_to;comment:跟踪人员;size:255;" binding:"required"`  //跟踪人员
    CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;" binding:"required"`  //创建时间
    UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:更新时间;" binding:"required"`  //更新时间
    DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:删除时间;" binding:"required"`  //删除时间
    Notes  *string `json:"notes" form:"notes" gorm:"column:notes;comment:事项备注;type:text;"`  //事项备注
}


// TableName 事项管理 OpsMattersTracking自定义表名 ops_matters_tracking
func (OpsMattersTracking) TableName() string {
    return "ops_matters_tracking"
}





