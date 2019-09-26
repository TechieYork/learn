package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "learn/Golang/test_grpc/basic/proto"
)

func main() {
	// interceptor
	var interceptor grpc.UnaryClientInterceptor
	interceptor = func(ctx context.Context, method string, req, reply interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Println("====== Enter interceptor ======")
		fmt.Printf("context:\r\n%v\r\n", ctx)
		fmt.Printf("method:\r\n%v\r\n", method)
		fmt.Printf("options:\r\n%v\r\n", opts)

		//继续处理请求
		err := invoker(ctx, method, req, reply, cc, opts...)

		if err != nil {
			fmt.Printf("invoke failed! error:%v\r\n", err)
		} else {
			fmt.Printf("invoke success! resp:%v\r\n", reply)
		}

		fmt.Println("====== Leave interceptor ======")
		return nil
	}

	// Dial remote server
	conn, err := grpc.Dial(":58888", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	// New client
	c := pb.NewUserServiceClient(conn)

	// RPC call
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

	time.Sleep(time.Second * 1)
}
