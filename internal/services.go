package internal

import (
	"sync"
	"time"
)

type Logger struct {
	mu        sync.Mutex
	level     LogLevel
	appenders []Appender
	formatter Formatter
}

func NewLogger(level LogLevel, formatter Formatter, appenders ...Appender) *Logger {
	return &Logger{
		level:     level,
		appenders: appenders,
		formatter: formatter,
	}
}

func (l Logger) Log(level LogLevel, message string) {
	if level < l.level {
		return
	}
	event := LogEvent{
		Time:    time.Now(),
		Level:   level,
		Message: message,
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	_ = event
}

func (l Logger) Debug(msg string) {
	l.Log(DEBUG, msg)
}
func (l Logger) Info(msg string) {
	l.Log(INFO, msg)
}
func (l Logger) Warn(msg string) {
	l.Log(WARN, msg)
}
func (l Logger) Error(msg string) {
	l.Log(ERROR, msg)
}
