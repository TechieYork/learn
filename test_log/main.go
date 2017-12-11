package main

import (
	"fmt"

	log4go "github.com/alecthomas/log4go"
	seelog "github.com/cihub/seelog"
	"time"
)

func testLog4go() {
	fmt.Println("====== Test log4go ======")

	// =OR= Can also specify manually via the following: (these are the defaults, this is equivalent to above)
	flw := log4go.NewFileLogWriter("log4go_test.log", false)
	flw.SetFormat("[%D %T] [%L] (%S) %M")
	flw.SetRotate(true)
	flw.SetRotateSize(500 * 1024 * 1024)
	flw.SetRotateLines(1000000)
	flw.SetRotateDaily(true)
	log4go.AddFilter("file", log4go.DEBUG, flw)
	//log4go.AddFilter("stdout", log4go.DEBUG, log4go.NewConsoleLogWriter())

	log4go.Debug("Test log")
	log4go.Finest("Test log")
	log4go.Fine("Test log")
	log4go.Trace("Test log")
	log4go.Info("Test log")
	log4go.Warn("Test log")
	log4go.Error("Test log")
	log4go.Critical("Test log")
}

func testSeelog() {
	fmt.Println("====== Test seelog ======")

	defer seelog.Flush()

	logger, err := seelog.LoggerFromConfigAsString(`
		<seelog>
			<outputs formatid="main">
				<rollingfile type="size" filename="./seelog_test.log" maxsize="5242888000" maxrolls="1000000"/>
				<console/>
			</outputs>
			<formats>
	    		<format id="main" format="%DateT%Time [%LEVEL] (%File %Func:%Line) => %Msg%n"/>
			</formats>
		</seelog>
	`)

	if err != nil {
		return
	}

	seelog.ReplaceLogger(logger)

	seelog.Trace("Test log")
	seelog.Debug("Test log")
	seelog.Info("Test log")
	seelog.Warn("Test log")
	seelog.Error("Test log")
	seelog.Critical("Test log")
}

func main() {
	fmt.Println("====== Test log ======")

	testLog4go()
	time.Sleep(time.Second * 1)
	testSeelog()
}
