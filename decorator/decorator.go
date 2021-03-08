package decorator

import (
	"encoding/json"
	"fmt"
	"time"
)

// 日志器
type ILogger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

// 日志器工厂
type ILoggerFactory interface {
	GetLogger(name string) ILogger
}

// 使用 console 来输出日志
type consoleLogger struct {
	name string
}

func nowString() string {
	return time.Now().Format("2006-01-02T15:04:05")
}

func (c *consoleLogger) log(msg string, level string) {
	fmt.Printf("%s [%s] %s %s\n", nowString(), c.name, level, msg)
}

func (c *consoleLogger) Debug(msg string) {
	c.log(msg, "DEBUG")
}

func (c *consoleLogger) Info(msg string) {
	c.log(msg, "INFO")
}

func (c *consoleLogger) Warn(msg string) {
	c.log(msg, "WARN")
}

func (c *consoleLogger) Error(msg string) {
	c.log(msg, "ERROR")
}

// 默认的日志器工厂
var ConsoleLoggerFactory = newConsoleLoggerFactory()

type consoleLoggerFactory struct {
}

func newConsoleLoggerFactory() ILoggerFactory {
	return &consoleLoggerFactory{}
}

func (cf *consoleLoggerFactory) GetLogger(name string) ILogger {
	return &consoleLogger{name: name}
}

// json 格式输出的日志器
type jsonLogger struct {
	xName   string
	xLogger ILogger
}

// json 格式的信息
type jsonMsg struct {
	Time  string
	Level string
	Msg   string
}

func (j *jsonMsg) String() string {
	js, err := json.Marshal(j)
	if err != nil {
		return ""
	}
	return string(js)
}

func newJsonLogger(name string, logger ILogger) ILogger {
	return &jsonLogger{
		xName:   name,
		xLogger: logger,
	}
}

func (j *jsonLogger) toJson(msg string, level string) string {
	js := &jsonMsg{
		Time:  nowString(),
		Level: level,
		Msg:   msg,
	}
	return js.String()
}

func (j *jsonLogger) Debug(msg string) {
	j.xLogger.Debug(j.toJson(msg, "DEBUG"))
}

func (j *jsonLogger) Info(msg string) {
	j.xLogger.Info(j.toJson(msg, "INFO"))
}

func (j *jsonLogger) Warn(msg string) {
	j.xLogger.Warn(j.toJson(msg, "WARN"))
}

func (j *jsonLogger) Error(msg string) {
	j.xLogger.Error(j.toJson(msg, "ERROR"))
}

// json 日志器的工厂
type jsonLoggerFactory struct {
}

func newJsonLoggerFactory() ILoggerFactory {
	return &jsonLoggerFactory{}
}

// 在 consoleLogger 的基础上加上了 json 的装饰，就成了 jsonLogger
func (jf *jsonLoggerFactory) GetLogger(name string) ILogger {
	logger := ConsoleLoggerFactory.GetLogger(name)
	return newJsonLogger(name, logger)
}

var JsonLoggerFactory = newJsonLoggerFactory()
