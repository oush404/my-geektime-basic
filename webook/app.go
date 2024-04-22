package main

import (
	"github.com/gin-gonic/gin"
	"my-geektime-basic/webook/internal/events"
)

type App struct {
	server    *gin.Engine
	consumers []events.Consumer
}
