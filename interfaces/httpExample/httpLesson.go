package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://www.google.fi")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//#1 Trivial Reader implementation
	byteStream := make([]byte, 1000)
	resp.Body.Read(byteStream)
	resp.Body.Close()
	fmt.Println(byteStream)

	//#2 io.Copy(Writer ,  Reader) this method expects two objects, one that implements
	//	 a Writer interface and the other that implements the Reader
	io.Copy(os.Stdout, resp.Body)
}
