package main

import (
	"io"
	"os"
)

func main() {
	myString := ""

	args := os.Args

	if len(args) == 1 {
		myString = "Please give me one more argument!"
	} else {
		myString = args[1]
	}

	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
