package main

import "fmt"

func funcNormal() {
	//Functions without input params and returns
}

func funcWithParams(a, b int) {
	//Functions with input params, but no returns
}

func funcWithReturns() error {
	//Functions without input params, but with returns
	return nil
}

func funcWithVariadic(a, b string, params ...string) {
	//Functions with variadic params
	for i, p := range params {
		fmt.Printf("Params[%v]: %v\r\n", i, p)
	}
}

func funcWithMultiReturns() (int, error) {
	return 0, nil
}

func funcWithMultiReturnsAndName() (code int, ret error) {
	code = 1
	ret = nil
	return
}

func main() {
	funcNormal()
	funcWithParams(1, 2)
	ret := funcWithReturns()
	fmt.Println(ret)

	funcWithVariadic("a", "b", "c", "d", "e")
	funcWithVariadic("a", "b", []string{"c", "d", "e"}...)

	code, ret := funcWithMultiReturns()
	fmt.Println(code, ret)
	code, ret = funcWithMultiReturnsAndName()
	fmt.Println(code, ret)
}
