
/**********************************
 * Author : Techie
 * Time : 2020-02-27 20:44:19
 **********************************/

package http_handler

import (
	"fmt"
	// log package
	log "github.com/cihub/seelog"
	"strconv"
	"user/pkg/user"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	// tracing "github.com/DarkMetrix/gofra/pkg/tracing/jaeger"

	"github.com/gin-gonic/gin"
)

// URI(for gin use): [GET] -> "/user/:id"
func GET_USER_ID(ctx *gin.Context) {
	log.Tracef("====== GET_USER_ID start ======")

	id := ctx.Param("id")

	// check params
	userID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Warnf("parse ID failed! error:%v", err.Error())
		ctx.AbortWithStatusJSON(400, gin.H{"error":err.Error()})
		return
	}

	if userID > int64(100000000) || userID < int64(1) {
		err = fmt.Errorf("user id should be in range of 1 <= id <= 100000000")
		log.Warnf("check param failed! error:%v", err.Error())
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	// reply success
	// since it's a fake server, just used to test istio features, reply with a fixed json
	var resp user.UserInfo

	resp.ID = userID
	resp.Name = fmt.Sprintf("Foo-%v", userID)
	resp.Pic = fmt.Sprintf("http://xxx.zzz/pic-%v.jpg", userID)

	ctx.JSON(200, gin.H{"user_info":resp})
}

