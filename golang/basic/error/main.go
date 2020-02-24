package main

import (
	"errors"
	"fmt"
)

func funcA() error {
	err := funcB()

	if err != nil {
		return fmt.Errorf("funcB failed! error:%w", err)
	}

	return nil
}

func funcB() error {
	err := funcC()

	if err != nil {
		return fmt.Errorf("funcC failed! error:%w", err)
	}

	return nil
}

func funcC() error {
	return errors.New("funcC exec error!")
}

func testError() {
	err := funcA()

	if err != nil {
		fmt.Println(err.Error())
	}

	tempErr := err

	for {
		tempErr = errors.Unwrap(tempErr)

		if tempErr != nil {
			fmt.Println(tempErr.Error())
		} else {
			break
		}
	}
}

func main() {
	testError()
}
