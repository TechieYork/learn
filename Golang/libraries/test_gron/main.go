package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/roylee0704/gron"
	"github.com/roylee0704/gron/xtime"
)

func main() {
	cron := gron.New()

	//Caution: gron used UTC time
	cron.AddFunc(gron.Every(1 * xtime.Day).At("11:00"), func() {
		fmt.Printf("Bingo!")
	})

	cron.Start()
	defer cron.Stop()

	// deal with signals, when interrupt was notified, server will stop gracefully
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	signalOccur := <- signalChannel

	fmt.Sprintf("Signal occured, signal:%v", signalOccur.String())
}
