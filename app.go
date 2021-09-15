package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
	c := Plus(1, 2)
	fmt.Println(c)
}

func Plus(a int, b int) int {
	return a + b
}