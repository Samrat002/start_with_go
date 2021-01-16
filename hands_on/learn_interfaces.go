package main

import "fmt"

type bot interface {
	speaks(string) string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) speaks(name string) string{
	return "Hello "+name
}


func (spanishBot) speaks(name string) string{
	return "Holla "+name
}

func main(){
	eb := englishBot{}
	sb := spanishBot{}
	prints(eb, "Sammy")
	prints(sb, "Susu")
}

func prints(b bot, name string){
	fmt.Println(b.speaks(name))
}
