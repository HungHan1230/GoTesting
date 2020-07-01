package main

import (
	"fmt"
	"strings"
	"time"
)

func Run_time_example() {
	p := fmt.Println
	now := time.Now()
	p(now)
	then := time.Date(2017, 06, 21, 20, 34, 58, 0, time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	p(then.Add(diff))
	p(then.Add(-diff))

	manipulation()
}
func printSplit(count int) {
	fmt.Println(strings.Repeat("#", count))
}
func manipulation() {
	fmt.Println("manipulation")
	// Add 時間相加
	now := time.Now()
	// ParseDuration parses a duration string.
	// A duration string is a possibly signed sequence of decimal numbers,
	// each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// 10分鍾前
	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(m)
	fmt.Println(m1)

	// 8個小時前
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(8 * h)
	fmt.Println(h1)

	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1)

	printSplit(50)

	// 10分鍾后
	mm, _ := time.ParseDuration("1m")
	mm1 := now.Add(mm)
	fmt.Println(mm1)

	// 8小時后
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(hh)
	fmt.Println(hh1)

	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println(dd1)

	printSplit(50)

	// Sub 計算兩個時間差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分鍾")
	var t float64 = subM.Minutes()
	fmt.Println(t)

	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小時")

	sumD := now.Sub(d1)
	fmt.Printf("%v 天\n", sumD.Hours()/24)
}
