package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// 自定义日志中间件
func Logger() gin.HandlerFunc {
	filePath := "log/mclog"
	// 软链接路径，链接到最新的日志文件
	linkName := "mclog.log"
	logFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err: ", err)
	}
	logger := logrus.New()

	logger.Out = logFile

	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := rotatelogs.New(
		filePath+".%Y%m%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.TraceLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(hook)

	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		StopTime := time.Since(startTime)
		duratime := fmt.Sprintf("%d ms", int(float64(StopTime.Milliseconds())))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unkonwn"
		}
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()
		userAgent := ctx.Request.UserAgent()
		dataSize := ctx.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := ctx.Request.Method
		path := ctx.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName": hostName,
			"Status":   statusCode,
			"DuraTime": duratime,
			"IP":       clientIP,
			"Method":   method,
			"Path":     path,
			"Agent":    userAgent,
		})

		if len(ctx.Errors) > 0 {
			entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		}

		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
