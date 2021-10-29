package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type LogFormatter struct{}

func (m *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var content = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	return []byte(content), nil
}

func InitAppLog() {
	Log = logrus.New()
	Log.SetFormatter(&LogFormatter{})
	Log.SetReportCaller(true)

	writer, err := rotatelogs.New(
		Config.App.LogPath+"log_%Y%m%d%H.log",
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(time.Duration(24)*time.Hour),
	)
	if err != nil {
		Log.Fatal(err)
	}
	Log.SetOutput(io.MultiWriter(writer, os.Stdout))
}

func LogMiddle(c *gin.Context) {
	var request string
	request += fmt.Sprintf("[request]\n%s %s %s %s\n", c.Request.Method, c.Request.URL, c.Request.Host, c.Request.RemoteAddr)
	request += "HEADER\n"
	for k, v := range c.Request.Header {
		request += fmt.Sprintf("%s:%s\n", k, strings.Join(v, ""))
	}
	request += "BODY\n"
	body, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		Log.Warn(err)
	}
	request += string(body) + "\n"
	Log.Info(request)
	c.Next()
}
