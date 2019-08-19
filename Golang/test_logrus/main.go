package main

import (
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})

	log.Infof("hello!{}:")

	// deal with signals, when interrupt was notified, server will stop gracefully
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	signalOccur := <- signalChannel

	log.Infof("Signal occured, signal:%v", signalOccur.String())
}
