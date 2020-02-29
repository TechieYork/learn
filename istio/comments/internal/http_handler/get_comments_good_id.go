
/**********************************
 * Author : Techie
 * Time : 2020-02-28 12:00:32
 **********************************/

package http_handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	// log package
	log "github.com/cihub/seelog"

	"github.com/gin-gonic/gin"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	// tracing "github.com/DarkMetrix/gofra/pkg/tracing/jaeger"

	"comments/pkg/comments"
	"comments/internal/pkg/config"
)

// URI(for gin use): [GET] -> "/comments/good/:id"
func GET_COMMENTS_GOOD_ID(ctx *gin.Context) {
	log.Tracef("====== GET_COMMENTS_GOOD_ID start ======")

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

	// since it's a fake server, just used to test istio features, reply with fixed data
	var goodComments []*comments.CommentsInfo

	for i := 1; i < 6; i++ {
		goodComments = append(goodComments, &comments.CommentsInfo{
			UserID: int64(i),
			Comment: fmt.Sprintf("This is comment from user id:%v", i),
		})
	}

	// send request to user to get user name and pic
	goodComments, err = getCommentsUserInfo(goodComments)

	if err != nil {
		log.Warnf("get comments user info failed! error:%v", err.Error())
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	// reply success
	ctx.JSON(200, gin.H{"comments": goodComments})
}

type UserBatchReq struct {
	UserIDs []int64 `json:"user_ids"`
}

type UserInfo struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Pic string `json:"pic"`
}

type UserBatchResp struct {
	UserInfos []UserInfo `json:"user_info_list"`
}

func getCommentsUserInfo(goodComments []*comments.CommentsInfo) ([]*comments.CommentsInfo, error) {
	// get all user id
	var userIDs []int64

	for _, v := range goodComments {
		userIDs = append(userIDs, v.UserID)
	}

	// get user info from remote user service
	userBatchResp, err := batchGetUserInfo(userIDs)

	if err != nil {
		return nil, err
	}

	// set user name and pic to comments
	userInfoMap := make(map[int64]UserInfo)

	for _, v := range userBatchResp.UserInfos {
		userInfoMap[v.ID] = v
	}

	for i := 0; i < len(goodComments); i++ {
		user, ok := userInfoMap[goodComments[i].UserID]

		if ok {
			goodComments[i].UserName = user.Name
			goodComments[i].UserPic = user.Pic
		} else {
			goodComments[i].UserName = "Unknown"
			goodComments[i].UserPic = "http://xxx.zzz/default.jpg"
		}
	}

	return goodComments, nil
}

func batchGetUserInfo(userIDs []int64) (*UserBatchResp, error) {
	// get user batch request
	var userBatchReq UserBatchReq

	for _, v := range userIDs{
		userBatchReq.UserIDs = append(userBatchReq.UserIDs, v)
	}

	// json marshal
	jsonBytes , err := json.Marshal(userBatchReq)

	if err != nil {
		log.Warnf("error:%v", err.Error())
		return nil, err
	}

	jsonBuf := bytes.NewBuffer(jsonBytes)

	// init http client & send request
	httpClient := &http.Client{}

	resp, err := httpClient.Post(fmt.Sprintf("%v/user/batch", config.GetConfig().User.HTTPAddr),
		"application/json",
		jsonBuf)

	if err != nil {
		log.Warnf("error:%v", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Warnf("send request to user failed! status code:%v", resp.StatusCode)
			return nil, fmt.Errorf("send request to user failed! status code:%v", resp.StatusCode)
		} else {
			log.Warnf("send request to user failed! status code:%v, body:%v", resp.StatusCode, string(body))
			return nil, fmt.Errorf("send request to user failed! status code:%v, body:%v", resp.StatusCode, string(body))
		}
	}

	// read response & unmarshal to user batch response
	jsonRespBuf, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Warnf("error:%v", err.Error())
		return nil, err
	}

	var userBatchResp UserBatchResp
	err = json.Unmarshal(jsonRespBuf, &userBatchResp)

	if err != nil {
		log.Warnf("error:%v", err.Error())
		return nil, err
	}

	return &userBatchResp, nil
}

