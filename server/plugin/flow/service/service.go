package service

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/flow/model"
	flowReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/flow/model/request"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type FlowService struct {
}

// CreateFlow 创建流程管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (flowService *FlowService) CreateFlow(flow *model.Flow) (err error) {
	flow.UUID, _ = uuid.NewV4()
	err = global.GVA_DB.Create(flow).Error
	return err
}

// DeleteFlow 删除流程管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (flowService *FlowService) DeleteFlow(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.Flow{}, "id = ?", ID).Error
	return err
}

// DeleteFlowByIds 批量删除流程管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (flowService *FlowService) DeleteFlowByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.Flow{}, "id in ?", IDs).Error
	return err
}

// UpdateFlow 更新流程管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (flowService *FlowService) UpdateFlow(flow model.Flow) (err error) {
	err = global.GVA_DB.Model(&model.Flow{}).Where("id = ?", flow.ID).Updates(&flow).Error
	return err
}

// GetFlow 根据ID获取流程管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (flowService *FlowService) GetFlow(ID string, uuid string) (flow model.Flow, err error) {
	db := global.GVA_DB.Model(&model.Flow{})

	if ID != "" {
		db = db.Where("id = ?", ID)
	}

	if uuid != "" {
		db = db.Or("uuid = ?", uuid)
	}

	err = db.First(&flow).Error
	return
}

func (flowService *FlowService) FindGraph(uuid string) (graph model.Graph, err error) {
	err = global.GVA_DB.Preload("Nodes").
		Preload("Nodes.Data").
		Preload("Edges").
		Preload("Edges.Data").
		Preload("Edges.Source").
		Preload("Edges.Target").
		First(&graph, "flow_uuid = ?", uuid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
		return
	}
	return
}

func (flowService *FlowService) SaveGraph(graph model.Graph) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var g model.Graph
		ferr := tx.Preload("Nodes").
			Preload("Nodes.Data").
			Preload("Edges").
			Preload("Edges.Data").
			Preload("Edges.Source").
			Preload("Edges.Target").First(&g, "flow_uuid = ?", graph.FlowUUID).Error
		if ferr != nil {
			tx = tx.Unscoped().Session(&gorm.Session{NewDB: true})
			txErr := tx.Create(&graph).Error
			return txErr
		}
		tx = tx.Unscoped().Session(&gorm.Session{NewDB: true})
		txErr := tx.Delete(&g).Error
		if txErr != nil {
			return txErr
		}

		var NodeIds []string
		for _, node := range g.Nodes {
			NodeIds = append(NodeIds, node.ID)
		}
		tx = tx.Session(&gorm.Session{NewDB: true})
		txErr = tx.Delete(&[]model.Node{}, "id in ?", NodeIds).Error
		if txErr != nil {
			return txErr
		}

		var EdgeIds []string
		var SourceIDS []string
		var TargetIDS []string
		for _, edge := range g.Edges {
			EdgeIds = append(EdgeIds, edge.ID)
			if edge.Source != nil {
				SourceIDS = append(SourceIDS, edge.Source.EdgeID)
			}
			if edge.Target != nil {
				TargetIDS = append(TargetIDS, edge.Target.EdgeID)
			}
		}

		tx = tx.Session(&gorm.Session{NewDB: true})
		txErr = tx.Delete(&[]model.Edge{}, "graph_id = ?", g.ID).Error
		if txErr != nil {
			return txErr
		}

		tx = tx.Session(&gorm.Session{NewDB: true})
		txErr = tx.Delete(&[]model.Source{}, "edge_id in ?", SourceIDS).Error
		if txErr != nil {
			return txErr
		}

		tx = tx.Session(&gorm.Session{NewDB: true})
		txErr = tx.Delete(&[]model.Target{}, "edge_id in ?", TargetIDS).Error
		if txErr != nil {
			return txErr
		}

		tx = tx.Session(&gorm.Session{NewDB: true})
		txErr = tx.Delete(&[]model.EdgeData{}, "edge_id in ?", EdgeIds).Error
		if txErr != nil {
			return txErr
		}

		tx = tx.Session(&gorm.Session{NewDB: true})
		txErr = tx.Delete(&[]model.NodeData{}, "node_id in ?", NodeIds).Error
		if txErr != nil {
			return txErr
		}

		return tx.Create(&graph).Error
	})
}

// GetFlowInfoList 分页获取流程管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (flowService *FlowService) GetFlowInfoList(info flowReq.FlowSearch) (list []model.Flow, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Flow{})
	var flows []model.Flow
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.UUID != "" {
		db = db.Where("uuid = ?", info.UUID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&flows).Error
	return flows, total, err
}

