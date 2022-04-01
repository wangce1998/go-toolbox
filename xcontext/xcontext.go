package xcontext

import (
	"context"
	beego "github.com/beego/beego/v2/server/web/context"
	uuid "github.com/satori/go.uuid"
)

const (
	ContextKey = "xcontext"
)

type XContext interface {
	context.Context
	RequestID() string
	SetRequestID(id string)
}

type defaultXContext struct {
	context.Context
	id string
}

func (d *defaultXContext) RequestID() string {
	return d.id
}

func (d *defaultXContext) SetRequestID(id string) {
	d.id = id
}

func New() XContext {
	return &defaultXContext{
		Context: context.TODO(),
		id:      uuid.NewV4().String(),
	}
}

func Wrap(ctx *beego.Context) XContext {
	xc, ok := ctx.Input.GetData(ContextKey).(XContext)
	if !ok {
		return New()
	}

	return &defaultXContext{
		Context: xc,
		id:      xc.RequestID(),
	}
}
