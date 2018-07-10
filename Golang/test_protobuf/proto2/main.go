package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pb2 "learn/Golang/test_protobuf/proto2/protocol"
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

	//Create object and set fields
	person := &pb2.Person{}
	person.Name = proto.String("techie")
	person.Id = proto.Int32(200508628)
	person.Email = proto.String("200508628@qq.com")

	phoneNumbers := []*pb2.Person_PhoneNumber{}
	phoneNumbers = append(phoneNumbers, &pb2.Person_PhoneNumber{Type:pb2.Person_HOME.Enum(), Number:proto.String("1234346")})

	person.Phone = phoneNumbers

	fmt.Printf("person set fields:%v\r\n", person.String())

	//Get object fields
	//if the field is optional, you have to compare to nil to check if the field exist or not
	//Use GetXXX function is safe, if the field dosen't exist, a default value would be returned
	if person.Id != nil {
		fmt.Println("Id is not nil: ", person.GetId())
	} else {
		fmt.Println("Id is nil!")
	}

	//Marshal
	buffer, err := proto.Marshal(person)

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
}

func main() {
	fmt.Println("====== Test protobuf ======")

	testProto2()
}
