package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Add the filename as an argument and retry")
		os.Exit(1)
	}

	readFileUsingIoutils(args)

	readFileUsingReadInterface(args[1])
}

func readFileUsingReadInterface(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("---Copying the file contents to stdout---")
	io.Copy(os.Stdout, file)
}

func readFileUsingIoutils(args []string) {
	if len(args) > 1 {
		filename := args[1]
		fmt.Println(filename)
		bs, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(string(bs))
	}
}
