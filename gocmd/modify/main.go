package main

import (
	"encoding/json"
	"log"
	"os"
)

type Car struct {
	Brand  string
	Engine string
}

func main() {
	err := json.NewEncoder(os.Stdout).Encode(Car{"Honda", "diesel"})
	if err != nil {
		log.Fatal(err)
	}
}
