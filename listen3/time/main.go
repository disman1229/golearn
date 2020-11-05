package main

import (
	"fmt"
	"time"
)

func testTime() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	send := now.Second()
	fmt.Printf("当前时间:%d年%d月%d日，%d时%d分%d秒\n", year, month, day, hour, minute, send)
	fmt.Printf("当前时间:%02d年%02d月%02d日，%02d时%02d分%02d秒\n", year, month, day, hour, minute, send)

	// 当前时间戳
	timestamp := now.Unix()
	fmt.Printf("timestamp is:%d\n", timestamp)
}

func testTimeStamp(timestamp int64) {

	timeObj := time.Unix(timestamp, 0)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	send := timeObj.Second()
	fmt.Printf("current timestamp:%d\n", timestamp)
	fmt.Printf("当前时间:%d年%d月%d日，%d时%d分%d秒\n", year, month, day, hour, minute, send)
}

// Second-秒，Millisecond-毫秒，Microsecond-微秒，Nanosecond-纳秒
func testTicker() {
	ticker := time.Tick(1 * time.Second)
	for i := range ticker {
		fmt.Printf("%v\n", i)
		processTask()
	}
}

func processTask() {
	fmt.Println("do tack")
}

//
func testFormat() {
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05") //日期必须写2006
	fmt.Println(timeStr)
}

func main() {
	// testTime()
	// timestamp := time.Now().Unix()
	// testTimeStamp(timestamp)
	// testTicker()
	testFormat()
}
