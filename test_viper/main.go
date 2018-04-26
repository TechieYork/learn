package main

import (
	"github.com/spf13/viper"
	"fmt"
)

type TracingInfo struct {
	Params []string "mapstructure:\"params\" json:\"params\""
}

type MonitorInfo struct {
	Params []string "mapstructure:\"params\" json:\"params\""
}

type ClientInfo struct {
	Pool PoolInfo "mapstructure:\"pool\" json:\"pool\""
}

type PoolInfo struct {
	InitConns int "mapstructure:\"init_conns\" json:\"init_conns\""
	MaxConns int "mapstructure:\"max_conns\" json:\"max_conns\""
	IdleTime int "mapstructure:\"idle_time\" json:\"idle_time\""
}

type ServerInfo struct {
	Addr string "mapstructure:\"addr\" json:\"addr\""
}

type Config struct {
	Monitor MonitorInfo "mapstructure:\"monitor\" json:\"monitor\""
	Tracing TracingInfo "mapstructure:\"tracing\" json:\"tracing\""
	Server ServerInfo "mapstructure:\"server\" json:\"server\""
	Client ClientInfo "mapstructure:\"client\" json:\"client\""
}

func testViperToml(path string) error {
	//Set viper setting
	viper.SetConfigType("toml")
	viper.SetConfigFile(path)
	viper.AddConfigPath("./conf/")

	//Read in config
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Read failed! error:%s\r\n", err.Error())
		return err
	}

	//Unmarshal config
	config := new(Config)
	err = viper.Unmarshal(config)

	if err != nil {
		fmt.Printf("Unmarshal failed! error:%s\r\n", err.Error())
		return err
	}

	fmt.Printf("testViperJson config:\r\n%v\r\n", config)

	return nil
}

func testViperJson(path string) error {
	//Set viper setting
	viper.SetConfigType("json")
	viper.SetConfigFile(path)
	viper.AddConfigPath("./conf/")

	//Read in config
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Read failed! error:%s\r\n", err.Error())
		return err
	}

	//Unmarshal config
	config := new(Config)
	err = viper.Unmarshal(config)

	if err != nil {
		fmt.Printf("Unmarshal failed! error:%s\r\n", err.Error())
		return err
	}

	fmt.Printf("testViperJson config:\r\n%v\r\n", config)

	return nil
}

func main() {
	err := testViperJson("./conf/config.json")

	if err != nil {
		return
	}

	err = testViperToml("./conf/config.toml")

	if err != nil {
		return
	}
}
