package main

import (
	"fmt"
	"math"
	"os"
	"time"

	TransposeMatrix "github.com/HungHan1230/GoTesting/TransposeMatrix"

	MyTestingReadFile "github.com/HungHan1230/GoTesting/MyTestingReadFile"
)

func main() {
	// mainTransposeMatrix()
	// mainMyTestingReadFile()
	testDate()
}

func mainTransposeMatrix() {
	sample := [][]string{
		[]string{"a1", "a2", "a3", "a4", "a5"},
		[]string{"b1", "b2", "b3", "b4", "b5"},
		[]string{"c1", "c2", "c3", "c4", "c5"},
	}
	ar := TransposeMatrix.Transpose(sample)
	fmt.Println(ar)
}
func mainMyTestingReadFile() {
	MyTestingReadFile.Run()
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func testDate() {
	var t1 int64 = 1588181730
	// var t2 int64 = 1589442289
	var t2 int64 = 1588250008
	unixTimeUTC1 := time.Unix(t1, 0) //gives unix time stamp in utc
	unixTimeUTC2 := time.Unix(t2, 0) //gives unix time stamp in utc

	fmt.Println("unix time stamp in UTC :--->", unixTimeUTC1)
	fmt.Println("unix time stamp in UTC :--->", unixTimeUTC2)

	diff := unixTimeUTC2.Sub(unixTimeUTC1)
	fmt.Println("days: ", diff.Hours()/24)
	fmt.Println("days: ", int(diff.Hours()/24))
	fmt.Println("days: ", math.Ceil(diff.Hours()/24))

	// dict := make(map[string]string)
	// dict["1588181730"] = "1"
	// dict["1588250008"] = "2"

	// if val, ok := dict["1588181730_"]; ok {
	// 	//do something here
	// 	fmt.Println(val)
	// }else{
	// 	fmt.Println("found nothing")
	// }


	// layout := "2006-01-02 15:04:05"
	// current_timestamp, err := time.Parse(layout, t1)

}
