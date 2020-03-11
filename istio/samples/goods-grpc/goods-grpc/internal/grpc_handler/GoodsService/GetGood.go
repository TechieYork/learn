
/**********************************
 * Author : Techie
 * Time : 2020-03-11 10:34:04
 **********************************/

package GoodsService 

import (
	"context"
	"fmt"

	"google.golang.org/grpc/status"

	// log package
	log "github.com/cihub/seelog"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	// opentracing "github.com/opentracing/opentracing-go"

	pool "github.com/DarkMetrix/gofra/pkg/grpc-utils/pool"
	config "goods-grpc/internal/pkg/config"

	pb "goods-grpc/api/protobuf_spec/goods"

	comments "goods-grpc/api/protobuf_spec/comments"
	stock "goods-grpc/api/protobuf_spec/stock"
)

func (this GoodsServiceImpl) GetGood (ctx context.Context, req *pb.GetGoodRequest) (*pb.GetGoodResponse, error) {
	// log.Infof("====== GetGood start ======")

	// check params
	if req.GetGoodID() > int64(100000) || req.GetGoodID() < int64(1) {
		err := fmt.Errorf("good id should be in range of 1 <= id <= 100000")
		log.Warnf("check param failed! error:%v", err.Error())
		return nil, err
	}

	resp := new(pb.GetGoodResponse)

	getStockCtx := ctx
	getCommentsCtx := ctx

	// get stock
	stockResp, err := getStock(getStockCtx, req.GetGoodID())

	if err != nil {
		log.Warnf("get stock failed! error:%v", err.Error())
		return nil, err
	}

	// get comments
	commentsResp, err := getComments(getCommentsCtx, req.GetGoodID())

	if err != nil {
		log.Warnf("get comments failed! error:%v", err.Error())
		return nil, err
	}

	resp.GoodID = req.GetGoodID()
	resp.GoodStock = stockResp.GetGoodStock()
	resp.Comments = commentsResp.GetComments()

	return resp, nil
}

func getStock(ctx context.Context, goodID int64) (*stock.GetStockResponse, error) {
	// init request
	req := &stock.GetStockRequest{}
	req.GoodID = goodID

	// get connection
	conn, err := pool.GetConnectionPool().GetConnection(context.Background(), config.GetConfig().Stock.Addr)

	if err != nil {
		log.Warnf("pool.GetConnection failed! error:%v", err.Error())
		return nil, err
	}

	// call
	c := stock.NewStockServiceClient(conn)

	resp, err := c.GetStock(ctx, req)

	if err != nil {
		stat, ok := status.FromError(err)

		if ok {
			log.Warnf("Get user info batch failed! stat code:%v, stat error:%v", stat.Code(), stat.Message())
			return nil, err
		} else {
			log.Warnf("Get user info batch failed! error:%v", err.Error())
			return nil, err
		}
	}

	return resp, nil
}

func getComments(ctx context.Context, goodID int64) (*comments.GetCommentsByGoodIDResponse, error) {
	// init request
	req := &comments.GetCommentsByGoodIDRequest{}
	req.GoodID = goodID

	// get connection
	conn, err := pool.GetConnectionPool().GetConnection(context.Background(), config.GetConfig().Comments.Addr)

	if err != nil {
		log.Warnf("pool.GetConnection failed! error:%v", err.Error())
		return nil, err
	}

	// call
	c := comments.NewCommentsServiceClient(conn)

	resp, err := c.GetCommentsByGoooID(ctx, req)

	if err != nil {
		stat, ok := status.FromError(err)

		if ok {
			log.Warnf("Get comments by good id failed! stat code:%v, stat error:%v", stat.Code(), stat.Message())
			return nil, err
		} else {
			log.Warnf("Get comments by good id failed! error:%v", err.Error())
			return nil, err
		}
	}

	return resp, nil
}

