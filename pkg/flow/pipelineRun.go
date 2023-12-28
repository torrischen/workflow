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
