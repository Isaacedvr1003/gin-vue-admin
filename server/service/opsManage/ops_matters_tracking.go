package opsManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/opsManage"
	opsManageReq "github.com/flipped-aurora/gin-vue-admin/server/model/opsManage/request"
)

type OpsMattersTrackingService struct{}

// CreateOpsMattersTracking 创建事项管理记录
// Author [yourname](https://github.com/yourname)
func (opsMattersTrackingService *OpsMattersTrackingService) CreateOpsMattersTracking(opsMattersTracking *opsManage.OpsMattersTracking) (err error) {
	err = global.GVA_DB.Create(opsMattersTracking).Error
	return err
}

// DeleteOpsMattersTracking 删除事项管理记录
// Author [yourname](https://github.com/yourname)
func (opsMattersTrackingService *OpsMattersTrackingService) DeleteOpsMattersTracking(id string) (err error) {
	err = global.GVA_DB.Delete(&opsManage.OpsMattersTracking{}, "id = ?", id).Error
	return err
}

// DeleteOpsMattersTrackingByIds 批量删除事项管理记录
// Author [yourname](https://github.com/yourname)
func (opsMattersTrackingService *OpsMattersTrackingService) DeleteOpsMattersTrackingByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]opsManage.OpsMattersTracking{}, "id in ?", ids).Error
	return err
}

// UpdateOpsMattersTracking 更新事项管理记录
// Author [yourname](https://github.com/yourname)
func (opsMattersTrackingService *OpsMattersTrackingService) UpdateOpsMattersTracking(opsMattersTracking opsManage.OpsMattersTracking) (err error) {
	err = global.GVA_DB.Model(&opsManage.OpsMattersTracking{}).Where("id = ?", opsMattersTracking.Id).Updates(&opsMattersTracking).Error
	return err
}

// GetOpsMattersTracking 根据id获取事项管理记录
// Author [yourname](https://github.com/yourname)
func (opsMattersTrackingService *OpsMattersTrackingService) GetOpsMattersTracking(id string) (opsMattersTracking opsManage.OpsMattersTracking, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&opsMattersTracking).Error
	return
}

// GetOpsMattersTrackingInfoList 分页获取事项管理记录
// Author [yourname](https://github.com/yourname)
func (opsMattersTrackingService *OpsMattersTrackingService) GetOpsMattersTrackingInfoList(info opsManageReq.OpsMattersTrackingSearch) (list []opsManage.OpsMattersTracking, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&opsManage.OpsMattersTracking{})
	var opsMattersTrackings []opsManage.OpsMattersTracking
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Id != nil {
		db = db.Where("id = ?", *info.Id)
	}
	if info.Description != nil && *info.Description != "" {
		db = db.Where("description LIKE ?", "%"+*info.Description+"%")
	}
	if info.Priority != nil && *info.Priority != "" {
		db = db.Where("priority = ?", *info.Priority)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	if info.AssignedTo != nil && *info.AssignedTo != "" {
		db = db.Where("assigned_to = ?", *info.AssignedTo)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["description"] = true
	orderMap["priority"] = true
	orderMap["status"] = true
	orderMap["assigned_to"] = true
	orderMap["created_at"] = true
	orderMap["updated_at"] = true
	orderMap["deleted_at"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}
	db.Order("status desc,priority desc,id desc")
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&opsMattersTrackings).Error
	return opsMattersTrackings, total, err
}
func (opsMattersTrackingService *OpsMattersTrackingService) GetOpsMattersTrackingPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
