package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
)
func main() {
	dates := []string{
		"20201202",
		"20201203",
		"20201204",
		"20201205",
		"20201206",
		"20201207",
		"20201208",
	}

	for _, date := range dates {
		hideStart := fmt.Sprintf("%v 00:00:00", date)
		hideEnd := fmt.Sprintf("%v 23:59:59", date)
		printTime("Hide", hideStart, hideEnd)

		activity1Start := fmt.Sprintf("%v 09:00:00", date)
		activity1End := fmt.Sprintf("%v 09:59:59", date)
		printTime("Activity-1", activity1Start, activity1End)

		activity2Start := fmt.Sprintf("%v 12:00:00", date)
		activity2End := fmt.Sprintf("%v 12:59:59", date)
		printTime("Activity-2", activity2Start, activity2End)

		activity3Start := fmt.Sprintf("%v 15:00:00", date)
		activity3End := fmt.Sprintf("%v 15:59:59", date)
		printTime("Activity-3", activity3Start, activity3End)

		activity4Start := fmt.Sprintf("%v 19:00:00", date)
		activity4End := fmt.Sprintf("%v 19:59:59", date)
		printTime("Activity-4", activity4Start, activity4End)

		activity5Start := fmt.Sprintf("%v 21:00:00", date)
		activity5End := fmt.Sprintf("%v 21:59:59", date)
		printTime("Activity-5", activity5Start, activity5End)

	}

	return
	fmt.Println(now.BeginningOfDay())
	fmt.Println(now.EndOfDay())
	fmt.Println(now.BeginningOfMonth())
	fmt.Println(now.EndOfMonth())
}

func printTime(name, start, end string) {
	loc, _ := time.LoadLocation("PRC")

	timeStart, _ := time.ParseInLocation("20060102 15:04:05", start, loc)
	timeEnd, _ := time.ParseInLocation("20060102 15:04:05", end, loc)

	fmt.Printf("Activity:%v,\tStartTime:%v(unix:%v),\tEndTime:%v(unix:%v)\r\n", name, timeStart, timeStart.Unix(), timeEnd, timeEnd.Unix())
}
