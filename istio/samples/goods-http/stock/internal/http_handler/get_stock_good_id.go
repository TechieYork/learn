
/**********************************
 * Author : Techie
 * Time : 2020-02-29 11:04:39
 **********************************/

package http_handler

import (
	"fmt"
	"strconv"

	// log package
	log "github.com/cihub/seelog"

	"github.com/gin-gonic/gin"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	// opentracing "github.com/opentracing/opentracing-go"

)

// URI(for gin use): [GET] -> "/stock/good/:id"
func GET_STOCK_GOOD_ID(ctx *gin.Context) {
	log.Tracef("====== GET_STOCK_GOOD_ID start ======")

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

	// reply success
	// since it's a fake server, just used to test istio features, reply with fixed data
	ctx.JSON(200, gin.H{"id":goodID, "stock":500})
}
