
/**********************************
 * Author : Techie
 * Time : 2020-02-29 11:39:02
 **********************************/

package http_handler

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	// log package
	log "github.com/cihub/seelog"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	"github.com/opentracing/opentracing-go"

	"github.com/gin-gonic/gin"

	config "goods/internal/pkg/config"
)

// URI(for gin use): [GET] -> "/good/:id"
func GET_GOOD_ID(ctx *gin.Context) {
	log.Tracef("====== GET_GOOD_ID start ======")

	// get tracing info
	httpCarrier := opentracing.HTTPHeadersCarrier(ctx.Request.Header)
	parentSpan, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, httpCarrier)

	if err != nil {
		rootSpan := opentracing.StartSpan("GET /good/${id}")
		defer rootSpan.Finish()

		parentSpan = rootSpan.Context()
	}

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
	stockResp, err := getStock(parentSpan, goodID)

	if err != nil {
		log.Warnf("get stock failed! error:%v", err.Error())
		ctx.AbortWithStatusJSON(400, gin.H{"error":err.Error()})
		return
	}

	// get comments
	commentsResp, err := getComments(parentSpan, goodID)

	if err != nil {
		log.Warnf("get comments failed! error:%v", err.Error())
		ctx.AbortWithStatusJSON(400, gin.H{"error":err.Error()})
		return
	}

	// reply success
	ctx.JSON(200, gin.H{"id": stockResp.ID, "stock": stockResp.Stock, "comments": commentsResp.Comments})
}

type StockResponse struct {
	ID int64 `json:"id"`
	Stock int64 `json:"stock"`
}

type CommentsInfo struct {
	UserID int64 `json:"user_id"`
	UserName string `json:"user_name"`
	UserPic string `json:"user_pic"`
	Comment string `json:"comment"`
}

type CommentsResponse struct {
	Comments []CommentsInfo
}

func getStock(parentSpan opentracing.SpanContext, goodID int64) (*StockResponse, error) {
	// get stock request
	// init http client & send request
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%v/stock/good/%v", config.GetConfig().Stock.HTTPAddr, goodID),
		nil)

	if err != nil {
		return nil, err
	}

	// inject parent span
	opentracing.GlobalTracer().Inject(
		parentSpan,
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header))

	// send request
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		log.Warnf("error:%v", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Warnf("send request to stock failed! status code:%v", resp.StatusCode)
			return nil, fmt.Errorf("send request to stock failed! status code:%v", resp.StatusCode)
		} else {
			log.Warnf("send request to stock failed! status code:%v, body:%v", resp.StatusCode, string(body))
			return nil, fmt.Errorf("send request to stock failed! status code:%v, body:%v", resp.StatusCode, string(body))
		}
	}

	// read response & unmarshal to user batch response
	jsonRespBuf, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Warnf("error:%v", err.Error())
		return nil, err
	}

	var stockResp StockResponse
	err = json.Unmarshal(jsonRespBuf, &stockResp)

	if err != nil {
		log.Warnf("error:%v", err.Error())
		return nil, err
	}

	return &stockResp, nil
}

func getComments(parentSpan opentracing.SpanContext, goodID int64) (*CommentsResponse, error) {
	// get stock request
	// init http client & send request
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%v/comments/good/%v", config.GetConfig().Comments.HTTPAddr, goodID),
		nil)

	if err != nil {
		return nil, err
	}

	// inject parent span
	opentracing.GlobalTracer().Inject(
		parentSpan,
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header))

	// send request
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		log.Warnf("error:%v", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Warnf("send request to comments failed! status code:%v", resp.StatusCode)
			return nil, fmt.Errorf("send request to comments failed! status code:%v", resp.StatusCode)
		} else {
			log.Warnf("send request to comments failed! status code:%v, body:%v", resp.StatusCode, string(body))
			return nil, fmt.Errorf("send request to comments failed! status code:%v, body:%v", resp.StatusCode, string(body))
		}
	}

	// read response & unmarshal to user batch response
	jsonRespBuf, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Warnf("error:%v", err.Error())
		return nil, err
	}

	var commentsResp CommentsResponse
	err = json.Unmarshal(jsonRespBuf, &commentsResp)

	if err != nil {
		log.Warnf("error:%v", err.Error())
		return nil, err
	}

	return &commentsResp, nil
}

