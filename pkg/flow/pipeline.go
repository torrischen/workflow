package flow

import "gorm.io/gorm"

type Pipeline struct {
	Base
	Topic       string        `json:"topic"`
	Head        string        `json:"head"` // head node ID
	Remark      string        `json:"remark"`
	Node        []Node        `json:"node"`
	PipelineRun []PipelineRun `json:"pipeline_run"`
}

func createPipeline(pipeline *Pipeline) error {
	if err := db.Create(pipeline).Error; err != nil {
		return err
	}

	return nil
}

func updatePipeline(id string, data map[string]interface{}) error {
	if err := db.Model(&Pipeline{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func deletePipeline(id string) error {
	if err := db.Delete(&Pipeline{}, id).Error; err != nil {
		return err
	}

	return nil
}

func getPipelineByID(id string, needNode, needPipelineRun bool) (*Pipeline, error) {
	var pipeline Pipeline

	_db := db
	if needNode {
		_db = _db.Preload("Node", func(db *gorm.DB) *gorm.DB {
			return db.Order("sequence asc")
		})
	}
	if needPipelineRun {
		_db = _db.Preload("PipelineRun")
	}

	if err := _db.
		Where("id=?", id).
		First(&pipeline).Error; err != nil {
		return nil, err
	}

	return &pipeline, nil
}
