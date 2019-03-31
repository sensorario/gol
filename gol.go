package gol

import (
	"log"
	"os"
)

func Info(message string) {
	level := "info"

	f, err := os.OpenFile(
		"/var/log/fussy/fussy.log",
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
