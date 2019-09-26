package main

import (
	"os"
	"os/signal"

	log "github.com/cihub/seelog"
	"fmt"
)

func main() {
	//If path is empty, use default setting
	var logger log.LoggerInterface
	var err error

	/*
	//Personal formatter
	err = log.RegisterCustomFormatter("Project", createProjectFormatter)

	if err != nil {
		fmt.Printf("Register custom formatter failed! error:%v\r\n", err.Error())
		return
	}
	*/

	logger, err = log.LoggerFromConfigAsFile("./log.config")

	if err != nil {
		fmt.Printf("Load config failed! error:%v\r\n", err.Error())
		return
	}

	err = log.ReplaceLogger(logger)

	if err != nil {
		fmt.Printf("Replace logger failed! error:%v\r\n", err.Error())
		return
	}

	log.Info("log init success!")

	// deal with signals, when interrupt was notified, server will stop gracefully
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	signalOccur := <- signalChannel

	log.Infof("Signal occured, signal:%v", signalOccur.String())

}

/*
func createProjectFormatter(params string) log.FormatterFunc {
	return func(message string, level log.LogLevel, context log.LogContextInterface) interface{} {
		return "test_seelog"
	}
}
*/
