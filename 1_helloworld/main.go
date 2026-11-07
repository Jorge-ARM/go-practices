package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello world, I'm Jorge")
	fmt.Println("Version:" + runtime.Version())
	/*
		Can be changed to: fmt.Printf("Version: %s\n", runtime.Version())

		%s: represents a text string.
		%d: represents an integer.
		%f: represents a decimal number.
		%v: represents a value in its default format.
		\n: line break.
	*/
}
