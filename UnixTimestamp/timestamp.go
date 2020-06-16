package main

import (
	"fmt"
	"time"
)

func main() {
	unixTimeUTC := time.Unix(1405544146, 0) //gives unix time stamp in utc

	unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339) // converts utc time to RFC3339 format

	fmt.Println("unix time stamp in UTC :--->", unixTimeUTC)
	fmt.Println("unix time stamp in unitTimeInRFC3339 format :->", unitTimeInRFC3339)

	fmt.Println(unixTimeUTC.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(unixTimeUTC.Format("2006-01-02 15:04:05"))

	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("current date is :%s \n", date) // run on local env

}
