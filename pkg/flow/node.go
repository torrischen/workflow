package flow

type Node struct {
	Base
	Name       string     `json:"name"`
	Sequence   int        `gorm:"index:idx_seq_pipe" json:"sequence"`
	PipelineID string     `gorm:"index:idx_seq_pipe" json:"pipeline_id"`
	PrevNodeID string     `json:"prev_node_id"`
	NextNodeID string     `json:"next_node_id"`
	Template   string     `json:"template"`
	Pipeline   *Pipeline  `json:"pipeline"`
	NodeData   []NodeData `json:"node_data"`
}

func createNode(node *Node) error {
	if err := db.Create(node).Error; err != nil {
		return err
	}

	return nil
}

func updateNode(id string, data map[string]interface{}) error {
	if err := db.Model(&Node{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func deleteNode(id string) error {
	if err := db.Delete(&Node{}, id).Error; err != nil {
		return err
	}

	return nil
}

func getNodeByID(id string, needPipeline, needNodeData bool) (*Node, error) {
	var node Node

	_db := db
	if needPipeline {
		_db = _db.Preload("Pipeline")
	}
	if needNodeData {
		_db = _db.Preload("NodeData")
	}

	if err := _db.
		Where("id=?", id).
		First(&node).Error; err != nil {
		return nil, err
	}

	return &node, nil
}
