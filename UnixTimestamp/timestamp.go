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
}
