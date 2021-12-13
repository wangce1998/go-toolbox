package xjob

import (
	"github.com/wangce1998/go-toolbox/xcontext"
)

type JobFunc func(ctx xcontext.XContext) error

type XJob interface {
	Name() string
	Describe() string
	Run(ctx xcontext.XContext) error
}

type defaultXJob struct {
	name     string
	describe string
	DoFunc   JobFunc
}

func (d defaultXJob) Run(ctx xcontext.XContext) error {
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
