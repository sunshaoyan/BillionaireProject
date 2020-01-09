package response

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v8"
	"hackathon/base/errs"
	"net/http"
	"strings"
)

type Result struct {
	Message   string      `json:"message"`
	Code      int         `json:"code"`
	RequestId string      `json:"request_id"`
	Data      interface{} `json:"data"`
}

// Response 封装了约定的返回结果
// 正常data返回  Response(c, data)
// 带pagination Response(c, data, pagination), Response(c, pagination, data) 顺序无关
// error返回，需要是 errs.Error类型，Response(c, errs.TestError{})
func Response(c *gin.Context, datalist ...interface{}) {
	ResponseWithCode(c, http.StatusOK, datalist...)
}

func ResponseWithCode(c *gin.Context, code int, datalist ...interface{}) {
	requestIDStr := uuid.NewV4().String()
	err := errs.Default
	var data interface{}
	for _, item := range datalist {
		if tmp, ok := item.(errs.Error); ok {
			err = tmp
			continue
		}
		if tmp, ok := item.(*errs.Error); ok {
			err = *tmp
			continue
		}
		if ve, ok := item.(validator.ValidationErrors); ok {
			buff := bytes.NewBufferString("")
			for key := range ve {
				key = getJsonKey(key)
				buff.WriteString(fmt.Sprintf(" field '%s' validate error", key))
			}
			err = *errs.ParameterError.AddMsgf(strings.TrimSpace(buff.String())).Log("reqId: " + requestIDStr)
			continue
		}
		if tmp, ok := item.(error); ok {
			err = *errs.UnKnownError.AddDetailMsg(tmp).Log("reqId: " + requestIDStr)
			continue
		}
		data = item
	}

	res := Result{
		Data:      data,
		RequestId: requestIDStr,
		Message:   err.Msg, Code: err.ErrCode}
	c.JSON(code, res)
}

func getJsonKey(key string) string {
	split := strings.Split(key, ".")
	str := split[1]
	data := make([]byte, 0, len(str)*2)
	j := false
	num := len(str)
	for i := 0; i < num; i++ {
		d := str[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(strings.TrimSpace(string(data[:])))
}
