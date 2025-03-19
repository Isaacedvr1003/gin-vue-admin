package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

const (
	START_NODE    = "start-node"    // 开始节点
	END_NODE      = "end-node"      // 结束节点
	EVENT_NODE    = "event-node"    // 事件节点
	DECISION_NODE = "decision-node" // 决策节点
	PARALLEL_NODE = "parallel-node" // 并行节点
	MERGE_NODE    = "merge-node"    // 合并节点
	//AUTO_NODE     = "auto-node"     // 自动节点
)

// Graph 表示流程图的结构，包括流程UUID，节点和边
type Graph struct {
	global.GVA_MODEL
	FlowUUID string `json:"flowUUID" form:"flowUUID" gorm:"column:flow_uuid;comment:流程uuid;"` //流程uuid
	Nodes    []Node `json:"nodes" form:"nodes" binding:"required"`                            //节点
	Edges    []Edge `json:"edges" form:"edges" binding:"required"`                            //边
}

// TableName 设置Graph的表名为"gva_flow_graph"
func (Graph) TableName() string {
	return "gva_flow_graph"
}

// NodeData 表示节点数据的结构，包括节点ID和名称
type NodeData struct {
	PID         uint   `gorm:"primarykey" json:"PID"` // 主键ID
	NodeID      string `json:"nodeId" gorm:"column:node_id;comment:节点ID"`
	Name        string `json:"name" gorm:"column:name;comment:名称"`
	Title       string `json:"title" gorm:"column:title;comment:标题"`
	Description string `json:"description" gorm:"column:description;comment:描述"`
	Logic       string `json:"logic" gorm:"column:logic;comment:逻辑"`
}

// TableName 设置NodeData的表名为"gva_flow_graph_node_data"
func (NodeData) TableName() string {
	return "gva_flow_graph_node_data"
}

// Node 表示节点的结构，包括ID，流程UUID，位置，大小，属性，可见性，形状，端口，数据和ZIndex
type Node struct {
	PID      uint           `gorm:"primarykey" json:"PID"` // 主键ID
	ID       string         `json:"id" gorm:"column:id;comment:ID;"`
	GraphID  uint           `json:"graphId" gorm:"column:graph_id;comment:流程ID"`
	Position datatypes.JSON `json:"position" gorm:"column:position;comment:位置;"`
	Size     datatypes.JSON `json:"size" gorm:"column:size;comment:大小;"`
	Attrs    datatypes.JSON `json:"attrs" gorm:"column:attrs;comment:属性;"`
	Visible  bool           `json:"visible" gorm:"column:visible;comment:可见性;type:bool;default:true"`
	Shape    string         `json:"shape" gorm:"column:shape;comment:形状"`
	Ports    datatypes.JSON `json:"ports" gorm:"column:ports;comment:端口;type:json"`
	Data     NodeData       `json:"data" gorm:"foreignKey:NodeID;references:ID;"`
	ZIndex   int            `json:"zIndex" gorm:"column:z_index;comment:Z索引;type:int;default:0"`
}

// TableName 设置Node的表名为"gva_flow_graph_node"
func (Node) TableName() string {
	return "gva_flow_graph_node"
}

// EdgeEndpoint 表示边的端点的结构，包括边ID，单元格和端口
type EdgeEndpoint struct {
	PID    uint   `gorm:"primarykey" json:"PID"` // 主键ID
	EdgeID string `json:"edgeId" gorm:"column:edge_id;comment:边ID"`
	Cell   string `json:"cell" gorm:"column:cell;comment:单元格"`
	Port   string `json:"port" gorm:"column:port;comment:端口"`
}

type Source struct {
	EdgeEndpoint
}

type Target struct {
	EdgeEndpoint
}

// TableName 设置EdgeEndpoint的表名为"gva_flow_graph_edge_endpoint"
func (Source) TableName() string {
	return "gva_flow_graph_edge_source"
}

func (Target) TableName() string {
	return "gva_flow_graph_edge_target"
}

type EdgeData struct {
	PID         uint   `gorm:"primarykey" json:"PID"` // 主键ID
	EdgeID      string `json:"edgeId" gorm:"column:edge_id;comment:节点ID;"`
	Name        string `json:"name" gorm:"column:name;comment:名称"`
	Title       string `json:"title" gorm:"column:title;comment:标题"`
	Description string `json:"description" gorm:"column:description;comment:描述"`
	Logic       string `json:"logic" gorm:"column:logic;comment:逻辑"`
}

// TableName 设置NodeData的表名为"gva_flow_graph_node_data"
func (EdgeData) TableName() string {
	return "gva_flow_graph_edge_data"
}

// Edge 表示边的结构，包括流程UUID，形状，属性，ID，ZIndex，源和目标
type Edge struct {
	PID     uint           `gorm:"primarykey" json:"PID"` // 主键ID
	ID      string         `json:"id" gorm:"column:id;comment:ID"`
	GraphID uint           `json:"graphId" gorm:"column:graph_id;comment:流程ID"`
	Shape   string         `json:"shape" gorm:"column:shape;comment:形状"`
	Attrs   datatypes.JSON `json:"attrs" gorm:"column:attrs;comment:属性;type:json"`
	ZIndex  int            `json:"zIndex" gorm:"column:z_index;comment:Z索引;type:int;default:0"`
	Data    EdgeData       `json:"data" gorm:"foreignKey:EdgeID;references:ID;"`
	Source  *Source        `json:"source" gorm:"foreignKey:EdgeID;references:ID;"`
	Target  *Target        `json:"target" gorm:"foreignKey:EdgeID;references:ID;"`
}

// TableName 设置Edge的表名为"gva_flow_graph_edge"
func (Edge) TableName() string {
	return "gva_flow_graph_edge"
}
