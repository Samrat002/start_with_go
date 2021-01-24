package main

import (
	"html/template"
	"log"
	"os"
)

/*
Create a data structure to pass to a template which
contains information about restaurant's menu including Breakfast, Lunch, and Dinner items
 */
type SingleItemDetails struct {
	Name string
	FoodType string
	Price float32
}
type Menu []SingleItemDetails
type restaurant struct {
	Breakfast, Lunch, Dinner Menu
}
var abspath = `/Users/debs/Personal/start_with_go/web-dev/hands_on_template/001_practice`
var tpl *template.Template

func init(){
	tpl = template.Must(template.New("xtemplate.gohtml").ParseFiles(abspath+"/xtemplate.gohtml"))
}

func main(){
	newFile, err := os.Create(abspath+"/index.html")
	if err!=nil{
		log.Fatal(err)
	}
	xrestaurantMenu:= restaurant{
		 Menu{
			{
				Name:     "Tea",
				FoodType: "Breverages",
				Price:    100,
			},
			{
				Name:     "Noodles",
				FoodType: "Continental",
				Price:    150.67,
			},
			{
				Name:     "Paratha",
				FoodType: "Indian",
				Price:    170.56,
			},
		},
		    Menu{
			{
				Name:     "Small Lunch Thali",
				FoodType: "Indian",
				Price:    500,
			},
			{
				Name:     "Large Lunch Variety Thali",
				FoodType: "Indian",
				Price:    700,
			},
		},
		   Menu{
			{
				Name:     "BigThali",
				FoodType: "Indian",
				Price:    700,
			},
		},
	}
	err = tpl.Execute(newFile, xrestaurantMenu)
	if err!=nil{
		log.Fatal(err)
	}
}