package codetool

import (
	"time"
)

// IsNumeric - Finds whether a variable is a number or a numeric string
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

// IsStringNumeric - return if the string var is numeric format
func IsStringNumeric(x string) bool {
	hasPeriod := false
	for i, c := range x {
		switch c {
		case '-':
			if i != 0 {
				return false
			}
		case '.':
			if hasPeriod {
				return false
			}
			hasPeriod = true
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			//Nothing here.
		default:
			return false
		}
	}
	return true
}

// Strtotime - use to Conversion string format time to unix timestamp
func Strtotime(timeStr string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	var sr int64 = 0

	if loc, err := time.LoadLocation("Local"); err == nil {
		if theTime, err := time.ParseInLocation(timeLayout, timeStr, loc); err == nil {
			sr := theTime.Unix()
			return sr
		}
	}

	return sr
}
