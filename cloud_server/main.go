package main

import (
	"github.com/gin-gonic/gin"
	// "hackathon/base/db"
	"hackathon/conf"
	"hackathon/service"
)

func main() {
	conf.ConfInit()
	// db.MongoConnect()
	r := gin.New()
	r.Use(gin.Recovery())

	service.NewService(&r.RouterGroup)

	r.Run("0.0.0.0:" + conf.Configure.Port)
}
