package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var s, sep string
	l := len(os.Args)
	fmt.Printf("args len=%d\n\r", l)
	t := time.Now()
	for i := 0; i < l; i++ {
		s += sep + os.Args[i]
		sep = " "
		time.Sleep(time.Nanosecond)
	}
	fmt.Println(s, " ", time.Since(t).Microseconds())
}
