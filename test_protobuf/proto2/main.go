package main

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"
	pb2 "learn/test_protobuf/proto2/protocol"
)

func testProto2() {
	fmt.Println("====== Test proto2 ======")

	//Create empty object
	personEmpty := &pb2.Person{}
	
	fmt.Println("person of empty:%v", personEmpty.String())

	//Create object with fields
	personWithFields := &pb2.Person{
		Name: proto.String("techie"),
	}

	fmt.Println("person with fields:%v", personWithFields.String())
}

func testProto3() {
	fmt.Println("====== Test proto3 ======")
}

func main() {
	fmt.Println("====== Test protobuf ======")

	testProto2()
	testProto3()
}
