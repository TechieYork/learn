package main

import (
	"fmt"
	"github.com/jinzhu/now"
)
func main() {
	fmt.Println(now.BeginningOfDay())
	fmt.Println(now.EndOfDay())
	fmt.Println(now.BeginningOfMonth())
	fmt.Println(now.EndOfMonth())
}
