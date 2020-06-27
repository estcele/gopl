package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	counts := make(map[string]int, 0)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}

}

func countLines(f *os.File, counts map[string]int) {
	fmt.Println(filepath.Abs(f.Name()))
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
