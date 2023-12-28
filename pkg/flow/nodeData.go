package flow

type NodeData struct {
	Base
	PipelineRunID string       `gorm:"index:idx_run_node" json:"pipeline_run_id"`
	NodeID        string       `gorm:"index:idx_run_node" json:"node_id"`
	Data          string       `json:"data"`
	Node          *Node        `json:"node"`
	PipelineRun   *PipelineRun `json:"pipeline_run"`
}
