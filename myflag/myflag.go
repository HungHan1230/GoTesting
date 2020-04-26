package main

import (
	"fmt"
	"flag"
)

//define flags
var inputName = flag.String("name", "XuChao", "Input your name")
var inputAge = flag.Int("age", 25, "Input your age")
var inputGender = flag.String("gender", "boy", "Input your gender")

func main(){
	flag.Parse()     //parse flag
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	
	fmt.Println("name=", *inputName)
	fmt.Println("age=", *inputAge)
	fmt.Println("gender=", *inputGender)
}