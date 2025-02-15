package opsManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/opsManage"
	opsManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/opsManage/request"
)

type BankWorkOrdersService struct{}

// CreateBankWorkOrders 创建工单管理记录
// Author [yourname](https://github.com/yourname)
func (bankWorkOrdersService *BankWorkOrdersService) CreateBankWorkOrders(bankWorkOrders *opsManage.BankWorkOrders) (err error) {
	err = global.GVA_DB.Create(bankWorkOrders).Error
	return err
}

// DeleteBankWorkOrders 删除工单管理记录
// Author [yourname](https://github.com/yourname)
func (bankWorkOrdersService *BankWorkOrdersService) DeleteBankWorkOrders(ID string) (err error) {
	err = global.GVA_DB.Delete(&opsManage.BankWorkOrders{}, "ID = ?", ID).Error
	return err
}

// DeleteBankWorkOrdersByIds 批量删除工单管理记录
// Author [yourname](https://github.com/yourname)
func (bankWorkOrdersService *BankWorkOrdersService) DeleteBankWorkOrdersByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]opsManage.BankWorkOrders{}, "ID in ?", IDs).Error
	return err
}

// UpdateBankWorkOrders 更新工单管理记录
// Author [yourname](https://github.com/yourname)
func (bankWorkOrdersService *BankWorkOrdersService) UpdateBankWorkOrders(bankWorkOrders opsManage.BankWorkOrders) (err error) {
	err = global.GVA_DB.Model(&opsManage.BankWorkOrders{}).Where("ID = ?", bankWorkOrders.ID).Updates(&bankWorkOrders).Error
	return err
}

// GetBankWorkOrders 根据ID获取工单管理记录
// Author [yourname](https://github.com/yourname)
func (bankWorkOrdersService *BankWorkOrdersService) GetBankWorkOrders(ID string) (bankWorkOrders opsManage.BankWorkOrders, err error) {
	err = global.GVA_DB.Where("ID = ?", ID).First(&bankWorkOrders).Error
	return
}

// GetBankWorkOrdersInfoList 分页获取工单管理记录
// Author [yourname](https://github.com/yourname)
func (bankWorkOrdersService *BankWorkOrdersService) GetBankWorkOrdersInfoList(info opsManageReq.BankWorkOrdersSearch) (list []opsManage.BankWorkOrders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&opsManage.BankWorkOrders{})
	var bankWorkOrderss []opsManage.BankWorkOrders
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.BankName != nil && *info.BankName != "" {
		db = db.Where("BankName = ?", *info.BankName)
	}
	if info.ID != nil && *info.ID != "" {
		db = db.Where("ID LIKE ?", "%"+*info.ID+"%")
	}
	if info.Summary != nil && *info.Summary != "" {
		db = db.Where("Summary LIKE ?", "%"+*info.Summary+"%")
	}
	if info.BusinessContact != nil && *info.BusinessContact != "" {
		db = db.Where("BusinessContact LIKE ?", "%"+*info.BusinessContact+"%")
	}
	if info.Details != nil {
		db = db.Where("Details LIKE ?", "%"+*info.Details+"%")
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("Status = ?", *info.Status)
	}
	if info.StartSubmitTime != nil && info.EndSubmitTime != nil {
		db = db.Where("SubmitTime BETWEEN ? AND ? ", info.StartSubmitTime, info.EndSubmitTime)
	}
	if info.SubmitPerson != nil && *info.SubmitPerson != "" {
		db = db.Where("SubmitPerson = ?", *info.SubmitPerson)
	}
	if info.StartTransferTime != nil && info.EndTransferTime != nil {
		db = db.Where("TransferTime BETWEEN ? AND ? ", info.StartTransferTime, info.EndTransferTime)
	}
	if info.StartDeliveryTime != nil && info.EndDeliveryTime != nil {
		db = db.Where("DeliveryTime BETWEEN ? AND ? ", info.StartDeliveryTime, info.EndDeliveryTime)
	}
	if info.DeliveryID != nil && *info.DeliveryID != "" {
		db = db.Where("DeliveryID LIKE ?", "%"+*info.DeliveryID+"%")
	}
	if info.StartCompleteTime != nil && info.EndCompleteTime != nil {
		db = db.Where("CompleteTime BETWEEN ? AND ? ", info.StartCompleteTime, info.EndCompleteTime)
	}
	if info.CompleteID != nil && *info.CompleteID != "" {
		db = db.Where("CompleteID LIKE ?", "%"+*info.CompleteID+"%")
	}
	if info.StartUpdateTime != nil && info.EndUpdateTime != nil {
		db = db.Where("UpdateTime BETWEEN ? AND ? ", info.StartUpdateTime, info.EndUpdateTime)
	}
	if info.Notes != nil {
		db = db.Where("notes LIKE ?", "%"+*info.Notes+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["BankName"] = true
	orderMap["ID"] = true
	orderMap["BusinessContact"] = true
	orderMap["Status"] = true
	orderMap["SubmitTime"] = true
	orderMap["SubmitPerson"] = true
	orderMap["TransferTime"] = true
	orderMap["DeliveryTime"] = true
	orderMap["DeliveryID"] = true
	orderMap["CompleteTime"] = true
	orderMap["CompleteID"] = true
	orderMap["UpdateTime"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}
	db.Order("SubmitTime desc,Status desc,id desc")
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&bankWorkOrderss).Error
	return bankWorkOrderss, total, err
}
func (bankWorkOrdersService *BankWorkOrdersService) GetBankWorkOrdersPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