// 获取当前节点的前面的所有节点
func (flowService *FlowService) Prev(nodeID string) (prev []model.Node, err error) {
	var targets []model.Target
	err = global.GVA_DB.Find(&targets, "cell = ?", nodeID).Error
	if err != nil {
		return
	}
	var edgeIDS []string
	for _, target := range targets {
		edgeIDS = append(edgeIDS, target.EdgeID)
	}
	var sources []model.Source
	err = global.GVA_DB.Find(&sources, "edge_id in ?", edgeIDS).Error
	if err != nil {
		return
	}
	var cellIDS []string
	for _, source := range sources {
		cellIDS = append(cellIDS, source.Cell)
	}
	var nodes []model.Node
	err = global.GVA_DB.Find(&nodes, "id in ?", cellIDS).Error
	if err != nil {
		return
	}
	for _, node := range nodes {
		prev = append(prev, node)
	}
	return
}

func (flowService *FlowService) Next(nodeID string, logic string) (next []model.NextFlow, err error) {
	var source []model.Source
	err = global.GVA_DB.Find(&source, "cell = ?", nodeID).Error
	var nowNode model.Node
	err = global.GVA_DB.First(&nowNode, "id = ?", nodeID).Error
	if err != nil {
		return
	}

	switch nowNode.Shape {
	case model.PARALLEL_NODE, model.MERGE_NODE:
		// 当前节点是并行节点的话 自动做什么事情
		// 平行节点会根据当前source节点的edgeID去找到下面所有的节点并且返回
		for _, s := range source {
			var target model.Target
			err = global.GVA_DB.First(&target, "edge_id = ?", s.EdgeID).Error
			if err != nil {
				return
			}
			var n model.NextFlow
			n.EdgeID = s.EdgeID
			n.NodeID = target.Cell
			var edge model.Edge
			err = global.GVA_DB.First(&edge, "id = ?", s.EdgeID).Error
			if err != nil {
				return
			}
			n.Edge = edge
			var node model.Node
			err = global.GVA_DB.First(&node, "id = ?", target.Cell).Error
			if err != nil {
				return
			}
			n.Node = node
			next = append(next, n)
		}
	case model.DECISION_NODE:
		// 当前节点是决策节点的话 自动做什么事情
		var edgeIDS []string
		for _, s := range source {
			edgeIDS = append(edgeIDS, s.EdgeID)
		}
		var edges []model.Edge
		err = global.GVA_DB.Preload("Data").Find(&edges, "id in ?", edgeIDS).Error
		var checkEdge model.Edge
		for _, edge := range edges {
			if edge.Data.Logic == logic {
				checkEdge = edge
				break
			}
		}
		var target model.Target
		err = global.GVA_DB.First(&target, "edge_id = ?", checkEdge.ID).Error
		if err != nil {
			return
		}
		var n model.NextFlow
		n.EdgeID = checkEdge.ID
		n.NodeID = target.Cell

		var edge model.Edge
		err = global.GVA_DB.First(&edge, "id = ?", checkEdge.ID).Error
		if err != nil {
			return
		}
		n.Edge = edge
		var node model.Node
		err = global.GVA_DB.First(&node, "id = ?", target.Cell).Error
		if err != nil {
			return
		}
		n.Node = node
		next = append(next, n)
	default:
		// 普通节点对source判断 如果多于一个返回普通节点不能有多个出口
		if len(source) > 1 || len(source) == 0 {
			err = errors.New("普通节点出口数量有误")
			return
		}

		var target model.Target
		err = global.GVA_DB.First(&target, "edge_id = ?", source[0].EdgeID).Error
		if err != nil {
			return
		}

		var n model.NextFlow
		n.EdgeID = source[0].EdgeID
		n.NodeID = target.Cell

		var edge model.Edge
		err = global.GVA_DB.First(&edge, "id = ?", source[0].EdgeID).Error
		if err != nil {
			return
		}
		n.Edge = edge
		var node model.Node
		err = global.GVA_DB.First(&node, "id = ?", target.Cell).Error
		if err != nil {
			return
		}
		n.Node = node
		next = append(next, n)
	}

	//下一个节点

	for i := range next {
		var targetNode model.Node
		err = global.GVA_DB.First(&targetNode, "id = ?", next[i].NodeID).Error
		switch targetNode.Shape {
		case model.EVENT_NODE, model.MERGE_NODE, model.END_NODE, model.START_NODE:
			// 事件 开始 结束 合并 四个节点都是直接返回当前节点
			return

		case model.DECISION_NODE, model.PARALLEL_NODE:
			// 返回当前节点和下面的多个节点
			n, err := flowService.Next(next[i].NodeID, logic)
			if err != nil {
				return nil, err
			}
			next[i].Next = append(next[i].Next, n...)
			return next, nil
			// 检查上方节点是否已经完全完成，如果完全完成则返回当前节点（login的拼接）和下一节点
		}
	}
	return
}
