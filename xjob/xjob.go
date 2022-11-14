package xjob

type JobFunc func() error

type XJob interface {
	Name() string
	Describe() string
	Run() error
}

type defaultXJob struct {
	name     string
	describe string
	DoFunc   JobFunc
}

func (d defaultXJob) Run() error {
	return d.DoFunc()
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
