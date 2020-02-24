
/**********************************
 * Author : Techie
 * Time : 2020-02-25 00:13:30
 **********************************/

package http_handler

import (
	// log package
	log "github.com/cihub/seelog"
	"strconv"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	// tracing "github.com/DarkMetrix/gofra/pkg/tracing/jaeger"

	"github.com/gin-gonic/gin"
)

// URI(for gin use): [GET] -> "/details/:id"
func GET_DETAILS_ID(ctx *gin.Context) {
	log.Tracef("====== GET_DETAILS_ID start ======")

	// parse request
	id := ctx.Param("id")

	// check params
	numID, err := strconv.ParseInt(id, 10, 64)

	if err != nil{
		log.Warnf("invalid product id: %v, error:%v", id, err.Error())
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	// reply success
	ctx.JSON(200, gin.H{
		"id" : numID,
		"author": "William Shakespeare",
		"year": 1595,
		"type" : "paperback",
		"pages" : 200,
		"publisher" : "PublisherA",
		"language" : "English",
		"ISBN-10" : "1234567890",
		"ISBN-13" : "123-1234567890",
	})
}

