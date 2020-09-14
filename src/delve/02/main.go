package main

import "fmt"

func add(left, right int) int {
	return left + right
}

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum = add(sum, i)
	}
	fmt.Println(sum)
}
