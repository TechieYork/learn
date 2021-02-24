package codetool

import (
	"fmt"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
)

// go test -v -gcflags=all=-l -coverpkg=./... -coverprofile=test.out ./...
// go tool cover -html=test.out -o test.html
func TestConveyIsNumeric(t *testing.T) {
	/*
	   	func IsNumeric(x interface{}) (result bool) {
	      	//Figure out result
	      	switch x.(type) {
	      	case int, uint:
	      		result = true
	      	case int8, uint8:
	      		result = true
	      	case int16, uint16:
	      		result = true
	      	case int32, uint32:
	      		result = true
	      	case int64, uint64:
	      		result = true
	      	case float32, float64:
	      		result = true
	      	case complex64, complex128:
	      		result = true
	      	case string:
	      		if xAsString, ok := x.(string); ok {
	      			result = IsStringNumeric(xAsString)
	      		} else {
	      			result = false
	      		}
	      	default:
	      		result = false
	      	}
	      	return result
	      }
	*/
	Convey("Test IsNumeric", t, func() {
		Convey("Test IsNumeric with different type", func() {
			type TestData struct {
				Type   string
				X      interface{}
				Expect bool
			}

			testCases := map[string]TestData{
				"int":            {Type: "int", X: int(0), Expect: true},
				"uint":           {Type: "uint", X: uint(0), Expect: true},
				"int8":           {Type: "int8", X: int8(0), Expect: true},
				"uint8":          {Type: "uint8", X: uint8(0), Expect: true},
				"int16":          {Type: "int16", X: uint16(0), Expect: true},
				"uint16":         {Type: "uint16", X: uint16(0), Expect: true},
				"int32":          {Type: "int32", X: int32(0), Expect: true},
				"uint32":         {Type: "uint32", X: uint32(0), Expect: true},
				"int64":          {Type: "uint64", X: uint64(0), Expect: true},
				"float32":        {Type: "float32", X: float32(0), Expect: true},
				"float64":        {Type: "float64", X: float64(0), Expect: true},
				"complex64":      {Type: "complex64", X: complex64(0), Expect: true},
				"complex128":     {Type: "complex128", X: complex128(0), Expect: true},
				"string":         {Type: "string", X: "string", Expect: false},
				"numeric string": {Type: "numeric string", X: "111", Expect: true},
				"struct":         {Type: "struct", X: struct{}{}, Expect: false},
			}

			for testCaseName, testCase := range testCases {
				Convey(testCaseName, func() {
					result := IsNumeric(testCase.X)
					So(result, ShouldEqual, testCase.Expect)
				})
			}
		})
	})
}

func TestConveyIsStringNumeric(t *testing.T) {
	Convey("Test IsStringNumeric", t, func() {
		Convey("Test IsStringNumeric normal", func() {
			result := IsStringNumeric("123.456")
			So(result, ShouldEqual, true)
		})

		Convey("Test IsStringNumeric with '.'", func() {
			var result bool

			result = IsStringNumeric("123.456")
			So(result, ShouldEqual, true)

			result = IsStringNumeric(".123456")
			So(result, ShouldEqual, true)

			result = IsStringNumeric("123456.")
			So(result, ShouldEqual, true)

			result = IsStringNumeric("123.456.")
			So(result, ShouldEqual, false)
		})

		Convey("Test IsStringNumeric with '-'", func() {
			var result bool

			result = IsStringNumeric("-123456")
			So(result, ShouldEqual, true)

			result = IsStringNumeric("123-456")
			So(result, ShouldEqual, false)

			result = IsStringNumeric("123456-")
			So(result, ShouldEqual, false)

			result = IsStringNumeric("-123-456")
			So(result, ShouldEqual, false)
		})

		Convey("Test IsStringNumeric with illegal chars", func() {
			result := IsStringNumeric("123.456abc")
			So(result, ShouldEqual, false)
		})
	})
}

func TestConveyStrtotime(t *testing.T) {
	Convey("Test Strtotime", t, func() {
		Convey("Test Strtotime normal", func() {
			var result int64

			result = Strtotime("2020-06-30 11:43:25")
			So(result, ShouldEqual, 1593488605)
		})

		Convey("Test Strtotime with illegal format", func() {
			var result int64

			result = Strtotime("2020-0630 11:43:25")
			So(result, ShouldEqual, 0)
		})

		Convey("Test Strtotime with LoadLocation failed", func() {
			p := gomonkey.ApplyFunc(time.LoadLocation, func(string) (*time.Location, error) {
				return nil, fmt.Errorf("time.LoadLocation failed")
			})

			defer p.Reset()

			var result int64

			result = Strtotime("2020-06-30 11:43:25")
			So(result, ShouldEqual, 0)
		})

		Convey("Test Strtotime with ParseInLocation failed", func() {
			p := gomonkey.ApplyFunc(time.ParseInLocation, func(string, string, *time.Location) (time.Time, error) {
				return time.Unix(0, 0), fmt.Errorf("time.ParseInLocation failed")
			})

			defer p.Reset()

			var result int64

			result = Strtotime("2020-06-30 11:43:25")
			t.Logf("result:%v", result)
			So(result, ShouldEqual, 0)
		})
	})
}
