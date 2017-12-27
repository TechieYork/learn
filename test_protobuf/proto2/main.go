package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pb2 "learn/test_protobuf/proto2/protocol"
)

func testProto2() {
	fmt.Println("====== Test proto2 ======")

	//Create empty object
	personEmpty := &pb2.Person{}
	
	fmt.Printf("person of empty:%v\r\n", personEmpty.String())

	//Create object with fields
	personWithFields := &pb2.Person{
		Name: proto.String("techie"),
		Id: proto.Int32(200508628),
		Email: proto.String("200508628@qq.com"),
		Phone: []*pb2.Person_PhoneNumber{
			{
				Type: pb2.Person_MOBILE.Enum(),
				Number: proto.String("56774"),
			},
			{
				Type: pb2.Person_WORK.Enum(),
				Number: proto.String("456789"),
			},
		},
	}

	fmt.Printf("person with fields:%v\r\n", personWithFields.String())
}

func main() {
	fmt.Println("====== Test protobuf ======")

	testProto2()
}
