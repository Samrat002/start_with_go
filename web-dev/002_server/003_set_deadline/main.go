package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	listener, err:= net.Listen("tcp", ":8080")
	if err!=nil{
		log.Panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err!=nil{
			log.Print(err)
			continue
		}
		go handle(conn)

	}
}

func handle(conn net.Conn){
	err:= conn.SetDeadline(time.Now().Add(10*time.Second))
	if err!=nil{
		log.Fatal("Connection TimeOut")
	}
	scanner:= bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	defer conn.Close()
	fmt.Println("Code goes here !!!!!!")


}
