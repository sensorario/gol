package gol

import (
	"log"
	"os"
)

type Conf struct {
	Application string
	LogFile     string
}

func Info(message string, conf Conf) {
	level := "info"
	app := conf.Application
	logfile := conf.LogFile

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
