package main

import (
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "learn/test_grpc/basic/proto"
)

// Implement UserService interface
type UserService struct{}

// UserService ...
var userService = UserService{}

func (h UserService) AddUser (ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	resp := new(pb.AddUserResponse)
	resp.Session = new(pb.Session)
	resp.Result = new(pb.Result)

	//Set session
	resp.Session = req.Session

	//Set result
	resp.GetResult().Code = 0
	resp.GetResult().Message = "Success"

	return resp, nil
}

func main() {
	// start
	fmt.Println("====== Test grpc server ======")

	listen, err := net.Listen("tcp", ":58888")
	if err != nil {
		fmt.Println("failed to listen: ", err.Error())
		return
	}

	var opts []grpc.ServerOption

	// interceptor
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("====== Enter interceptor ======")
		fmt.Printf("context:\r\n%v\r\n", ctx)
		fmt.Printf("method:\r\n%v\r\n", info.FullMethod)
		fmt.Printf("server:\r\n%v\r\n", info.Server)

		// 继续处理请求
		fmt.Println("====== interceptor handle ======")
		fmt.Printf("req:%v\r\n", req)

		resp, err = handler(ctx, req)

		if err != nil {
			fmt.Printf("handle failed! error:%v\r\n", err)
		} else {
			fmt.Printf("handle success! resp:%v\r\n", resp)
		}

		fmt.Println("====== Leave interceptor ======")

		return resp, err
	}

	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	// init grpc Server
	s := grpc.NewServer(opts...)

	// register UserService
	pb.RegisterUserServiceServer(s, userService)

	fmt.Println("Listen on :58888")

	s.Serve(listen)
}
