package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.MatchString("11111.*", "1111"))
}
