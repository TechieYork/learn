package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "learn/test_grpc/basic/proto"
)

func main() {
	// 连接
	conn, err := grpc.Dial(":58888", grpc.WithInsecure())

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewUserServiceClient(conn)

	// 调用方法
	req := new(pb.AddUserRequest)
	req.Session = new(pb.Session)
	req.User = new(pb.User)

	req.Session.Seq = "12345678"
	req.User.Name = "techieliu"
	req.User.Sex = 1

	resp, err := c.AddUser(context.Background(), req)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.String())
}
