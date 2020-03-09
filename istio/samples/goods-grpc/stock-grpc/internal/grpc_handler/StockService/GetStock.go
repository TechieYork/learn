
/**********************************
 * Author : Techie
 * Time : 2020-03-09 11:51:46
 **********************************/

package StockService 

import (
	"context"
	"fmt"

	// log package
	log "github.com/cihub/seelog"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	// tracing "github.com/DarkMetrix/gofra/pkg/tracing/zipkin"

	pb "stock/api/protobuf_spec/stock"
)

func (this StockServiceImpl) GetStock (ctx context.Context, req *pb.GetStockRequest) (*pb.GetStockResponse, error) {
	// log.Infof("====== GetStock start ======")

	// check param
	if req.GoodID > int64(100000) || req.GoodID < int64(1) {
		err := fmt.Errorf("good id should be in range of 1 <= id <= 100000")
		log.Warnf("check param failed! error:%v", err.Error())
		return nil, err
	}

	// reply success
	// since it's a fake server, just used to test istio features, reply with fixed data
	resp := new(pb.GetStockResponse)

	resp.GoodID = req.GoodID
	resp.GoodStock = 999

	return resp, nil
}
