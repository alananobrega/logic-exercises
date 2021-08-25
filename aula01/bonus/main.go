// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.

//receber um num
//varrer num by num do num que chega
//eh multiplo
//a = a + num
//retorna a

package main

import "fmt"

func main() { //O(2n)
	a := sum_mutiple(1000)           //n
	b := sum_mutiple(2000)           //n
	fmt.Printf("a soma é: %d \n", a) //1
	fmt.Printf("a soma é: %d \n", b) //1
}

func sum_mutiple(num int) int { //O(n)
	soma := 0

	for i := 1; i < num; i++ { //n
		if i%3 == 0 || i%5 == 0 {
			soma += i //1
		}
	}
	return soma
}
