package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int, 0)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() { //ctrl+D中断
		counts[input.Text()]++
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
