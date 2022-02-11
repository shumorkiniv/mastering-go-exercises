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

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
