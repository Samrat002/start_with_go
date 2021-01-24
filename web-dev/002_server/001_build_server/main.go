package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err!=nil{
		log.Panic(err)
	}
	defer listener.Close()
	for{
		conn, err:=listener.Accept()
		if err!=nil{
			log.Println(err)
		}
		io.WriteString(conn, "\nWelcome to tcp connection!!!!!\n")
		io.WriteString(conn, "Closing the requester connection")
		conn.Close()
	}
}
