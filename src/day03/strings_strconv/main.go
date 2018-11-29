package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now)

    fmt.Println(time.Time{})
    fmt.Println(time.Now().Day())
    fmt.Println(time.Now().Minute())
    fmt.Println(time.Now().Month())
    fmt.Println(time.Now().Year())

    fmt.Printf("%02d/%02d/%02d %02d:%02d:%02d",
        time.Now().Year(), time.Now().Month(), time.Now().Day(),
        time.Now().Hour(), time.Now().Minute(), time.Now().Second())

    fmt.Println(time.Now().Format("02/1/2006 15:04"))
    fmt.Println(time.Now().Format("2006/1/02 15:04"))
    fmt.Println(time.Now().Format("2006/1/02"))


}


