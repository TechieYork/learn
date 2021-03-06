
/**********************************
 * Author : Techie
 * Time : 2020-02-27 20:37:29
 **********************************/

package application

import (
	"context"
	"fmt"
	"time"

	"os"
	"os/signal"

	"net/http"
	_ "net/http/pprof"

	log "github.com/cihub/seelog"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	// component
	logger "github.com/DarkMetrix/gofra/pkg/logger/seelog"
	monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"
	performance "github.com/DarkMetrix/gofra/pkg/performance"
	tracing "github.com/DarkMetrix/gofra/pkg/tracing/zipkin"

	logMiddleware "github.com/DarkMetrix/gofra/pkg/gin-utils/middleware/log_middleware/seelog"
	monitorMiddleware "github.com/DarkMetrix/gofra/pkg/gin-utils/middleware/monitor_middleware/statsd"
	recoveryMiddleware "github.com/DarkMetrix/gofra/pkg/gin-utils/middleware/recovery_middleware/recovery"
	// gin relative
	"github.com/gin-gonic/gin"

	tracingInterceptor "github.com/DarkMetrix/gofra/pkg/grpc-utils/interceptor/opentracing_interceptor"
	// gRPC relative
	logInterceptor "github.com/DarkMetrix/gofra/pkg/grpc-utils/interceptor/seelog_interceptor"
	monitorInterceptor "github.com/DarkMetrix/gofra/pkg/grpc-utils/interceptor/statsd_interceptor"

	// common
	pool "github.com/DarkMetrix/gofra/pkg/grpc-utils/pool"
	commonUtils "github.com/DarkMetrix/gofra/pkg/utils"

	"user/internal/pkg/common"
	"user/internal/pkg/config"

	// http handler
	httpHandler "user/internal/http_handler"
)

var globalApplication *Application

type Application struct {
	ClientOpts []grpc.DialOption
}

// new Application
func newApplication() *Application {
	return &Application{}
}

// get singleton application
func GetApplication() *Application {
	if globalApplication == nil {
		globalApplication = newApplication()
	}

	return globalApplication
}

// init application
func (app *Application) Init(conf *config.Config) error {
	// process conf.Server.HTTPAddr
	conf.Server.HTTPAddr = commonUtils.GetRealAddrByNetwork(conf.Server.HTTPAddr)

	// init log
	err := logger.Init("../configs/log.config", common.ProjectName)

	if err != nil {
		log.Warnf("Init logger failed! error:%v", err.Error())
	}

	// init pprof
	if conf.Pprof.Active != 0 {
		go func() {
			log.Infof("Begin pprof at addr:%v", conf.Pprof.Addr)
			err = http.ListenAndServe(conf.Pprof.Addr, nil)

			if err != nil {
				log.Warnf("Pprof http.ListenAndServe failed! error:%v", err.Error())
			}
		}()
	}

	// init monitor
	err = monitor.Init(config.GetConfig().Monitor.Params...)

	if err != nil {
		log.Warnf("Init monitor failed! error:%v", err.Error())
	}

	// init tracing
	err = tracing.Init(config.GetConfig().Tracing.Params...)

	if err != nil {
		log.Warnf("Init tracing failed! error:%v", err.Error())
	}

	// set client interceptor
	app.ClientOpts = append(app.ClientOpts, grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			tracingInterceptor.GetClientInterceptor(),
			logInterceptor.GetClientInterceptor(),
			monitorInterceptor.GetClientInterceptor())), grpc.WithInsecure())

	// init pool
	err = pool.GetConnectionPool().Init(app.ClientOpts)

	if err != nil {
		log.Warnf("Init pool failed! error:%v", err.Error())
		return err
	}

	// init performance
	if conf.Performance.Active != 0 {
		switch conf.Performance.Type {
		case "log":
			go performance.BeginMemoryPerformanceMonitorWithLog()
			go performance.BeginGoroutinePerformanceMonitorWithLog()
		case "statsd":
			go performance.BeginMemoryPerformanceMonitorWithStatsd()
			go performance.BeginGoroutinePerformanceMonitorWithStatsd()
		default:
			log.Warnf("Performance type not found! Type:%v", conf.Performance.Type)
		}
	}

	return nil
}

// run application
func (app *Application) Run(address string) error {
	defer log.Flush()
	defer tracing.Close()

	// run to serve http
	httpClose, err := app.runHTTPServer(address)

	if err != nil {
		log.Warnf("app.runHTTPServer failed! error:%v", err.Error())
		return err
	}

	defer httpClose()

	// deal with signals, when interrupt was notified, server will stop gracefully
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	signalOccur := <- signalChannel

	log.Infof("Signal occurred, signal:%v", signalOccur.String())

	return nil
}

type httpCloseFunc func()

func (app *Application) runHTTPServer(address string) (httpCloseFunc, error) {
	// check address
	if address == "" {
		return nil, fmt.Errorf("HTTP listen address is empty! address:%v", address)
	}

	// set release or debug mode
	if config.GetConfig().Server.GinDebug == 0 {
		gin.SetMode(gin.ReleaseMode)
		log.Infof("gin runs in release mode!")
	} else {
		gin.SetMode(gin.DebugMode)
		log.Infof("gin runs in debug mode!")
	}

	// init engine
	engine := gin.Default()

	group := engine.Group("/",
		recoveryMiddleware.GetMiddleware(),
		logMiddleware.GetMiddleware(),
		monitorMiddleware.GetMiddleware())

	// add http handler
	// !!!DO NOT EDIT!!!
	group.GET("/health", httpHandler.GET_HEALTH)
	group.GET("/user/:id", httpHandler.GET_USER_ID)
	group.POST("/user/batch", httpHandler.POST_USER_BATCH)
	/*@REGISTER_HTTP_STUB*/

	httpServer := &http.Server{
		Addr: address,
		Handler: engine,
	}

	go func() {
		err := httpServer.ListenAndServe()

		if err != nil {
			log.Errorf("Serve http failed! error:%v", err.Error())

			time.Sleep(time.Second)
			os.Exit(-2)
		} else {
			log.Infof("Serve http quit")
		}
	}()

	return func(){
		// stop http service gracefully
		ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()

		err := httpServer.Shutdown(ctx)

		if err != nil {
			log.Warnf("Http server stopped gracefully failed! error:%s", err.Error())
		} else {
			log.Infof("Http server stopped gracefully!")
		}
	}, nil
}
