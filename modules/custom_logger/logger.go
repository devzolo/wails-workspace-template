package custom_logger

import (
	"strings"
)

const (
	TRACE = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

type MyLogger struct {
	level int
}

func NewMyLogger() *MyLogger {
	return &MyLogger{level: INFO}
}

func (m *MyLogger) SetLogLevel(level int) {
	m.level = level
}

func (m *MyLogger) isBlocked(message string) bool {
	if strings.HasPrefix(message, "[AssetHandler] Handling request") {
		return true
	}
	if strings.HasPrefix(message, "[AssetHandler] File") && strings.Contains(message, "not found, serving") {
		return true
	}
	if strings.HasPrefix(message, "[ExternalAssetHandler] Loading") {
		return true
	}
	return false
}

func (m *MyLogger) Write(level int, message string) {
	if level >= m.level {
		if m.isBlocked(message) {
			return
		}

		println(message)
	}
}

func (m *MyLogger) Print(message string) {
	m.Write(TRACE, message)
}

func (m *MyLogger) Trace(message string) {
	m.Write(TRACE, message)
}

func (m *MyLogger) Debug(message string) {
	m.Write(DEBUG, message)
}

func (m *MyLogger) Info(message string) {
	m.Write(INFO, message)
}

func (m *MyLogger) Warning(message string) {
	m.Write(WARNING, message)
}

func (m *MyLogger) Error(message string) {
	m.Write(ERROR, message)
}

func (m *MyLogger) Fatal(message string) {
	m.Write(FATAL, message)
}
