package main

import (
	"errors"
	"testing"

	"github.com/bouk/monkey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckParamPrefix(t *testing.T) {
	//Test convey
	Convey("Test CheckParamPrefix function", t, func() {
		Convey("Test empty string", func() {
			So(CheckParamPrefix(""), ShouldBeError)
		})

		Convey("Test length invalid", func() {
			So(CheckParamPrefix("111222333444"), ShouldBeError)
		})

		Convey("Test prefix regexp", func() {
			So(CheckParamPrefix("test_ok"), ShouldBeNil)
			So(CheckParamPrefix("test-not"), ShouldBeError)
		})
	})

	//Test monkey
	Convey("Test CheckParamPrefix with monkey", t, func() {
		defer monkey.UnpatchAll()

		Convey("Test empty string", func() {
			monkey.Patch(CheckParamPrefix, func(content string) error {
				return errors.New("monkey patch error")
			})

			So(CheckParamPrefix(""), ShouldBeError)

			monkey.Unpatch(CheckParamPrefix)
		})

		Convey("Test length invalid", func() {
			monkey.Patch(CheckParamPrefix, func(content string) error {
				return errors.New("monkey patch error")
			})

			So(CheckParamPrefix("111222333444"), ShouldBeError)

			monkey.Unpatch(CheckParamPrefix)
		})

		Convey("Test prefix regexp", func() {
			So(CheckParamPrefix("test_ok"), ShouldBeNil)
			So(CheckParamPrefix("test-not"), ShouldBeError)
		})
	})
}
