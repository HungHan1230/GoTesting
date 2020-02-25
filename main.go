package main

import (
	"fmt"

	TransposeMatrix "github.com/HungHan1230/GoTesting/TransposeMatrix"

	MyTestingReadFile "github.com/HungHan1230/GoTesting/MyTestingReadFile"
)

func main() {
	mainTransposeMatrix()
	mainMyTestingReadFile()
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
