package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("test.md")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))

}
