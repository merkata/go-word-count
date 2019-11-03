package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "count lines")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}

func count(l io.Reader, countLines bool) int {
	b := bufio.NewScanner(l)
	if !countLines {
		b.Split(bufio.ScanWords)
	}
	wc := 0
	for b.Scan() {
		wc++
	}
	return wc
}
