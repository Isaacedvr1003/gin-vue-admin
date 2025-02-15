
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type OpsMattersTrackingSearch struct{
    Id  *int `json:"id" form:"id" `
    Description  *string `json:"description" form:"description" `
    Priority  *string `json:"priority" form:"priority" `
    Status  *string `json:"status" form:"status" `
    AssignedTo  *string `json:"assignedTo" form:"assignedTo" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
