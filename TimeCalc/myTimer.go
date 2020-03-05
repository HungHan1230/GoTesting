package main

import (
    "fmt"
    "time"
)

func StartCac() {
    t1 := time.Now() // get current time
    //logic handlers
    for i := 0; i < 1000; i++ {
        fmt.Print("*")
    }
    elapsed := time.Since(t1)
    fmt.Println("App elapsed: ", elapsed)
}

func main(){
    StartCac()
}
