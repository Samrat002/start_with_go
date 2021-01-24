package main

import (
	"fmt"
	"net/http"
)

type MyFirstType int
func (MyFirstType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is code")

}
func main() {
	var d MyFirstType
	http.ListenAndServe(":8080", d)
}
