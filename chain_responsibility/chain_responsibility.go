package chain_responsibility

import (
	"fmt"
	"io"
)

// 日志器接口
type ILogger interface {
	Debug(msg string)
	Info(msg string)
	Error(msg string)
}

// 实现 ILogger 接口, 内部使用责任链模式分别处理 Debug/Info/Error 请求
type simpleLogger struct {
	chain ILoggerFilter
}

func newSimpleLogger() ILogger {
	vErrorLogger := newLoggerFilter(newFileWriter("error.log"), LEVEL_ERROR, nil)
	vInfoLogger := newLoggerFilter(newFileWriter("info.log"), LEVEL_INFO, nil)
	vDebugLogger := newLoggerFilter(newFileWriter("debug.log"), LEVEL_DEBUG, nil)

	vDebugLogger.Next(vInfoLogger)
	vInfoLogger.Next(vErrorLogger)

	return &simpleLogger{chain: vDebugLogger}
}

func (s *simpleLogger) Debug(msg string) {
	s.chain.Handle(LEVEL_DEBUG, msg)
}

func (s *simpleLogger) Info(msg string) {
	s.chain.Handle(LEVEL_INFO, msg)
}

func (s *simpleLogger) Error(msg string) {
	s.chain.Handle(LEVEL_ERROR, msg)
}

// 日志责任链节点的接口
type LoggingLevel string

const LEVEL_DEBUG LoggingLevel = "DEBUG"
const LEVEL_INFO LoggingLevel = "INFO"
const LEVEL_ERROR LoggingLevel = "ERROR"

type ILoggerFilter interface {
	Next(filter ILoggerFilter)
	Handle(level LoggingLevel, msg string)
}

// 实现 ILoggerFilter 接口
type loggerFilter struct {
	writer io.StringWriter
	level  LoggingLevel
	chain  ILoggerFilter
}

func newLoggerFilter(writer io.StringWriter, level LoggingLevel, filter ILoggerFilter) ILoggerFilter {
	return &loggerFilter{
		writer: writer,
		level:  level,
		chain:  filter,
	}
}

func (l *loggerFilter) Next(filter ILoggerFilter) {
	l.chain = filter
}

func (l *loggerFilter) Handle(level LoggingLevel, msg string) {
	if l.level == level {
		_, _ = l.writer.WriteString(fmt.Sprintf("%v %s", l.level, msg))
	} else {
		if l.chain != nil {
			l.chain.Handle(level, msg)
		}
	}
}

// 负责日志输出, 实现 io.StringWriter 接口
type fileWriter struct {
	file string
}

func newFileWriter(file string) io.StringWriter {
	return &fileWriter{
		file: file,
	}
}

func (f *fileWriter) WriteString(s string) (n int, e error) {
	fmt.Printf("fileWriter.WriteString, file=%s, msg=%s\n", f.file, s)
	return len(s), nil
}
