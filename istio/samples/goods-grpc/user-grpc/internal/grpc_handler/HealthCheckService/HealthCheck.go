
/**********************************
 * Author : Techie
 * Time : 2020-03-08 16:43:45
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

	pb "user/api/protobuf_spec/health_check"
)

func (this HealthCheckServiceImpl) HealthCheck (ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	// log.Infof("====== HealthCheck start ======")

	resp := new(pb.HealthCheckResponse)

	return resp, nil
}
