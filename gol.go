package gol

import (
	"log"
	"os"
)

type Logger struct {
	Application string
	LogFile     string
}

func (l *Logger) Info(message string) {
	app := l.Application
	level := "info"
	logfile := l.LogFile

	f, err := os.OpenFile(
		"/var/log/"+app+"/"+logfile+".log",
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
