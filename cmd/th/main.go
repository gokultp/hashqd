package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()
	fmt.Println(a.Nanosecond(), a.Nanosecond()/1000)

}
