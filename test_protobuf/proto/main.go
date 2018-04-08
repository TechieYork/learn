package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pb2 "learn/test_protobuf/proto/protocol2"
	pb3 "learn/test_protobuf/proto/protocol3"
)

func testProto() {
	/********************************************************************
	 * Proto2
	 ********************************************************************/
	fmt.Println("====== Test proto ======")

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

	//Marshal
	buffer, err := proto.Marshal(personWithFields)

	if err != nil {
		fmt.Println("Marshal failed! error:", err.Error())
		return
	} else {
		fmt.Println("Marshal successs! buffer:", buffer)
	}

	//Unmarshal
	personUmmarshal := &pb2.Person{}

	err = proto.Unmarshal(buffer, personUmmarshal)

	if err != nil {
		fmt.Println("Unmarshal failed! error:", err.Error())
		return
	} else {
		fmt.Println("Unmarshal success! person:", personUmmarshal.String())
	}

	/********************************************************************
	 * Proto3
	 ********************************************************************/
	fmt.Println("====== Test proto3 ======")

	//Create empty object
	personInfoEmpty := &pb3.PersonInfo{}

	fmt.Printf("person of empty:%v\r\n", personInfoEmpty.String())

	//Create object with fields
	personInfoWithFields := &pb3.PersonInfo{
		Name: "techie",
		Id: 200508628,
		Email: "200508628@qq.com",
		Phone: []*pb3.PersonInfo_PhoneNumber{
			{
				Type: pb3.PersonInfo_MOBILE,
				Number: "56774",
			},
			{
				Type: pb3.PersonInfo_WORK,
				Number: "456789",
			},
		},
	}

	fmt.Printf("person with fields:%v\r\n", personInfoWithFields.String())

	//Get object fields
	fmt.Println("GetId() return: ", personInfoWithFields.GetId())
	fmt.Println("Id return: ", personInfoWithFields.Id)

	//Marshal
	bufferInfo, err := proto.Marshal(personInfoWithFields)

	if err != nil {
		fmt.Println("Marshal failed! error:", err.Error())
		return
	} else {
		fmt.Println("Marshal successs! buffer:", bufferInfo)
	}

	//Unmarshal
	personInfoUmmarshal := &pb3.PersonInfo{}

	err = proto.Unmarshal(bufferInfo, personInfoUmmarshal)

	if err != nil {
		fmt.Println("Unmarshal failed! error:", err.Error())
		return
	} else {
		fmt.Println("Unmarshal success! person:", personInfoUmmarshal.String())
	}
}

func main() {
	fmt.Println("====== Test protobuf ======")

	testProto()
}
