package main

import "fmt"

//create an interface type that both person and secretAgent implement
//declare a func with a parameter of the interface’s type
//call that func in main and pass in a value of type person
//call that func in main and pass in a value of type secretAgent
type humans interface {
	speak() string
}

func printWhatTheySpeak(h humans) {
	fmt.Println(h.speak())
}
func main() {
	david:= person{
		name: "David Aron",
		age:  23,
	}
	james:= secreatAgent{
		person:        person{
			name: "James Bond",
			age:  32,
		},
		licenceToKill: true,
	}
	printWhatTheySpeak(david)
	printWhatTheySpeak(james)

}
