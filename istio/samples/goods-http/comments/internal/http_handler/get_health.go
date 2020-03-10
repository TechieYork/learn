
/**********************************
 * Author : Techie
 * Time : 2020-02-28 11:59:16
 **********************************/

package http_handler

import (
	// log package
	log "github.com/cihub/seelog"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	// opentracing "github.com/opentracing/opentracing-go"

	"github.com/gin-gonic/gin"
)

// URI(for gin use): [POST] -> "/health"
func GET_HEALTH(ctx *gin.Context) {
	log.Tracef("====== GET_HEALTH start ======")

	// reply success
	ctx.JSON(200, gin.H{"msg":"success"})
}

