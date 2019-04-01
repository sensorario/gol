package gol

import (
	"log"
	"os"
)

type Logger struct {
	Application string
	LogFile     string
	CustomLog   bool
}

func (l *Logger) Path() string {
	if l.CustomLog == true {
		return l.Application + "/" + l.LogFile + ".log"
	}

	return "/var/log/" + l.Application + "/" + l.LogFile + ".log"
}

func NewLogger(app string) Logger {
	return Logger{
		Application: app,
		LogFile:     "logger",
		CustomLog:   false,
	}
}

func NewCustomLogger(app string) Logger {
	return Logger{
		Application: app,
		LogFile:     "logger",
		CustomLog:   true,
	}
}

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

func (l *Logger) Info(message string) {
	l.log("info", message)
}

func (l *Logger) Error(message string) {
	l.log("error", message)
}
