package xcontext

import (
	"context"
	"github.com/satori/go.uuid"
)

const XRequestIDKey = "X-Request-ID"

type XContext interface {
	context.Context
	RequestID() string
}

type defaultXContext struct {
	context.Context
	requestID string
}

func (c defaultXContext) RequestID() string {
	return c.requestID
}

func New() XContext {
	c, id := xRequestIDCtx(context.TODO())

	return &defaultXContext{
		Context:   c,
		requestID: id,
	}
}

func Wrap(c context.Context) XContext {
	if cc, ok := c.(XContext); ok {
		return cc
	}

	cc, id := xRequestIDCtx(c)

	return &defaultXContext{
		Context:   cc,
		requestID: id,
	}
}

func xRequestIDCtx(c context.Context) (context.Context, string) {
	if c == nil {
		panic("xRequestIDCtx：context.Context 不能传入 nil")
	}
	val := c.Value(XRequestIDKey)
	if id, ok := val.(string); ok && id != "" {
		return c, id
	}
	id := uuid.NewV4().String()

	return context.WithValue(c, XRequestIDKey, id), id
}
