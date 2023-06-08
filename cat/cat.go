// cat.go
package main

import (
        "bufio" // buffered I/O https://pkg.go.dev/bufio
        "fmt"   // formatting and printing https://pkg.go.dev/fmt
        "os"    // operating system functionality https://pkg.go.dev/os
)

func main() {
        scanner := bufio.NewScanner(os.Stdin) // new scanner to read input from standard input
        for scanner.Scan() {                  // read input line by line
                fmt.Println(scanner.Text()) // current input line to standard output
        }

        if err := scanner.Err(); err != nil { // if any error occurred during scanning
                fmt.Fprintln(os.Stderr, "reading standard input:", err) // error message to standard error
        }
}

/*
$ go run cat.go
Hello, World!
Hello, World!
This is a simple implementation of Unix "cat".
This is a simple implementation of Unix "cat".
*/