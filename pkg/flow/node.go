package flow

type Node struct {
	Base
	Name       string     `json:"name"`
	PipelineID string     `json:"pipeline_id"`
	PrevNodeID string     `json:"prev_node_id"`
	NextNodeID string     `json:"next_node_id"`
	Template   string     `json:"template"`
	Pipeline   *Pipeline  `json:"pipeline"`
	NodeData   []NodeData `json:"node_data"`
}
