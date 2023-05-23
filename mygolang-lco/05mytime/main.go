package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format(time.Kitchen))
	fmt.Println(presentTime.Format("Jan 02 2006 15:04:05 PM"))

	createdDate := time.Date(1996, time.September, 19, 5, 45, 0, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("Jan 02, 2006"))
}
