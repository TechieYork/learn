
/**********************************
 * Author : Techie
 * Time : 2020-02-29 11:39:02
 **********************************/

package http_handler

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"strconv"

	// log package
	log "github.com/cihub/seelog"
	"google.golang.org/grpc/status"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	tracing "github.com/DarkMetrix/gofra/pkg/tracing/zipkin"

	"github.com/gin-gonic/gin"

	pool "github.com/DarkMetrix/gofra/pkg/grpc-utils/pool"
	config "goods/internal/pkg/config"

	comments "goods/api/protobuf_spec/comments"
	stock "goods/api/protobuf_spec/stock"
)

// URI(for gin use): [GET] -> "/good/:id"
func GET_GOOD_ID(ctx *gin.Context) {
	log.Tracef("====== GET_GOOD_ID start ======")

	// get tracing info
	httpCarrier := opentracing.HTTPHeadersCarrier(ctx.Request.Header)
	parentSpan, err := tracing.GetTracer().Extract(opentracing.HTTPHeaders, httpCarrier)

	grpcCtx := context.Background()

	if err != nil {
		rootSpan := opentracing.StartSpan("GET /good/${id}")
		defer rootSpan.Finish()

		grpcCtx = opentracing.ContextWithSpan(grpcCtx, rootSpan)
	} else {
		childSpan := opentracing.StartSpan("GET /good/${id}",
			opentracing.ChildOf(parentSpan))

		defer childSpan.Finish()

		grpcCtx = opentracing.ContextWithSpan(grpcCtx, childSpan)
	}

	getStockCtx := grpcCtx
	getCommentsCtx := grpcCtx

	// get good id
	id := ctx.Param("id")

	// check params
	goodID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Warnf("parse ID failed! error:%v", err.Error())
		ctx.AbortWithStatusJSON(400, gin.H{"error":err.Error()})
		return
	}

	if goodID > int64(100000) || goodID < int64(1) {
		err = fmt.Errorf("good id should be in range of 1 <= id <= 100000")
		log.Warnf("check param failed! error:%v", err.Error())
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	// get stock
	stockResp, err := getStock(getStockCtx, goodID)

	if err != nil {
		log.Warnf("get stock failed! error:%v", err.Error())
		ctx.AbortWithStatusJSON(400, gin.H{"error":err.Error()})
		return
	}

	// get comments
	commentsResp, err := getComments(getCommentsCtx, goodID)

	if err != nil {
		log.Warnf("get comments failed! error:%v", err.Error())
		ctx.AbortWithStatusJSON(400, gin.H{"error":err.Error()})
		return
	}

	// reply success
	ctx.JSON(200, gin.H{"id": stockResp.GoodID, "stock": stockResp.GoodStock, "comments": commentsResp.Comments})
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

