// wc.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // new scanner to read input from standard input
	var label string
	switch os.Args[1] { // determine scanner split function based on CLI flag
	case "-l":
		scanner.Split(bufio.ScanLines)
		label = "Lines"
	case "-w":
		scanner.Split(bufio.ScanWords)
		label = "Words"
	case "-c":
		scanner.Split(bufio.ScanRunes)
		label = "Runes"
	default:
		scanner.Split(bufio.ScanBytes)
		label = "Bytes"
	}
	count := 0
	for scanner.Scan() { // read input splitted by the separator
		count++
	}
	if err := scanner.Err(); err != nil { // if any error occurred during scanning
		fmt.Fprintln(os.Stderr, "reading input:", err) // error message to standard error
	}
	fmt.Printf("%s #: %d\n", label, count)
}

/*
$ go run wc.go -l

Go is a powerful, efficient,
and expressive
programming 世界.

Lines #: 5
$ go run wc.go -w

Go is a powerful, efficient,
and expressive
programming 世界.

Words #: 9
$ go run wc.go -c

Go is a powerful, efficient,
and expressive
programming 世界.

Runes #: 62
$
*/
