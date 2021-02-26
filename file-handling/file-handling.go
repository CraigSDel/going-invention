package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Look at us go :P")
	file.Close()

	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	line := string(stream)
	fmt.Println(line)
}
