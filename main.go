package main

import (
	"github.com/Rishikesh01/amazon-clone-backend/config"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	engine := new(config.Engine)
	engine.Run()
}
