
/**********************************
 * Author : Techie
 * Time : 2020-03-11 10:18:28
 **********************************/

package application

import (
	"net"
	"os"
	"os/signal"
	"time"

	"net/http"
	_ "net/http/pprof"

	log "github.com/cihub/seelog"
	viper "github.com/spf13/viper"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	logger "github.com/DarkMetrix/gofra/pkg/logger/seelog"
	monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"
	performance "github.com/DarkMetrix/gofra/pkg/performance"
	tracing "github.com/DarkMetrix/gofra/pkg/tracing/zipkin"

	tracingInterceptor "github.com/DarkMetrix/gofra/pkg/grpc-utils/interceptor/opentracing_interceptor"
	recoverInterceptor "github.com/DarkMetrix/gofra/pkg/grpc-utils/interceptor/recover_interceptor"
	logInterceptor "github.com/DarkMetrix/gofra/pkg/grpc-utils/interceptor/seelog_interceptor"
	monitorInterceptor "github.com/DarkMetrix/gofra/pkg/grpc-utils/interceptor/statsd_interceptor"

	pool "github.com/DarkMetrix/gofra/pkg/grpc-utils/pool"
	commonUtils "github.com/DarkMetrix/gofra/pkg/utils"

	"goods-grpc/internal/pkg/common"
	"goods-grpc/internal/pkg/config"

	// !!!DO NOT EDIT!!!
	health_check "goods-grpc/api/protobuf_spec/health_check"
	goods "goods-grpc/api/protobuf_spec/goods"
	/*@PROTO_STUB*/
	HealthCheckServiceHandler "goods-grpc/internal/grpc_handler/HealthCheckService"
	GoodsServiceHandler "goods-grpc/internal/grpc_handler/GoodsService"
	/*@HANDLER_STUB*/
)

var globalApplication *Application

type Application struct {
	ServerOpts []grpc.ServerOption
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
	// process conf.Server.Addr
	conf.Server.Addr = commonUtils.GetRealAddrByNetwork(conf.Server.Addr)

	// init log
	err := logger.Init(viper.GetString("log.config.path"), common.ProjectName)

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

	// set server interceptor
	app.ServerOpts = append(app.ServerOpts, grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			recoverInterceptor.GetServerInterceptor(),
			tracingInterceptor.GetServerInterceptor(),
			logInterceptor.GetServerInterceptor(),
			monitorInterceptor.GetServerInterceptor())))

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
	defer pool.GetConnectionPool().Close()

	// run to serve grpc
	grpcClose, err := app.runGRPCServer(address)

	if err != nil {
		log.Warnf("app.runGRPCServer failed! error:%v", err.Error())
		return err
	}

	defer grpcClose()

	// deal with signals, when interrupt was notified, server will stop gracefully
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	signalOccur := <- signalChannel

	log.Infof("Signal occurred, signal:%v", signalOccur.String())

	return nil
}

type grpcCloseFunc func()

func (app *Application) runGRPCServer(address string) (grpcCloseFunc, error) {
	// listen
	listen, err := net.Listen("tcp", address)

	if err != nil {
		return nil, err
	}

	// init grpc server
	s := grpc.NewServer(app.ServerOpts ...)

	// register services
	// !!!DO NOT EDIT!!!
	health_check.RegisterHealthCheckServiceServer(s, HealthCheckServiceHandler.HealthCheckServiceImpl{})
	goods.RegisterGoodsServiceServer(s, GoodsServiceHandler.GoodsServiceImpl{})
	/*@REGISTER_STUB*/

	// run to serve
	go func() {
		err = s.Serve(listen)

		if err != nil {
			log.Errorf("Serve gRPC failed! error:%v", err.Error())

			time.Sleep(time.Second)
			os.Exit(-2)
		} else {
			log.Infof("Serve gRPC quit!")
		}
	}()

	return func() {
		// stop grpc service gracefully
		s.GracefulStop()

		log.Infof("gRPC server stopped gracefully!")
	}, nil
}
