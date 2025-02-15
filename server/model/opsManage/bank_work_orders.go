
// 自动生成模板BankWorkOrders
package opsManage
import (
	"time"
)

// 工单管理 结构体  BankWorkOrders
type BankWorkOrders struct {
    BankName  *string `json:"BankName" form:"BankName" gorm:"column:BankName;comment:银行名称;size:255;" binding:"required"`  //银行名称
    ID  *string `json:"ID" form:"ID" gorm:"primarykey;column:ID;comment:工单编号;size:255;" binding:"required"`  //工单编号
    Summary  *string `json:"Summary" form:"Summary" gorm:"column:Summary;comment:工单摘要;size:500;" binding:"required"`  //工单摘要
    BusinessContact  *string `json:"BusinessContact" form:"BusinessContact" gorm:"column:BusinessContact;comment:业务联系人;size:255;" binding:"required"`  //业务联系人
    Details  *string `json:"Details" form:"Details" gorm:"column:Details;comment:工单详情;type:text;" binding:"required"`  //工单详情
    Status  *string `json:"Status" form:"Status" gorm:"column:Status;comment:工单进度;size:255;" binding:"required"`  //工单进度
    SubmitTime  *time.Time `json:"SubmitTime" form:"SubmitTime" gorm:"column:SubmitTime;comment:工单申请时间;size:255;" binding:"required"`  //工单申请时间
    SubmitPerson  *string `json:"SubmitPerson" form:"SubmitPerson" gorm:"column:SubmitPerson;comment:工单负责人;size:255;" binding:"required"`  //工单负责人
    TransferTime  *time.Time `json:"TransferTime" form:"TransferTime" gorm:"column:TransferTime;comment:工单转交付时间;size:255;"`  //工单转交付时间
    DeliveryTime  *time.Time `json:"DeliveryTime" form:"DeliveryTime" gorm:"column:DeliveryTime;comment:工单交付时间;size:255;"`  //工单交付时间
    DeliveryID  *string `json:"DeliveryID" form:"DeliveryID" gorm:"column:DeliveryID;comment:交付企微审批编号;size:255;"`  //交付企微审批编号
    CompleteTime  *time.Time `json:"CompleteTime" form:"CompleteTime" gorm:"column:CompleteTime;comment:工单上线时间;size:255;"`  //工单上线时间
    CompleteID  *string `json:"CompleteID" form:"CompleteID" gorm:"column:CompleteID;comment:上线堡垒机编号;size:255;"`  //上线堡垒机编号
    UpdateTime  *time.Time `json:"UpdateTime" form:"UpdateTime" gorm:"column:UpdateTime;comment:工单更新时间;size:255;"`  //工单更新时间
    Notes  *string `json:"notes" form:"notes" gorm:"column:notes;comment:工单备注;size:255;type:text;"`  //工单备注
}


// TableName 工单管理 BankWorkOrders自定义表名 bank_work_orders
func (BankWorkOrders) TableName() string {
    return "bank_work_orders"
}





