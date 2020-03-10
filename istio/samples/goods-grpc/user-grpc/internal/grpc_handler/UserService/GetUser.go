
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

func (this UserServiceImpl) GetUser (ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// log.Infof("====== GetUser start ======")

	// check param
	if req.Id > int64(100000000) || req.Id < int64(1) {
		err := fmt.Errorf("user id should be in range of 1 <= id <= 100000000")
		log.Warnf("check param failed! error:%v", err.Error())
		return nil, err
	}

	// reply success
	// since it's a fake server, just used to test istio features, reply with fixed data
	resp := new(pb.GetUserResponse)
	resp.UserInfo = &pb.UserInfo{}

	resp.GetUserInfo().Id = req.GetId()
	resp.GetUserInfo().Name = fmt.Sprintf("Foo-%v", req.GetId())
	resp.GetUserInfo().Pic = fmt.Sprintf("http://xxx.zzz/pic-%v.jpg", req.GetId())

	return resp, nil
}
