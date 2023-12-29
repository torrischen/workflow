package flow

type NodeData struct {
	Base
	PipelineRunID string       `gorm:"index:idx_run_node" json:"pipeline_run_id"`
	NodeID        string       `gorm:"index:idx_run_node" json:"node_id"`
	Data          string       `json:"data"`
	Node          *Node        `json:"node"`
	PipelineRun   *PipelineRun `json:"pipeline_run"`
}

func CreateNodeData(nodeData *NodeData) error {
	if err := db.Create(nodeData).Error; err != nil {
		return err
	}

	return nil
}

func UpdateNodeData(id string, data map[string]interface{}) error {
	if err := db.Model(&NodeData{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func DeleteNodeData(id string) error {
	if err := db.Delete(&NodeData{}, id).Error; err != nil {
		return err
	}

	return nil
}

func GetNodeDataByID(id string) (*NodeData, error) {
	var nodeData NodeData
	if err := db.Preload("PipelineRun").Preload("Node").First(&nodeData, id).Error; err != nil {
		return nil, err
	}

	return &nodeData, nil
}
