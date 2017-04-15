package main

import "os"

func main() {
	src, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()
}
