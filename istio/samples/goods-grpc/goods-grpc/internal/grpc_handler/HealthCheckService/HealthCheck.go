
/**********************************
 * Author : Techie
 * Time : 2020-03-11 10:18:28
 **********************************/

package HealthCheckService 

import (
	"context"

	// log package
	// log "github.com/cihub/seelog"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	// opentracing "github.com/opentracing/opentracing-go"

	pb "goods-grpc/api/protobuf_spec/health_check"
)

func (this HealthCheckServiceImpl) HealthCheck (ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	// log.Infof("====== HealthCheck start ======")

	resp := new(pb.HealthCheckResponse)

	return resp, nil
}
