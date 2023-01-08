package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// r := make([]byte, 9999)
	// resp.Body.Read(r)
	// fmt.Println(string(r))
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
