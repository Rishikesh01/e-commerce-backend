package main

import (
	"github.com/Rishikesh01/amazon-clone-backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := new(config.Engine)
	engine.Run()
}
