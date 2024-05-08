package main

import (
	"my-geektime-basic/webook/internal/events"
	"my-geektime-basic/webook/pkg/ginx"
	"my-geektime-basic/webook/pkg/grpcx"
)

type App struct {
	consumers   []events.Consumer
	server      *grpcx.Server
	adminServer *ginx.Server
}
