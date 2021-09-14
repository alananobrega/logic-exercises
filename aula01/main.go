package main

import "fmt"

func main() {

	fmt.Printf("o fatorial eh: %d \n", factorial(6))
}

func factorial(num int) int { //O(n)
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

// 2! = 2*1! = 2
// 3! = 3*2! = 6
// 4! = 4*3! = 24
// 4! = 4*3! = 24
// 4! = 4*3! = 24
// n! = n*(n-1)!

//quantas vezes que ocorre uma chamada = n

//CALCULO COMPLEXIDADE ALG. RECURSIVO
//Para calcular a complexidade de um algoritmo precisamos entender
//quantas vezes ocorre uma chamada da funcao que esta dentro,
//nesse caso a chamada occorre n vezes porque n representa o numero que entra,
//veja que caso entre o num fatorial de 4 a funcao fatorial sera chamada 4 vezes
//logo a complexidade seria O(4) e no geral O(n)
