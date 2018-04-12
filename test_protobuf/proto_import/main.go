package main

import (
	"github.com/gogo/protobuf/proto"
	person "learn/test_protobuf/proto_import/protocol/proto_import"
	student "learn/test_protobuf/proto_import/protocol/proto_main"
	"fmt"
)
func main() {
	s := &student.Student{
		Person: &person.Person{
			Name: proto.String("test"),
			Id: proto.Int32(1234),
			Email: proto.String("asdf"),
		},
		School: proto.String("High"),
		Grade: proto.String("Grade 1"),
	}

	fmt.Println("Student:", s)
}
