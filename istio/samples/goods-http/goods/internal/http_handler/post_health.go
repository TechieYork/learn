
/**********************************
 * Author : Techie
 * Time : 2020-02-29 11:34:44
 **********************************/

package http_handler

import (
	// log package
	log "github.com/cihub/seelog"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	// tracing "github.com/DarkMetrix/gofra/pkg/tracing/jaeger"

	"github.com/gin-gonic/gin"
)

// URI(for gin use): [POST] -> "/health"
func POST_HEALTH(ctx *gin.Context) {
	log.Tracef("====== POST_HEALTH start ======")

	/*
	// parse request
	// TODO: Bind json to request
	var req xxx

	err := ctx.BindJSON(&req)

	if err != nil {
		log.Warnf("ctx.BindJSON failed! error:%v", err.Error())
		ctx.AbortWithStatusJSON(520, gin.H{"ret":-1, "msg":"Bad json body!"})
		return
	}

	// check params
	// TODO: Check params
	err = checkPOST_HEALTHParams(&req)

	if err != nil {
		log.Warnf("checkPOST_HEALTHParams failed! error:%v", err.Error())
		ctx.AbortWithStatusJSON(520, gin.H{"ret":-1, "msg":fmt.Sprintf("Param invalid! error:%v", err.Error())})
		return
	}
	*/

	// reply success
	ctx.JSON(200, gin.H{"ret":0, "msg":"success"})
}

/*
// TODO: Implement checkPOST_HEALTHParams function
func checkPOST_HEALTHParams(req *xxx) error {
	return nil
}
*/
