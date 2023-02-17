package main

import (
	"os"
  "time"
)

func appendLog(file string, action string, detail string) {
	// open file
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// do logging
	if _, err := f.WriteString(time.Now().Format("2006-01-02 15:04") + " [" + action + "] " + detail + "\n"); err != nil {
		panic(err)
	}

	// flush file stream
	f.Sync()
}

func startLogger(logFile string, msgs chan Message) {
	appendLog(logFile, "START", "Logging Started.")

	// handle msgs coming over the channel
	for {
		msg := <-msgs
		if msg.command == Quit {
		appendLog(logFile, "STOP", "Logging Stopped.")
			return
		}

		appendLog(logFile, msg.command, msg.message)
	}
}
