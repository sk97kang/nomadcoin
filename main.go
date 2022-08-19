package main

import "fmt"

func plus(a ...int) int {
	return a[0] + a[1]
}

func main() {
	result := plus(1, 2)
	fmt.Println(result)
}