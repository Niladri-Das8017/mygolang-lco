package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Time Package ")
	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	creatdDate := time.Date(2020, time.October, 12, 23, 23, 0, 0, time.UTC )

	fmt.Println(creatdDate.Format("01-02-2006 Monday"))

}
