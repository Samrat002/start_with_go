package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

)

func handle(conn net.Conn){
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln:=scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err!=nil{
		log.Panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err!=nil{
			log.Print(err)
		}
		go handle(conn)
	}
}
