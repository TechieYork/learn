package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/xerrors"
	pkgErrors "github.com/pkg/errors"
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

func funcX() error {
	err := funcY()

	if err != nil {
		return xerrors.Errorf("funcY failed! error:%w", err)
	}

	return nil
}

func funcY() error {
	err := funcZ()

	if err != nil {
		return xerrors.Errorf("funcZ failed! error:%w", err)
	}

	return nil
}

func funcZ() error {
	return xerrors.Errorf("funcZ exec error!")
}

func testXerror() {
	err := funcX()

	if err != nil {
		fmt.Printf("%+v", err)
	}
}

func func1() error {
	err := func2()

	if err != nil {
		return pkgErrors.Wrapf(err, "func2 failed! error:%v", err)
	}

	return nil
}

func func2() error {
	err := func3()

	if err != nil {
		return pkgErrors.Wrapf(err, "func3 failed! error:%v", err)
	}

	return nil
}

func func3() error {
	return pkgErrors.Errorf("func3 exec error!")
}

func testPkgError() {
	err := func1()

	if err != nil {
		fmt.Printf("%+v", err)
	}
}

func main() {
	awardIDs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}


	// split awardIDs to different group
	awardIDGroup := make([][]string, 0)
	tempGroup := make([]string, 0)

	for _, awardID := range awardIDs {
		tempGroup = append(tempGroup, awardID)

		if len(tempGroup) >= 10 {
			awardIDGroup = append(awardIDGroup, tempGroup)
			tempGroup = make([]string, 0)
		}
	}

	if len(tempGroup) != 0 {
		awardIDGroup = append(awardIDGroup, tempGroup)
	}
	fmt.Println(awardIDGroup)
	return

	curDate, err := time.Parse("20060102", "19700101")

	if err != nil {
		fmt.Printf("time.Parse failed! error:%v", err)
		return
	}

	fmt.Println(curDate.String())

	tommorow := curDate.Add(time.Hour * 24)
	fmt.Println(tommorow.String())

	return

	fmt.Println("====== test xerrors ======")
	testXerror()
	fmt.Println("")

	fmt.Println("====== test errors ======")
	testError()
	fmt.Println("")

	fmt.Println("====== test pkg errors ======")
	testPkgError()
	fmt.Println("")
}
