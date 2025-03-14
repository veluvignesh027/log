# Logger Package

This package provides a simple logging utility in Go. It allows you to configure logging to either the console (stdout) or to a file, with options for JSON formatting.

## Configuration

The logging behavior is controlled by the `Config` struct, which includes the following fields:

- `LogName` (string): The name of the logger. Default is "GLOBAL".
- `IsFileLog` (bool): A flag indicating whether to log to a file. Default is `false`.
- `File` (io.Writer): The output destination for logs. Default is `os.Stdout`.
- `IsFormatJson` (bool): A flag indicating whether to format logs as JSON. Default is `false`.

## Functions

### DefaultConfig

```go
func DefaultConfig() Config
func NewLogger(config Config) *Logger


package main

import (
	"os"
	"log"
)

func main() {
	config := DefaultConfig()
	config.IsFileLog = true
	config.File, _ = os.OpenFile("output.log", os.O_CREATE|os.O_APPEND, 0666)
	l := NewLogger(config)
	l.Debug("hello")
	l.Info("hello")
	l.Warn("hello")
}
