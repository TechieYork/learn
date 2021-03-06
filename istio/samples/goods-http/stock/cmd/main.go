
/**********************************
 * Author : Techie
 * Time : 2020-02-29 11:03:20
 **********************************/

package main

import (
	"fmt"
	"os"

	application "stock/internal/app"
	config "stock/internal/pkg/config"
)

func main() {
	// start
	fmt.Println("====== Server [stock] Start ======")

	// init config
	conf := config.GetConfig()

	err := conf.Init("../configs/config.toml")

	if err != nil {
		fmt.Printf("Init config failed! error:%v\r\n", err.Error())
		os.Exit(-1)
	}

	// init application
	app := application.GetApplication()

	if app == nil {
		fmt.Printf("Application get failed!\r\n")
		os.Exit(-2)
	}

	err = app.Init(conf)

	if err != nil {
		fmt.Printf("Application init failed! error:%v\r\n", err.Error())
		os.Exit(-3)
	}

	fmt.Printf("Listen on port [%v]\r\n", conf.Server.HTTPAddr)

	// run application
	err = app.Run(conf.Server.HTTPAddr)

	if err != nil {
		fmt.Printf("Application run failed! error:%v\r\n", err.Error())
		os.Exit(-4)
	}
}
