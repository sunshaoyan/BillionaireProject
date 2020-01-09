package logger

import (
	"bufio"
	"emq-auth/conf"
	"log"
	"os"
	"path"
	"time"

	"github.com/pkg/errors"

	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func LoggerInit() {
	var logFormat logrus.Formatter
	if conf.Configure.Logger.Format == "json" {
		logFormat = &logrus.JSONFormatter{}
	} else {
		logFormat = &logrus.TextFormatter{}
	}
	logrus.SetFormatter(logFormat)

	if _, err := os.Stat(conf.Configure.Logger.Logdir); os.IsNotExist(err) {
		err := os.MkdirAll(conf.Configure.Logger.Logdir, 0777)
		if err != nil {
			panic(err)
		}
	}

	baseLogPath := path.Join(conf.Configure.Logger.Logdir, conf.Configure.Logger.FileName)

	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),   // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(24*time.Hour),    // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{"module": "init_logger"}).Errorf("%+v\n", errors.WithStack(err))
	}

	switch conf.Configure.Logger.Level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetOutput(os.Stderr)
	case "info":
		// info 需要将log 通过stdout 输出到 elk，所以不setNull()
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		setNull()
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		setNull()
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		setNull()
		logrus.SetLevel(logrus.InfoLevel)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, logFormat)
	logrus.AddHook(lfHook)
}

func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println("errs", err)
	}
	writer := bufio.NewWriter(src)
	logrus.SetOutput(writer)
}
