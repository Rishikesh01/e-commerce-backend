package config

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Engine struct{}

func (e *Engine) Run() {
	router := gin.Default()
	start(router)
}

func start(router *gin.Engine) {
	//util := &util.BcryptUtil{}
	//db := repository.Init()

	err := router.Run()
	if err != nil {
		log.Println(err)
	}
}
