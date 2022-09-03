package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	tempFile := os.Args[1]
	f, err := os.Open(tempFile)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	c, err := io.Copy(os.Stdout, f)

	fmt.Println(c)
}
