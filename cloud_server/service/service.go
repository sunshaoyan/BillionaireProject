package service

import (
	"github.com/gin-gonic/gin"
	"hackathon/service/controller"
)

func NewService(base *gin.RouterGroup) {
	routerRegister(base)
}

func routerRegister(base *gin.RouterGroup) {
	ds := base.Group("/hackathon")
	ds.GET("/config", controller.GetConfig)
	ds.POST("/set", controller.SetConfig)
}
