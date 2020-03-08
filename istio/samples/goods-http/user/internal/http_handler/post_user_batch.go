
/**********************************
 * Author : Techie
 * Time : 2020-02-27 20:44:43
 **********************************/

package http_handler

import (
	"fmt"
	"user/pkg/user"

	// log package
	log "github.com/cihub/seelog"
	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	// tracing "github.com/DarkMetrix/gofra/pkg/tracing/jaeger"

	"github.com/gin-gonic/gin"
)

type UserBatchReq struct {
	UserIDs []int64 `json:"user_ids"`
}

// URI(for gin use): [POST] -> "/user/batch"
func POST_USER_BATCH(ctx *gin.Context) {
	log.Tracef("====== POST_USER_BATCH start ======")

	// parse request
	var req UserBatchReq

	err := ctx.BindJSON(&req)

	if err != nil {
		log.Warnf("ctx.BindJSON failed! error:%v", err.Error())
		ctx.AbortWithStatusJSON(400, gin.H{"error":err.Error()})
		return
	}

	// check params
	// TODO: Check params
	err = checkPOST_USER_BATCHParams(&req)

	if err != nil {
		log.Warnf("check param failed! error:%v", err.Error())
		ctx.AbortWithStatusJSON(520, gin.H{"ret":-1, "msg":fmt.Sprintf("Param invalid! error:%v", err.Error())})
		return
	}

	// get users info
	// since it's a fake server, just used to test istio features, reply with fixed data
	var resp []user.UserInfo

	for _, v := range req.UserIDs {
		resp = append(resp, user.UserInfo{
			ID:   v,
			Name: fmt.Sprintf("Foo-%v", v),
			Pic:  fmt.Sprintf("http://xxx.zzz/pic-%v.jpg", v),
		})
	}

	// reply success
	ctx.JSON(200, gin.H{"user_info_list":resp})
}

// TODO: Implement checkPOST_USER_BATCHParams function
func checkPOST_USER_BATCHParams(req *UserBatchReq) error {
	for _, v := range req.UserIDs {
		if v > int64(100000000) || v < int64(1) {
			return fmt.Errorf("user id should in the range of 1 <= id <= 100000000")
		}
	}

	return nil
}
