package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(count(os.Stdin))
}

func count(l io.Reader) int {
	b := bufio.NewScanner(l)
	b.Split(bufio.ScanWords)
	wc := 0
	for b.Scan() {
		wc++
	}
	return wc
}
