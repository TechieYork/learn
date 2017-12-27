package main

import (
	"fmt"

	//proto "github.com/golang/protobuf/proto"
	pb3 "learn/test_protobuf/proto3/protocol"
)

func testProto3() {
	fmt.Println("====== Test proto3 ======")

	//Create empty object
	personEmpty := &pb3.Person{}

	fmt.Printf("person of empty:%v\r\n", personEmpty.String())

	//Create object with fields
	personWithFields := &pb3.Person{
		Name: "techie",
		Id: 200508628,
		Email: "200508628@qq.com",
		Phone: []*pb3.Person_PhoneNumber{
			{
				Type: pb3.Person_MOBILE,
				Number: "56774",
			},
			{
				Type: pb3.Person_WORK,
				Number: "456789",
			},
		},
	}

	fmt.Printf("person with fields:%v\r\n", personWithFields.String())
}

func main() {
	fmt.Println("====== Test protobuf ======")

	testProto3()
}
