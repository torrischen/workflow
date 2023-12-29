package flow

type Pipeline struct {
	Base
	Topic       string        `json:"topic"`
	Head        string        `json:"head"` // head node ID
	Remark      string        `json:"remark"`
	Node        []Node        `json:"node"`
	PipelineRun []PipelineRun `json:"pipeline_run"`
}

func CreatePipeline(pipeline *Pipeline) error {
	if err := db.Create(pipeline).Error; err != nil {
		return err
	}

	return nil
}

func UpdatePipeline(id string, data map[string]interface{}) error {
	if err := db.Model(&Pipeline{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func DeletePipeline(id string) error {
	if err := db.Delete(&Pipeline{}, id).Error; err != nil {
		return err
	}

	return nil
}

func GetPipelineByID(id string) (*Pipeline, error) {
	var pipeline Pipeline
	if err := db.Preload("Node").Preload("PipelineRun").First(&pipeline, id).Error; err != nil {
		return nil, err
	}

	return &pipeline, nil
}
