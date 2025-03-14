package log

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var GlobalLogger = NewLogger(DefaultConfig())

func FormatTheLog(level string, a ...any) []any {
	_, filename, lineno, ok := runtime.Caller(2)
	if !ok {
		filename = ""
		lineno = 0
	}
	baseFilename := filepath.Base(filename)
	logTime := time.Now().Format("2006-01-02T15:04:05 -070000")
	logMessage := []any{logTime, "[" + level + "]", fmt.Sprintf("%s:%d", baseFilename, lineno)}

	if GlobalLogger.isFormatJson {
		obj := struct {
			Time       string `json:"time"`
			Level      string `json:"loglevel"`
			Filename   string `json:"filename"`
			LineNumber int    `json:"linenumber"`
			Message    []any  `json:"message"`
		}{
			Time:       logTime,
			Level:      level,
			Filename:   baseFilename,
			LineNumber: lineno,
			Message:    a,
		}
		bytes, _ := json.Marshal(obj)
		return []any{string(bytes)}
	}
	return append(logMessage, a...)
}

func (l *Logger) log(level string, msg ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	all := FormatTheLog(level, msg...)
	fmt.Fprintln(l.file, all...)
}

func (l *Logger) Info(msg ...any) {
	l.log(INFO, msg...)
}

func (l *Logger) Debug(msg ...any) {
	l.log(DEBUG, msg...)
}

func (l *Logger) Warn(msg ...any) {
	l.log(WARNING, msg...)
}

func (l *Logger) Error(msg ...any) {
	l.log(ERROR, msg...)
}

func (l *Logger) Fatal(msg ...any) {
	l.log(FATAL, msg...)
	os.Exit(1)
}

// Global logging functions
func Info(a ...any) {
	GlobalLogger.Info(a...)
}
func Debug(a ...any) {
	GlobalLogger.Debug(a...)
}
func Error(a ...any) {
	GlobalLogger.Error(a...)
}
func Warn(a ...any) {
	GlobalLogger.Warn(a...)
}
func Fatal(a ...any) {
	GlobalLogger.Fatal(a...)
}


