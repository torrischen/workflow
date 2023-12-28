package flow

type Pipeline struct {
	Base
	Topic       string        `json:"topic"`
	Remark      string        `json:"remark"`
	Node        []Node        `json:"node"`
	PipelineRun []PipelineRun `json:"pipeline_run"`
}
