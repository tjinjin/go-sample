package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, _ := os.Open("./file.txt")
	message, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(message))
}
