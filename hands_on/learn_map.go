package main

import "fmt"

func main(){
	colors := map[string]string{
		"red": "#fffff",
		"yellow":"#de43534",
		"green": "#ffde34",

	}
	fmt.Println(colors)
	// adding more keys
	colors["brown"] = "#23edf"
	fmt.Println(colors)

	colours:=make( map[string]string)
	colours["gold"] = "gg"
	colours["silver"] = "s"
	fmt.Println(colours)

	// delete a key
	delete(colors, "red")
	fmt.Println(colors)

	// iterate over the map
	for key, value:= range colors {
		fmt.Println(key, value)
	}
}
