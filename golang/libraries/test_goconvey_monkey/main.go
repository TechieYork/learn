package main

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

func CheckParamPrefix(content string) error {
	return validation.Validate(content, validation.NilOrNotEmpty, validation.Length(1, 10), validation.Match(regexp.MustCompile("^test_.*")))
}

func main() {
	fmt.Println(CheckParamPrefix(""))
}
