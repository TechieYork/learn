package main

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"
	jsonpb "github.com/golang/protobuf/jsonpb"

	pb2 "learn/Golang/test_protobuf/pb2json/protocol"
)

func main() {
	m := jsonpb.Marshaler{
		EnumsAsInts: true,
		EmitDefaults: true,
		Indent: "    ",
		OrigName: true,
	}

	person := new(pb2.Person)

	person.StrName = proto.String("techie")
	person.StrEmail = proto.String("test@gmail.com")
	person.Int32Id = proto.Int32(123)

	person.RptPhone = append(person.RptPhone, &pb2.PhoneNumber{
		StrNumber: proto.String("1234566"),
		EnumType: pb2.PhoneType_HOME.Enum(),
	})
	person.RptPhone = append(person.RptPhone, &pb2.PhoneNumber{
		StrNumber: proto.String("999012312"),
		EnumType: pb2.PhoneType_MOBILE.Enum(),
	})

	person.Address = new(pb2.Address)
	person.Address.StrBlock = proto.String("Nanshan street")
	person.Address.ZipCode = proto.String("518000")

	jsonStr, err := m.MarshalToString(person)

	if err != nil {
		fmt.Printf("Marshal failed! error:%v\r\n", err.Error())
		return
	}

	fmt.Printf("Marshal success! json:\r\n%v\r\n", jsonStr)
}

