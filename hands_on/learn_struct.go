package main

import "fmt"

type Person struct {
	name string
	age int
	address Address

}

type Address struct {
	street1 string
	street2 string
	pin string
	state string
	country string
}
func (p *Person) UpdateName(n string){
	(*p).name = n

}
func main() {

	alex := Person{
		name:    "Alex Xander",
		age:     10,
		address: Address{
			street1: "Streat 1 ",
			street2: "taste the food",
			pin:     "786145",
			state:   "Calafornia",
			country: "USA",
		},
	}
	fmt.Println(alex)
	var david Person
	david.name = "David Halson"
	david.age = 20
	david.address = Address{
		street1: "taste the first",
		street2: "Taste the second",
		pin:     "23456",
		state:   "Perkins",
		country: "Canada",
	}
	fmt.Println(david)
	fmt.Printf("%+v\n", david)

	///////////////////////////////// LEARNING POINTER /////////////////////////////////////////
	// using & before variable makes it a pointer
	// using * before an address makes it to a actual value
	alexPointer := &alex
	alexPointer.UpdateName("Alexender")
	fmt.Println(alex)

	/////// exploring go shortcut
	david.UpdateName("Davion")
	fmt.Println(david)


}

