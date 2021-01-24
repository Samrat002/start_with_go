package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err!=nil{
		log.Panic(err)
	}
	defer conn.Close()
	bs, err := ioutil.ReadAll(conn)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}
