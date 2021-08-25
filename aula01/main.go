package main

import "fmt"

func main() {

	fmt.Printf("o fatorial eh: %d \n", factorial(21))
	fmt.Println(factorial2(3))
}

func factorial(num int) int { //O(n)
	factorial := 1

	for i := num; i > 1; i-- { //n
		factorial = factorial * i //1
	}

	return factorial
}

// 3! = 3*2*1 = 6
// 4! = 4*3*2*1 = 24
