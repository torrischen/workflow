package flow

import (
	"context"
	"time"
)

type Flow struct {
	ctx   context.Context
	delay time.Duration
}

func NewFlow(mysqlOpt *MysqlOption, opts ...FlowOption) *Flow {
	f := &Flow{
		ctx: context.Background(),
	}
	for _, opt := range opts {
		opt(f)
	}

	initMysql(
		mysqlOpt.User,
		mysqlOpt.Password,
		mysqlOpt.Host,
		mysqlOpt.Port,
		mysqlOpt.DbName,
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

func (f *Flow) UpdatePipeline(id string, data map[string]interface{}) error {
	return updatePipeline(id, data)
}
