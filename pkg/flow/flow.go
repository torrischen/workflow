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
