package errs

import (
	"bytes"
	"fmt"
	"os"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	InternalError    = Error{ErrCode: 1003, Msg: "Internal services error"}
	UnKnownError     = Error{ErrCode: 1004, Msg: "unknown error"}
	Default          = Error{ErrCode: 0, Msg: "success"}
	ParameterError   = Error{ErrCode: 1001, Msg: "Parameter error: %s "}
	MongoInsertError = Error{ErrCode: 1101, Msg: "Value insert error: %s "}
	MongoUpdateError = Error{ErrCode: 1102, Msg: "Value update error: %s "}
	MongoNotFound    = Error{ErrCode: 1103, Msg: "Value not found: %s "}
	MongoQueryError  = Error{ErrCode: 1104, Msg: "Query error: %s"}
	MongoCreateError = Error{ErrCode: 1105, Msg: "Create error: %s"}
	MongoDupError    = Error{ErrCode: 1106, Msg: "Dup error: %s"}
	FileReadError    = Error{ErrCode: 1201, Msg: "Dup error: %s"}

	HttpFailed = Error{ErrCode: 1201, Msg: "Http connection abnormal: %s"}
)

type Error struct {
	ErrCode   int    `json:"err_code"`
	Msg       string `json:"msg"`
	DetailMsg string `json:"detail_msg"`
}

func (e Error) Error() string {
	return e.String()
}

func (e Error) String() string {
	return fmt.Sprintf("errCode: %d, msg: %s, detail_msg: %s", e.ErrCode, e.Msg, e.DetailMsg)
}

func (e Error) AddMsgf(msgs ...interface{}) Error {
	e.Msg = fmt.Sprintf(e.Msg, msgs...)
	return e
}

func (e Error) AddDetailMsg(msgs ...error) Error {
	var buf bytes.Buffer
	buf.WriteString(" Detail: ")
	for _, item := range msgs {
		if item == nil {
			continue
		}
		buf.WriteString(item.Error())
	}
	e.DetailMsg = buf.String()
	return e
}

func (e Error) Log(funcName string) *Error {

	serverEnv := os.Getenv("SERVER_ENV")
	if serverEnv != "" && serverEnv == "unittest" {
		return &e
	}
	logrus.WithFields(logrus.Fields{"module": funcName, "error_code": e.ErrCode}).Errorf("%+v\n", errors.WithStack(e))
	return &e
}
