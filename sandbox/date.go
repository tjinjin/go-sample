package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Date(2016, 10, 01, 0, 0, 0, 0, time.Local)
	currentTime := time.Date(2016, 11, 10, 0, 0, 0, 0, time.Local)
	duration := currentTime.Sub(startTime)
	fmt.Println(startTime)
	fmt.Println(currentTime)
	fmt.Println(duration)
	// 日付感のdurationを取得して何日差か日で示
	fmt.Println(duration.Hours() / 24)
	fmt.Println(startTime.Add(duration))
	t := time.Now()
	fmt.Println(t.Month())
}
