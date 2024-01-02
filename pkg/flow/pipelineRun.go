package flow

const (
	PipelineRunStatusProcessing = "processing"
	PipelineRunStatusPending    = "pending"
	PipelineRunStatusSuccess    = "success"
	PipelineRunStatusFailed     = "failed"
)

type PipelineRun struct {
	Base
	PipelineID string     `gorm:"index:idx_pipe_stage" json:"pipeline_id"`
	Stage      string     `gorm:"idnex:idx_pipe_stage" json:"stage"` //node ID
	Status     string     `json:"status"`                            // processing, pending, success, failed
	Pipeline   *Pipeline  `json:"pipeline"`
	NodeData   []NodeData `json:"node_data"`
}

func createPipelineRun(pipelineRun *PipelineRun) error {
	if err := db.Create(pipelineRun).Error; err != nil {
		return err
	}

	return nil
}

func updatePipelineRun(id string, data map[string]interface{}) error {
	if err := db.Model(&PipelineRun{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func deletePipelineRun(id string) error {
	if err := db.Delete(&PipelineRun{}, id).Error; err != nil {
		return err
	}

	return nil
}

func getPipelineRunByID(id string) (*PipelineRun, error) {
	var pipelineRun PipelineRun
	if err := db.
		Preload("Pipeline").
		Preload("NodeData").
		First(&pipelineRun, id).Error; err != nil {
		return nil, err
	}

	return &pipelineRun, nil
}
