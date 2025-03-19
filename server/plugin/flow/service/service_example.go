package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/flow/model"
	flowReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/flow/model/request"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

// 开始示例
func (flowService *FlowService) StartTransitionExample(flow model.TransitionExample, operation uint) (workUUID uuid.UUID, err error) {
	workUUID, _ = uuid.NewV4()
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var node model.Node
		var graph model.Graph
		err := tx.First(&graph, "flow_uuid = ?", flow.FlowUUID).Error
		if err != nil {
			return err
		}
		err = tx.First(&node, "graph_id = ? AND shape = ?", graph.ID, model.START_NODE).Error
		if err != nil {
			return err
		}
		var transition model.TransitionExample
		transition.FlowUUID = flow.FlowUUID
		transition.WorkUUID = workUUID
		transition.WorkType = flow.WorkType
		transition.Operator = operation
		err = tx.Create(&transition).Error
		if err != nil {
			return err
		}
		var step model.TransitionExampleStep
		step.FlowUUID = flow.FlowUUID
		step.WorkUUID = workUUID
		step.WorkType = flow.WorkType
		step.NodeID = node.ID
		step.Operator = operation
		step.Status = model.START_NODE
		err = tx.Create(&step).Error
		if err != nil {
			return err
		}
		var now flowReq.TransitionExampleNext
		now.WorkUUID = workUUID
		now.Attr = flow.Attr
		now.Logic = ""
		now.NodeID = node.ID
		return flowService.TransitionExampleNext(tx, now, operation)
	})
	return workUUID, err
}

// GetTransitionExample 根据ID获取流转示例记录
// Author [piexlmax](https://github.com/piexlmax)
func (flowService *FlowService) GetTransitionExampleStep(UUID string) (step []model.TransitionExampleStep, err error) {
	err = global.GVA_DB.Where("work_uuid = ?", UUID).Find(&step).Error
	return
}

// GetTransitionExampleInfoList 分页获取流转示例记录
// Author [piexlmax](https://github.com/piexlmax)
func (flowService *FlowService) GetTransitionExampleInfoList(info flowReq.TransitionExampleSearch) (list []model.TransitionExample, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.TransitionExample{})
	var transitionExamples []model.TransitionExample
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&transitionExamples).Error
	return transitionExamples, total, err
}

// GetTransitionExampleInfoList 分页获取流转示例记录
// Author [piexlmax](https://github.com/piexlmax)
func (flowService *FlowService) TransitionExampleNext(tx *gorm.DB, step flowReq.TransitionExampleNext, operator uint) (err error) {
	if tx == nil {
		tx = global.GVA_DB
	}
	nodes, err := flowService.Next(step.NodeID, step.Logic)
	var transition model.TransitionExample
	err = tx.Where("work_uuid = ?", step.WorkUUID).First(&transition).Error
	if err != nil {
		return err
	}
	if step.StepID != 0 {
		var s model.TransitionExampleStep
		err := tx.First(&s, "id = ?", step.StepID).Error
		if err != nil {
			return err
		}
		s.Operator = operator
		s.Attr = step.Attr
		s.Logic = step.Logic
		s.Status = "done"
		err = tx.Save(&s).Error
	}

	if nodes == nil || len(nodes) == 0 {
		return nil
	}

	for i := range nodes {
		if nodes[i].Node.Shape == model.MERGE_NODE {
			prev, err := flowService.Prev(nodes[i].NodeID)
			if err != nil {
				return err
			}
			var prevIDS []string
			for _, v := range prev {
				prevIDS = append(prevIDS, v.ID)
			}

			var merge []model.TransitionExampleStep
			ferr := tx.Where("node_id in (?) and work_uuid = ?", prevIDS, step.WorkUUID).Find(&merge).Error
			if ferr != nil {
				return ferr
			}
			var comLogic string
			for i := range merge {
				comLogic += merge[i].Logic
				if merge[i].Status == "active" {
					return nil
				}
			}

			// 拿到合并节点的所有来源节点，并且根据其逻辑值进行拼接传递向下个节点
			nextNodes, err := flowService.Next(nodes[0].NodeID, comLogic)
			if err != nil {
				return err
			}
			nodes[i].Next = nextNodes
		}
	}
	// 合并节点处理

	return tx.Transaction(func(tx *gorm.DB) error {
		return createNodes(tx, step, transition, operator, nodes)
	})
}

func createNodes(tx *gorm.DB, now flowReq.TransitionExampleNext, transition model.TransitionExample, operator uint, nodes []model.NextFlow) error {
	for i := range nodes {
		var step model.TransitionExampleStep
		step.FlowUUID = transition.FlowUUID
		step.WorkUUID = transition.WorkUUID
		step.WorkType = transition.WorkType
		step.NodeID = nodes[i].NodeID

		switch nodes[i].Node.Shape {
		case model.EVENT_NODE:
			step.Status = "active"
		default:
			step.Status = nodes[i].Node.Shape
		}
		err := tx.Create(&step).Error
		if err != nil {
			return err
		}
		if nodes[i].Next != nil && len(nodes[i].Next) > 0 {
			err = createNodes(tx, now, transition, operator, nodes[i].Next)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
