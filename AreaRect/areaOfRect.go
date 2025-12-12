package main

import "fmt"

//Your code goes here

type Rectangle struct {
	Width  float32
	Height float32
}

func (rect Rectangle) Area() float32 {
	area := rect.Width * rect.Height
	return area
}

// Do not change the code in the main function
func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area:", rect.Area())

}
