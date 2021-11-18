package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	defer filenotFoundCatch()
	readFile()
}

func readFile() {
	if _, err := ioutil.ReadFile("./not-exist.txt"); err != nil {
		panic("File not found")
	}
}

func filenotFoundCatch() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
