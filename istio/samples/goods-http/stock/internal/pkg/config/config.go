
/**********************************
 * Author : Techie
 * Time : 2020-02-29 11:03:20
 **********************************/

package config

import (
	"github.com/spf13/viper"
)

// global config
var globalConfig *Config

// pprof
type PprofInfo struct {
	Addr string "mapstructure:\"addr\" json:\"addr\""
	Active uint32 "mapstructure:\"active\" json:\"active\""
}

// performance info
type PerformanceInfo struct {
	Active uint32 "mapstructure:\"active\" json:\"active\""
	Type string "mapstructure:\"type\" json:\"type\""
}

// tracing info
type TracingInfo struct {
	Params []string "mapstructure:\"params\" json:\"params\""
}

// monitor info
type MonitorInfo struct {
	Params []string "mapstructure:\"params\" json:\"params\""
}

// server config
type ServerInfo struct {
	HTTPAddr string "mapstructure:\"http_addr\" json:\"http_addr\""
	GinDebug int "mapstructure:\"gin_debug\" json:\"gin_debug\""
}

// client config
type ClientInfo struct {
}

// pool config
type PoolInfo struct {
	InitConns int "mapstructure:\"init_conns\" json:\"init_conns\""
	MaxConns int "mapstructure:\"max_conns\" json:\"max_conns\""
	IdleTime int "mapstructure:\"idle_time\" json:\"idle_time\""
}

// config structure
type Config struct {
	Pprof PprofInfo "mapstructure:\"pprof\" json:\"pprof\""
	Performance PerformanceInfo "mapstructure:\"performance\" json:\"performance\""
	Monitor MonitorInfo "mapstructure:\"monitor\" json:\"monitor\""
	Tracing TracingInfo "mapstructure:\"tracing\" json:\"tracing\""
	Server ServerInfo "mapstructure:\"server\" json:\"server\""
	Client ClientInfo "mapstructure:\"client\" json:\"client\""
}

// new Config
func newConfig() *Config {
	return &Config{}
}

// get singleton config
func GetConfig() *Config {
	if globalConfig == nil {
		globalConfig = newConfig()
	}

	return globalConfig
}

// init config from json file
func (config *Config) Init (path string) error {
	// set viper setting
	viper.SetConfigType("toml")
	viper.SetConfigFile(path)
	viper.AddConfigPath("../configs/")

	// read in config
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	// unmarshal config
	err = viper.Unmarshal(config)

	if err != nil {
		return err
	}

	return nil
}
