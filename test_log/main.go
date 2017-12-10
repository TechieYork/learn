package main

import (
	"fmt"

	logrus "github.com/sirupsen/logrus"
	//blog4go "github.com/YoungPioneers/blog4go"
	//seelog "github.com/cihub/seelog"
)

func testLogrus() {
	fmt.Println("====== Test logrus ======")
	logrus.Infoln("test logrus")
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func testBlog4go() {
	fmt.Println("====== Test blog4go ======")
}

func testSeelog() {
	fmt.Println("====== Test seelog ======")
}

func main() {
	fmt.Println("====== Test log ======")

	testSeelog()
	testBlog4go()
	testLogrus()
}
