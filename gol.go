package gol

import (
	"log"
	"os"
)

// The logger struct
type Logger struct {
	Application string
	LogFile     string
	CustomLog   bool
}

// Path generate the full file path,
func (l *Logger) Path() string {
	if l.CustomLog == true {
		return l.Application + "/" + l.LogFile + ".log"
	}

	return "/var/log/" + l.Application + "/" + l.LogFile + ".log"
}

/// NewLogger creates a logger.
func NewLogger(app string) Logger {
	return Logger{
		Application: app,
		LogFile:     "logger",
		CustomLog:   false,
	}
}

// NewCustomLogger create customizable logger.
func NewCustomLogger(app string) Logger {
	return Logger{
		Application: app,
		LogFile:     "logger",
		CustomLog:   true,
	}
}

// log do something.
func (l *Logger) log(level string, message string) {
	loggerPath := l.Path()
	f, err := os.OpenFile(
		loggerPath,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	log.SetOutput(f)
	log.Print("[" + level + "] - " + message + "")
}

// Info is vary low level, (similar to edbug).
func (l *Logger) Info(message string) {
	l.log("info", message)
}

// WArning means that something went wrong.
func (l *Logger) Warning(message string) {
	l.log("warning", message)
}

// Error level means a very dungerous stuff is happened.
func (l *Logger) Error(message string) {
	l.log("error", message)
}
