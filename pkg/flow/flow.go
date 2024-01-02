package flow

import (
	"context"
	"errors"
	"time"
)

type Flow struct {
	ctx   context.Context
	delay time.Duration
}

func NewFlow(storage *StorageOption, opts ...FlowOption) *Flow {
	f := &Flow{
		ctx: context.Background(),
	}
	for _, opt := range opts {
		opt(f)
	}

	initMysql(
		storage.Mysql.User,
		storage.Mysql.Password,
		storage.Mysql.Host,
		storage.Mysql.Port,
		storage.Mysql.DbName,
	)

	return f
}

func (f *Flow) NewPipeline(args *PipelineArgs) error {
	pipe := &Pipeline{
		Topic:  args.Topic,
		Remark: args.Remark,
	}
	if err := createPipeline(pipe); err != nil {
		return err
	}

	seq := 0
	var prev string

	for _, nodeArgs := range args.Nodes {
		seq++
		node := &Node{
			Name:       nodeArgs.Name,
			Sequence:   seq,
			PipelineID: pipe.ID,
			PrevNodeID: prev,
			Template:   nodeArgs.Template,
		}
		if err := createNode(node); err != nil {
			return err
		}

		if seq != 1 {
			if err := updateNode(prev, map[string]interface{}{"NextNodeID": node.ID}); err != nil {
				return err
			}
		}

		prev = node.ID
	}

	return nil
}

func (f *Flow) GetPipelineByID(id string, needNode, needPipelineRun bool) (*Pipeline, error) {
	return getPipelineByID(id, needNode, needPipelineRun)
}

func (f *Flow) UpdatePipelineTopic(id string, topic string) error {
	return updatePipeline(id, map[string]interface{}{
		"topic": topic,
	})
}

func (f *Flow) UpdatePipelineRemark(id, remark string) error {
	return updatePipeline(id, map[string]interface{}{
		"remark": remark,
	})
}

// DeletePipeline will delete the pipeline and all nodes in it.
func (f *Flow) DeletePipeline(id string) error {
	pipe, err := getPipelineByID(id, true, true)
	if err != nil {
		return err
	}

	if len(pipe.PipelineRun) > 0 {
		return errors.New("pipeline has been run, cannot delete")
	}

	for _, node := range pipe.Node {
		if err := deleteNode(node.ID); err != nil {
			return err
		}
	}

	if err := deletePipeline(id); err != nil {
		return err
	}

	return nil
}

func (f *Flow) NewPipelineRun(args *PipelineRunArgs) error {
	pipe, err := getPipelineByID(args.PipelineID, true, false)
	if err != nil {
		return err
	}

	if len(pipe.Node) == 0 {
		return errors.New("pipeline has no node")
	}

	pipeRun := &PipelineRun{
		PipelineID: pipe.ID,
		Stage:      pipe.Node[0].ID,
		Status:     PipelineRunStatusProcessing,
	}
	if err := createPipelineRun(pipeRun); err != nil {
		return err
	}

	for _, node := range pipe.Node {
		nodeRun := &NodeData{
			PipelineRunID: pipeRun.ID,
			NodeID:        node.ID,
		}
		if err := createNodeData(nodeRun); err != nil {
			return err
		}
	}

	return nil
}
