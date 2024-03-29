
/**********************************
 * Author : Techie
 * Time : 2020-03-08 16:43:45
 **********************************/

package main

import (
	"fmt"
	"os"

	viper "github.com/spf13/viper"
	pflag "github.com/spf13/pflag"

	config "user/internal/pkg/config"
	application "user/internal/app"
)

func main() {
	// start
	fmt.Println("====== Server [user-grpc] Start ======")

	// init flags
	pflag.String("config.path", "../configs/config.toml", "Config file path, default '../configs/config.toml'")
	pflag.String("log.config.path", "../configs/log.config", "Log config file path, default '../configs/log.config'")

	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	// init config
	conf := config.GetConfig()

	err := conf.Init(viper.GetString("config.path"))

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

	fmt.Printf("Listen on port [%v]\r\n", conf.Server.Addr)

	// run application
	err = app.Run(conf.Server.Addr)

	if err != nil {
		fmt.Printf("Application run failed! error:%v\r\n", err.Error())
		os.Exit(-4)
	}
}
