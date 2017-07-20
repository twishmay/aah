// Copyright (c) Jeevanandam M (https://github.com/jeevatkm)
// go-aah/log source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
	"io"
	slog "log"

	"aahframework.org/config.v0"
)

var std *Logger

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Logger methods
//_______________________________________

// Error logs message as `ERROR`. Arguments handled in the mananer of `fmt.Print`.
func Error(v ...interface{}) {
	std.output(LevelError, 3, nil, v...)
}

// Errorf logs message as `ERROR`. Arguments handled in the mananer of `fmt.Printf`.
func Errorf(format string, v ...interface{}) {
	std.output(LevelError, 3, &format, v...)
}

// Warn logs message as `WARN`. Arguments handled in the mananer of `fmt.Print`.
func Warn(v ...interface{}) {
	std.output(LevelWarn, 3, nil, v...)
}

// Warnf logs message as `WARN`. Arguments handled in the mananer of `fmt.Printf`.
func Warnf(format string, v ...interface{}) {
	std.output(LevelWarn, 3, &format, v...)
}

// Info logs message as `INFO`. Arguments handled in the mananer of `fmt.Print`.
func Info(v ...interface{}) {
	std.output(LevelInfo, 3, nil, v...)
}

// Infof logs message as `INFO`. Arguments handled in the mananer of `fmt.Printf`.
func Infof(format string, v ...interface{}) {
	std.output(LevelInfo, 3, &format, v...)
}

// Debug logs message as `DEBUG`. Arguments handled in the mananer of `fmt.Print`.
func Debug(v ...interface{}) {
	std.output(LevelDebug, 3, nil, v...)
}

// Debugf logs message as `DEBUG`. Arguments handled in the mananer of `fmt.Printf`.
func Debugf(format string, v ...interface{}) {
	std.output(LevelDebug, 3, &format, v...)
}

// Trace logs message as `TRACE`. Arguments handled in the mananer of `fmt.Print`.
func Trace(v ...interface{}) {
	std.output(LevelTrace, 3, nil, v...)
}

// Tracef logs message as `TRACE`. Arguments handled in the mananer of `fmt.Printf`.
func Tracef(format string, v ...interface{}) {
	std.output(LevelTrace, 3, &format, v...)
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Logger methods - Drop-in replacement
// for Go standard logger
//_______________________________________

// Print logs message as `INFO`. Arguments handled in the mananer of `fmt.Print`.
func Print(v ...interface{}) {
	std.output(LevelInfo, 3, nil, v...)
}

// Printf logs message as `INFO`. Arguments handled in the mananer of `fmt.Printf`.
func Printf(format string, v ...interface{}) {
	std.output(LevelInfo, 3, &format, v...)
}

// Println logs message as `INFO`. Arguments handled in the mananer of `fmt.Printf`.
func Println(format string, v ...interface{}) {
	std.output(LevelInfo, 3, &format, v...)
}

// Fatal logs message as `FATAL` and call to os.Exit(1).
func Fatal(v ...interface{}) {
	std.output(levelFatal, 3, nil, v...)
	exit(1)
}

// Fatalf logs message as `FATAL` and call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	std.output(levelFatal, 3, &format, v...)
	exit(1)
}

// Fatalln logs message as `FATAL` and call to os.Exit(1).
func Fatalln(v ...interface{}) {
	std.output(levelFatal, 3, nil, v...)
	exit(1)
}

// Panic logs message as `PANIC` and call to panic().
func Panic(v ...interface{}) {
	std.output(levelPanic, 3, nil, v...)
	panic("")
}

// Panicf logs message as `PANIC` and call to panic().
func Panicf(format string, v ...interface{}) {
	std.output(levelPanic, 3, &format, v...)
	panic(fmt.Sprintf(format, v...))
}

// Panicln logs message as `PANIC` and call to panic().
func Panicln(format string, v ...interface{}) {
	std.output(levelPanic, 3, &format, v...)
	panic(fmt.Sprintf(format, v...))
}

// Writer method returns the writer of default logger.
func Writer() io.Writer {
	return std.receiver.Writer()
}

// ToGoLogger method wraps the current log writer into Go Logger instance.
func ToGoLogger() *slog.Logger {
	return std.ToGoLogger()
}

// SetDefaultLogger method sets the given logger instance as default logger.
func SetDefaultLogger(l *Logger) {
	std = l
}

// SetLevel method sets log level for default logger.
func SetLevel(level string) error {
	return std.SetLevel(level)
}

// SetPattern method sets the log format pattern for default logger.
func SetPattern(pattern string) error {
	return std.SetPattern(pattern)
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Logger level and assertion methods
//___________________________________

// Level method returns currently enabled logging level.
func Level() string {
	return std.Level()
}

// IsLevelInfo method returns true if log level is INFO otherwise false.
func IsLevelInfo() bool {
	return std.IsLevelInfo()
}

// IsLevelError method returns true if log level is ERROR otherwise false.
func IsLevelError() bool {
	return std.IsLevelError()
}

// IsLevelWarn method returns true if log level is WARN otherwise false.
func IsLevelWarn() bool {
	return std.IsLevelWarn()
}

// IsLevelDebug method returns true if log level is DEBUG otherwise false.
func IsLevelDebug() bool {
	return std.IsLevelDebug()
}

// IsLevelTrace method returns true if log level is TRACE otherwise false.
func IsLevelTrace() bool {
	return std.IsLevelTrace()
}

func init() {
	cfg, _ := config.ParseString("log { }")
	std, _ = New(cfg)
}
