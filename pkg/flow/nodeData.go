package flow

type NodeData struct {
	Base
	PipelineRunID string       `gorm:"index:idx_run_node" json:"pipeline_run_id"`
	NodeID        string       `gorm:"index:idx_run_node" json:"node_id"`
	Data          string       `json:"data"`
	Node          *Node        `json:"node"`
	PipelineRun   *PipelineRun `json:"pipeline_run"`
}

func createNodeData(nodeData *NodeData) error {
	if err := db.Create(nodeData).Error; err != nil {
		return err
	}

	return nil
}

func updateNodeData(id string, data map[string]interface{}) error {
	if err := db.Model(&NodeData{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func deleteNodeData(id string) error {
	if err := db.Delete(&NodeData{}, id).Error; err != nil {
		return err
	}

	return nil
}

func getNodeDataByID(id string, needNode, needPipelineRun bool) (*NodeData, error) {
	var nodeData NodeData

	_db := db
	if needNode {
		_db = _db.Preload("Node")
	}
	if needPipelineRun {
		_db = _db.Preload("PipelineRun")
	}

	if err := _db.
		Where("id=?", id).
		First(&nodeData).Error; err != nil {
		return nil, err
	}

	return &nodeData, nil
}
