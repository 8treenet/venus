package main

import (
	quickgin "github.com/8treenet/venus/quick-gin"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := quickgin.New()
	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	//engine.Run()
	engine.RunH2C()
}
