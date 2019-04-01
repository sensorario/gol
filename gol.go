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

func NewLogger(app string) Logger {
	return Logger{
		Application: app,
		LogFile:     "logger",
	}
}

func (l *Logger) log(level string, message string) {
	app := l.Application
	logfile := l.LogFile

	loggerPath := "/var/log/" + app + "/" + logfile + ".log"
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
