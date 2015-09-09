package mux

import (
	"golang.org/x/net/context"

	"github.com/julienschmidt/httprouter"
)

type key string

const muxKey key = "mux"

func NewContext(ctx context.Context, params httprouter.Params) context.Context {
	return context.WithValue(ctx, muxKey, params)
}

func FromContext(ctx context.Context) httprouter.Params {
	client, _ := ctx.Value(muxKey).(httprouter.Params)
	return client
}
