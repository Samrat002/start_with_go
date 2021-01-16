package main

import (
	"fmt"
	"math"
)

// HANDS ON 1
// create a type square
// create a type circle
// attach a method to each that calculates area and returns it
// create a type shape which defines an interface as anything which has the area method
// create a func info which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

type square struct{
	side float64
}

func (s square) area() float64 {
	return s.side*s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi*math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	my_square := square{side: 15.5}
	my_circle := circle{radius: 15.5}
	fmt.Println("Printing area of a circle")
	info(my_circle)
	fmt.Println("Printing area of a square")
	info(my_square)

}
