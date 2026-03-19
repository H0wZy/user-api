package config

import (
	"io"
	"log"
	"os"
)

// ANSI escape codes for colors
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, Blue+"DEBUG: "+Reset, logger.Flags()),
		info:    log.New(writer, Green+"INFO: "+Reset, logger.Flags()),
		warning: log.New(writer, Yellow+"WARNING: "+Reset, logger.Flags()),
		err:     log.New(writer, Red+"ERROR: "+Reset, logger.Flags()),
		writer:  writer,
	}
}

// Simple logs

func (l *Logger) Debug(v ...any) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...any) {
	l.info.Println(v...)
}
func (l *Logger) Warn(v ...any) {
	l.warning.Println(v...)
}
func (l *Logger) Error(v ...any) {
	l.err.Println(v...)
}

// Formated logs

func (l *Logger) Debugf(format string, v ...any) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}
func (l *Logger) Warnf(format string, v ...any) {
	l.warning.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...any) {
	l.err.Printf(format, v...)
}
