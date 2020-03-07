package main

import (
	"fmt"
	"time"
)

func main() {
	const t1 = "2006-01-02 15：04：05 -0700"
	const t2 = "06-Jan-02 3:4:5.000 PM MST -0700"

	st1 := time.Now().Format(t1)
	st2 := time.Now().Format(t2)

	dt1, _ := time.Parse("Jan 2 15:04:05 -0700 2006", "Mar 7 18:01:30 +0800 2020")
	fmt.Printf("\rt1=%s\nt2=%s\n", st1, st2)
	fmt.Printf("\rdt1=%v\n", dt1)
}
