package flow

type PipelineArgs struct {
	Topic  string
	Remark string
	Nodes  []NodeArgs
}

type NodeArgs struct {
	Name       string
	PipelineID string `json:"pipeline_id"`
	Template   string `json:"template"`
}

type PipelineRunArgs struct {
	PipelineID string
}
