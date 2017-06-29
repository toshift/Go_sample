package main

import (
	"fmt"
	"os"
)

const BUFSIZE = 1024

func main() {
	file, err := os.Open(`C:\\work\\Gowork\\hello.go`)
	if err != nil {
		//		fmt.Printf(err)

	}
	defer file.Close()
	defer fmt.Printf("end!")

	buf := make([]byte, BUFSIZE)
	for {

		readbuff, err := file.Read(buf)
		if (readbuff == 0) || (err != nil) {
			// Readエラー
			fmt.Printf("Read Error \n")
		}
		fmt.Printf(string(buf[:readbuff]))
	}
}
