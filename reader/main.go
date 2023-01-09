package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fn := ""
	for _, arg := range os.Args[1:] {
		fn = arg
	}
	fmt.Println(fn)
	f, e := os.Open(fn)
	if e != nil {
		fmt.Println("ERROR: File could not be opened!")
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
