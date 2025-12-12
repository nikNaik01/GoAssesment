package main

import "fmt"

func swap(x, y *int) {
	//Your code goes here
	temp := *x
	*x = *y
	*y = temp
}

// Do not change the code in the main function
func main() {
	x := 10
	y := 20

	swap(&x, &y)
	fmt.Println(x, y)
}
