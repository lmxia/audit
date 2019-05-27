package common

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLogger init log options
func InitLogger() {
	logrus.SetLevel(logrus.DebugLevel)
	defaultWriter := &lumberjack.Logger{
		Filename:   "./blockly.log",
		MaxSize:    50, // megabytes
		MaxBackups: 5,
		MaxAge:     28,    //days
		Compress:   false, // disabled by default
	}
	logrus.SetOutput(defaultWriter)
}

//Logger return a self defined handler for log info
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		logrus.Infof("|%3d |%13v |%15s |%s |%s |",
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
