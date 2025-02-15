package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BankWorkOrdersSearch struct {
	BankName          *string    `json:"BankName" form:"BankName" `
	ID                *string    `json:"ID" form:"ID" `
	Summary           *string    `json:"Summary" form:"Summary" `
	BusinessContact   *string    `json:"BusinessContact" form:"BusinessContact" `
	Details           *string    `json:"Details" form:"Details" `
	Status            *string    `json:"Status" form:"Status" `
	StartSubmitTime   *time.Time `json:"startSubmitTime" form:"startSubmitTime"`
	EndSubmitTime     *time.Time `json:"endSubmitTime" form:"endSubmitTime"`
	SubmitPerson      *string    `json:"SubmitPerson" form:"SubmitPerson" `
	StartTransferTime *time.Time `json:"startTransferTime" form:"startTransferTime"`
	EndTransferTime   *time.Time `json:"endTransferTime" form:"endTransferTime"`
	StartDeliveryTime *time.Time `json:"startDeliveryTime" form:"startDeliveryTime"`
	EndDeliveryTime   *time.Time `json:"endDeliveryTime" form:"endDeliveryTime"`
	DeliveryID        *string    `json:"DeliveryID" form:"DeliveryID" `
	StartCompleteTime *time.Time `json:"startCompleteTime" form:"startCompleteTime"`
	EndCompleteTime   *time.Time `json:"endCompleteTime" form:"endCompleteTime"`
	CompleteID        *string    `json:"CompleteID" form:"CompleteID" `
	StartUpdateTime   *time.Time `json:"startUpdateTime" form:"startUpdateTime"`
	EndUpdateTime     *time.Time `json:"endUpdateTime" form:"endUpdateTime"`
	Notes             *string    `json:"notes" form:"notes" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
