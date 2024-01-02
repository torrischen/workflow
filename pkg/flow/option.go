package flow

import (
	"context"
	"time"
)

type FlowOption func(*Flow)

type MysqlOption struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type RedisOption struct {
	Host     string
	Port     string
	Password string
	Db       int
}

type StorageOption struct {
	Mysql *MysqlOption
	Redis *RedisOption
}

func WithContext(ctx context.Context) FlowOption {
	return func(f *Flow) {
		f.ctx = ctx
	}
}

func WithDelay(d time.Duration) FlowOption {
	return func(f *Flow) {
		f.delay = d
	}
}
