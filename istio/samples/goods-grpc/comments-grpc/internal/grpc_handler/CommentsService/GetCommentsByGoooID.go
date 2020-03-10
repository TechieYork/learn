
/**********************************
 * Author : Techie
 * Time : 2020-03-09 12:01:15
 **********************************/

package CommentsService 

import (
	"fmt"
	"context"

	// log package
	log "github.com/cihub/seelog"
	"google.golang.org/grpc/status"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	// opentracing "github.com/opentracing/opentracing-go"

	pool "github.com/DarkMetrix/gofra/pkg/grpc-utils/pool"
	config "comments/internal/pkg/config"

	pb "comments/api/protobuf_spec/comments"
	user "comments/api/protobuf_spec/user"
)

func (this CommentsServiceImpl) GetCommentsByGoooID (ctx context.Context, req *pb.GetCommentsByGoodIDRequest) (*pb.GetCommentsByGoodIDResponse, error) {
	// log.Infof("====== GetCommentsByGoooID start ======")

	// check param
	if req.GoodID > int64(100000) || req.GoodID < int64(1) {
		err := fmt.Errorf("good id should be in range of 1 <= id <= 100000")
		log.Warnf("check param failed! error:%v", err.Error())
		return nil, err
	}

	// since it's a fake server, just used to test istio features, reply with fixed data
	resp := new(pb.GetCommentsByGoodIDResponse)
	var userIDs []int64

	for i := 1; i < 6; i++ {
		userIDs = append(userIDs, int64(i))
		resp.Comments = append(resp.Comments, &pb.CommentInfo{
			UserID:   int64(i),
			UserName: "",
			UserPic:  "",
			Comments: fmt.Sprintf("This is comment from user id:%v", int64(i)),
		})	
	}

	// get user infos
	userInfos, err := getUserInfoBatch(ctx, userIDs)

	if err != nil {
		log.Warnf("Get user info batch failed! error:%v", err.Error())
		return nil, err
	}

	// fill user info
	userInfoMap := make(map[int64]*user.UserInfo)

	for _, info := range userInfos {
		userInfoMap[info.Id] = info
	}

	for i := 0; i < len(resp.Comments); i++ {
		user, ok := userInfoMap[resp.Comments[i].UserID]

		if ok {
			resp.Comments[i].UserName = user.Name
			resp.Comments[i].UserPic = user.Pic
		} else {
			resp.Comments[i].UserName = "Unknown"
			resp.Comments[i].UserPic = "http://xxx.zzz/default.jpg"
		}
	}

	return resp, nil
}

func getUserInfoBatch(ctx context.Context, ids []int64) ([]*user.UserInfo, error){
	// init request
	req := &user.GetUserBatchRequest{}
	req.Ids = ids

	// get connection
	conn, err := pool.GetConnectionPool().GetConnection(context.Background(), config.GetConfig().User.Addr)

	if err != nil {
		log.Warnf("pool.GetConnection failed! error:%v", err.Error())
		return nil, err
	}

	// call
	c := user.NewUserServiceClient(conn)

	resp, err := c.GetUserBatch(ctx, req)

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

	return resp.UserInfos, nil
}
