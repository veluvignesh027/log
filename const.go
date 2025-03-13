package main

import (
	"io"
	"os"
	"sync"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
)

const (
	INFO    = Green + "INFO" + Reset
	ERROR   = Red + "ERROR" + Reset
	DEBUG   = Yellow + "DEBUG" + Reset
	WARNING = Blue + "WARNING" + Reset
	FATAL   = Magenta + "FATAL" + Reset
)

type Logger struct {
	name         string
	file         io.Writer
	isWriteFile  bool
	isFormatJson bool
	mu           sync.Mutex
}

type Config struct {
	LogName      string
	IsFileLog    bool
	File         io.Writer
	IsFormatJson bool
}

func DefaultConfig() Config {
	return Config{
		LogName:   "GLOBAL",
		IsFileLog: false,
		File:      os.Stdout,
	}
}

func NewLogger(config Config) *Logger {
	return &Logger{
		name:         config.LogName,
		file:         config.File,
		isWriteFile:  config.IsFileLog,
		isFormatJson: config.IsFormatJson,
	}
}
