package main

import (
	"errors"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":3003")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go ClientHandler(conn)
	}
}

func ClientHandler(c net.Conn) {
	defer func() {
		if v := recover(); v != nil {
			log.Println("capture a panic:", v)
			log.Println("avoid crashing the program")
		}
		c.Close()
	}()
	panic(errors.New("Panic in every client"))
}
