package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hackathon/base/errs"
	"hackathon/base/response"
)

var Setting string

type ConfigRequst struct {
	Data string `json:"data"`
}

func GetConfig(c *gin.Context) {
	response.Response(c, Setting)
	Setting = ""
}

func SetConfig(c *gin.Context) {
	request := ConfigRequst{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.Response(c, errs.ParameterError.AddMsgf(err.Error()))
		return
	}
	Setting = request.Data
	fmt.Println(request)
	response.Response(c)
}
