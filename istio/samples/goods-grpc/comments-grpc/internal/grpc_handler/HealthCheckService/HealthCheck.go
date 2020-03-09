
/**********************************
 * Author : Techie
 * Time : 2020-03-08 16:43:04
 **********************************/

package HealthCheckService 

import (
	"context"

	// log package
	// log "github.com/cihub/seelog"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	// tracing "github.com/DarkMetrix/gofra/pkg/tracing/zipkin"

	pb "comments/api/protobuf_spec/health_check"
)

func (this HealthCheckServiceImpl) HealthCheck (ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	// log.Infof("====== HealthCheck start ======")

	resp := new(pb.HealthCheckResponse)

	return resp, nil
}
