package tools

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Entry

func NewLogger() {
	logStore := logrus.New()
	logStore.SetLevel(logrus.DebugLevel)

	// 同时写到多个输出
	w1 := os.Stdout //写到控制台
	// 写到文件
	w2, _ := os.OpenFile("./vote.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	logStore.SetOutput(io.MultiWriter(w1, w2)) // io.MultiWriter 返回一个 io.Writer 对象

	logStore.SetFormatter(&logrus.JSONFormatter{})
	Logger = logStore.WithFields(logrus.Fields{
		"name": "磊",
		"app":  "voteV2",
	})

	// 钩子函数
	// logStore.AddHook()  // 重大错误邮箱通知，日志分割等等

	// 用的最多
	// 上下文context
	// logStore.WithContext()
}
