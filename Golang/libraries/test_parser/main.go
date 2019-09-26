package main

import (
	"fmt"

	"github.com/tallstoat/pbparser"
)

func main() {
	fmt.Println("====== Test protobuf parser ======")

	file := "./user.proto"

	// invoke ParseFile() API to parse the file
	pf, err := pbparser.ParseFile(file)
	if err != nil {
		fmt.Printf("Unable to parse proto file: %v \n", err)
		return
	}

	// print attributes of the returned datastructure
	fmt.Printf("PackageName: %v, Syntax: %v\r\n",
		pf.PackageName, pf.Syntax)

	fmt.Printf("Dependencies: %v\r\n", pf.Dependencies)

	fmt.Println("====== Print service defination ======")
	for _, elem := range pf.Services {
		fmt.Printf("  Name:%v\r\n  Doc:%v\r\n  Opt:%v\r\n  QName:%v\r\n",
			elem.Name, elem.Documentation, elem.Options, elem.QualifiedName)

		fmt.Println("====== Print rpc defination ======")
		for _, rpc := range elem.RPCs {
			fmt.Printf("  Name:%v\r\n  Doc:%v\r\n  Opt:%v\r\n  ReqType:%v\r\n  RspType:%v\r\n",
				rpc.Name, rpc.Documentation, rpc.Options, rpc.RequestType, rpc.ResponseType)
		}
	}
}
