package xjob

import (
	"context"
)

type JobFunc func(ctx context.Context) error

type XJob interface {
	Name() string
	Describe() string
	Run(ctx context.Context) error
}

type defaultXJob struct {
	name     string
	describe string
	DoFunc   JobFunc
}

func (d defaultXJob) Run(ctx context.Context) error {
	// TODO::锁
	return d.DoFunc(ctx)
}

func (d defaultXJob) Name() string {
	return d.name
}

func (d defaultXJob) Describe() string {
	return d.describe
}

func New(name string, describe string, do JobFunc) XJob {
	return &defaultXJob{
		name:     name,
		describe: describe,
		DoFunc:   do,
	}
}
