package xjob

import (
	"github.com/wangce1998/go-toolbox/xcontext"
)

type XJob interface {
	Name() string
	Describe() string
	Run(ctx xcontext.XContext) error
}

type defaultXJob struct {
	name     string
	describe string
}

func (d defaultXJob) Name() string {
	return d.name
}

func (d defaultXJob) Describe() string {
	return d.describe
}

func (d defaultXJob) Run(ctx xcontext.XContext) error {
	panic("待实现的job")
}

func New(name string, describe string) XJob {
	return &defaultXJob{
		name:     name,
		describe: describe,
	}
}
