
/**********************************
 * Author : Techie
 * Time : 2020-03-08 20:46:22
 **********************************/

package UserService 

import (
	"context"
	"fmt"

	// log package
	log "github.com/cihub/seelog"

	// monitor package
	// monitor "github.com/DarkMetrix/gofra/pkg/monitor/statsd"

	// tracing package
	// opentracing "github.com/opentracing/opentracing-go"

	pb "user/api/protobuf_spec/user"
)

func (this UserServiceImpl) GetUserBatch (ctx context.Context, req *pb.GetUserBatchRequest) (*pb.GetUserBatchResponse, error) {
	// log.Infof("====== GetUserBatch start ======")

	// check param
	err := checkGetUserBatchParam(req)

	if err != nil {
		log.Warnf("check param failed! error:%v", err.Error())
		return nil, err
	}

	// reply success
	// since it's a fake server, just used to test istio features, reply with fixed data
	resp := new(pb.GetUserBatchResponse)

	for _, id := range req.GetIds() {
		newUser := &pb.UserInfo{
			Id: id,
			Name: fmt.Sprintf("Foo-%v", id),
			Pic: fmt.Sprintf("http://xxx.zzz/pic-%v.jpg", id),
		}

		resp.UserInfos = append(resp.UserInfos, newUser)
	}

	return resp, nil
}

func checkGetUserBatchParam(req *pb.GetUserBatchRequest) error {
	for _, id := range req.Ids {
		if id > int64(100000000) || id < int64(1) {
			return fmt.Errorf("user id should be in range of 1 <= id <= 100000000")
		}
	}

	return nil
}
