// Copyright 2020 Jing Li. All rights reserved.

/*
Package log implements a logger
 - It extends the built-in log package and provides the methods to log at
 selected Syslog levels
 - It formats and writes each log line to Logger's $Wi in the format of:
   - <utc-date> <utc-time> <file:line> <level-char> <formatted-message>
 - It is thread-safe for using the built-in log package underneath, but has no
 supports of log rotation and throttling
*/
package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Level int

// A subset of Syslog levels
const (
	LogPanic Level = 0
	LogError       = 3
	LogWarn        = 4
	LogInfo        = 6
	LogDebug       = 7
)

var logSs = [...]string{"p", "", "", "e", "w", "", "i", "d"}

type Logger struct {
	Log   *log.Logger
	Level Level
	Fn    string
	Wi    io.Writer
}

/*
Create and return a logger $l to write log lines to $i
 - $i accepts 1) a filename string or 2) an object implementing io.Writer
 - If error happens, it closes the opened file provided by $i string
 - It sets $l.Level to $LogInfo by default, so it writes nothing for Debug
 level unless the env-var WG_DEBUG has been set
*/
func NewLog(i interface{}) (l *Logger, e error) {
	l = &Logger{}

	defer func() {
		// Finalize $l if error happens
		if e != nil {
			FreeLog(l)
		}
	}()

	// Assert type
	switch v := i.(type) {
	case string:
		l.Fn = v
	case io.Writer:
		l.Wi = v
	default:
		e = fmt.Errorf("Invalid type: %v", v)
		return
	}

	// Open $l.Fn if any
	if l.Fn != "" {
		l.Wi, e = os.OpenFile(l.Fn, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
		if e != nil {
			return
		}
	}

	// Create and init logger
	l.Log = log.New(l.Wi, "", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	// Set level
	l.Level = LogInfo
	_, ok := os.LookupEnv("WG_DEBUG")
	if ok {
		l.Level = LogDebug
	}

	return
}

// Finalize $l
func FreeLog(l *Logger) (e error) {
	if l != nil && l.Fn != "" && l.Wi != nil {
		e = l.Wi.(*os.File).Close()
	}
	return
}

func logAt(l *Logger, level Level, fs string, as ...interface{}) {
	// Sanity checks
	if l == nil || l.Log == nil {
		return
	}
	// Check $level
	if level > l.Level {
		return
	}

	// Use 3 as the call-stack depth for having this func in the middle
	l.Log.Output(3, logSs[level]+" "+fmt.Sprintf(fs, as...))
}

// Format $fs and $as and write it at Panic level
func (l *Logger) Panic(fs string, as ...interface{}) {
	logAt(l, LogPanic, fs, as...)
}

// Format $fs and $as and write it at Error level
func (l *Logger) Error(fs string, as ...interface{}) {
	logAt(l, LogError, fs, as...)
}

// Format $fs and $as and write it at Warn level
func (l *Logger) Warn(fs string, as ...interface{}) {
	logAt(l, LogWarn, fs, as...)
}

// Format $fs and $as and write it at Info level
func (l *Logger) Info(fs string, as ...interface{}) {
	logAt(l, LogInfo, fs, as...)
}

// Format $fs and $as and write it at Debug level
func (l *Logger) Debug(fs string, as ...interface{}) {
	logAt(l, LogDebug, fs, as...)
}
