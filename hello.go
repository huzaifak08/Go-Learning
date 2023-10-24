package main

import "fmt"

func simpleAdd(x int, y int) int {
	return x + y
}

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello World We are learning Golang Allhamdullilah")

	fmt.Println(add(1, 2))       // 3
	fmt.Println(simpleAdd(3, 4)) // 7

}
