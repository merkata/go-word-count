package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

// ErrInvalidFlags raises an error if multiple flags are set
// where you'd want only one, example - print line and word count
var ErrInvalidFlags = errors.New("invalid flags or combination of supplied")

func main() {
	lines := flag.Bool("l", false, "count lines")
	bytes := flag.Bool("b", false, "count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(l io.Reader, countLines bool, countBytes bool) (int, error) {
	b := bufio.NewScanner(l)

	if countLines && countBytes {
		return 0, ErrInvalidFlags
	}
	if !countLines && !countBytes {
		b.Split(bufio.ScanWords)
	} else if !countLines {
		b.Split(bufio.ScanBytes)
	}

	wc := 0
	for b.Scan() {
		wc++
	}
	return wc, nil
}
