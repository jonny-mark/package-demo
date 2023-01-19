package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jonny-mark/package-demo/rabbitMq/app"
	logger "github.com/jonny-mark/package-demo/rabbitMq/pkg/log"
)

func main() {
	logger.Init()
	router := gin.New()
	router.POST("/publish", app.Publish)
	router.Run()
}
