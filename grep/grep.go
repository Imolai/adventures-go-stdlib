// grep.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// Pattern to search for and the text to search in
	text := "Hello, world of worlds! Welcome to the World of Go."
	pattern := `\b[A-Z][a-z]+\b`

	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalln("Error compiling pattern:", err)
	}

	// Search for all matches
	matches := re.FindAllString(text, -1)

	// Print matches to STDOUT through a buffered writer
	w := bufio.NewWriter(os.Stdout)
	for _, match := range matches {
		fmt.Fprint(w, match+"\n")
	}

	// Flush the buffer to write all data to STDOUT
	w.Flush()
}

/*
OUTPUT:
Hello
Welcome
World
Go
*/
