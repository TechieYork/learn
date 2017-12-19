package main

import (
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "learn/test_grpc/basic/proto"
)

// 定义UserService并实现约定的接口
type UserService struct{}

// UserService ...
var userService = UserService{}

func (h UserService) AddUser (ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	fmt.Println(req.String())

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
	listen, err := net.Listen("tcp", ":58888")
	if err != nil {
		fmt.Println("failed to listen: ", err.Error())
		return
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterUserServiceServer(s, userService)

	fmt.Println("Listen on :58888")

	s.Serve(listen)
}
