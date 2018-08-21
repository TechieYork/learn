package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pb3 "learn/Golang/test_protobuf/proto3/protocol"
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

	//Create object and set fields
	person := &pb3.Person{}
	person.Name = "techie"
	person.Id = 200508628
	person.Email = "200508628@qq.com"

	phoneNumbers := []*pb3.Person_PhoneNumber{}
	phoneNumbers = append(phoneNumbers, &pb3.Person_PhoneNumber{Type:pb3.Person_HOME, Number:"1234346"})

	person.Phone = phoneNumbers

	fmt.Printf("person set fields:%v\r\n", person.String())

	//Get object fields
	fmt.Println("GetId() return: ", person.GetId())
	fmt.Println("Id return: ", person.Id)

	//Marshal
	buffer, err := proto.Marshal(person)

	if err != nil {
		fmt.Println("Marshal failed! error:", err.Error())
		return
	} else {
		fmt.Println("Marshal successs! buffer:", buffer)
	}

	//Unmarshal
	personUmmarshal := &pb3.Person{}

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

	testProto3()
}
