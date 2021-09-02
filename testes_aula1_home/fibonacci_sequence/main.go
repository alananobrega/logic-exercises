package main

// 1)
//     A sequência Fibonacci é gerada de tal forma que a soma de 2 elementos prévios sempre geram o elemento seguinte. O início da sequência em questão é o seguinte:
//     1, 1, 2, 3, 5, 8, 13, 21, 34, ...
//     (
//         Isto é:
//         1 + 1 = 2
//         2 + 1 = 3
//         3 + 2 = 5
//         E assim consecutivamente
//     )
//     Se somarmos os 7 primeiros termos da sequência, teríamos o seguinte resultado:
//     1 + 1 + 2 + 3 + 5 + 8 + 13 = 33
//     Sendo assim, crie um algoritmo que dado determinado índice, calcule a soma de todos os termos da sequência Fibonacci até o índice em questão.
//     Exemplo:
//     somaFibonnaci(7) // Resultado = 33

import "fmt"

func main() {
	fmt.Println(sumFibonacciSequence(7))
}

//algoritimo que calcule a sequencia de fibonacci
//[1, 1]
//array[0]+[1]
//guarda a soma
//[1]+[2]
//[2]+[3]
//[3]+[4]
//[a]+[a+1]
//a[0.. indice]
//algoritimo que some essa sequencia
func sumFibonacciSequence(indice int) int { //O(n)
	sequence := []int{1, 1}         //1
	sum := 2                        //1
	for i := 1; i < indice-1; i++ { //n
		num := sequence[i-1] + sequence[i] //1
		sequence = append(sequence, num)   //1
		sum += num                         //1
	}
	return sum
}

//Duvidas:
//como refatorar???
