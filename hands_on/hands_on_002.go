package main

import "fmt"

// HANDS ON 2
// create a struct that holds person fields
// create a struct that holds secret agent fields and embeds person type
// attach a method to person: pSpeak
// attach a method to secret agent: saSpeak
// create a variable of type person
// create a variable of type secret agent
// print a field from person
// run pSpeak attached to the variable of type person
// print a field from secret agent
// run saSpeak attached to the variable of type secret agent
// run pSpeak attached to the variable of type secret agent
type person struct {
	name string
	age int

}
func (p person) speak() string{
	return "i am citizen"
}

type secreatAgent struct {
	person
	licenceToKill bool
}
func (sa secreatAgent) speak() string{
	return "Iam secret agent"
}

func main() {
	per:= person{
		name: "Alex",
		age:  28,
	}
	sa := secreatAgent{
		person:        person{
			name: "James Bond",
			age:  34,
		},
		licenceToKill: true,
	}
	fmt.Println("Person speaks: - "+per.speak())
	fmt.Println("Secret Agent speaks: - "+sa.speak())
	fmt.Println("Secret Agent also speaks: - "+sa.person.speak())
}
