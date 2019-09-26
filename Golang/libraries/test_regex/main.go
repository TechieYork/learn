package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.MatchString("^person_.*", "person_sdfsdfdsf"))
}
