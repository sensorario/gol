package gol

import (
	"log"
	"os"

	"golang.org/x/sys/unix"
)

type Logger struct {
	Application string
	LogFile     string
}

func (l *Logger) Info(message string) {
	level := "info"
	app := l.Application
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

	path := "/var/log/" + app
	if unix.Access(path, unix.W_OK) != nil {
		log.Print("[warning] - cant write into `" + path + "` log")
	}
}
