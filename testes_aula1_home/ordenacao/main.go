package main

import (
	"fmt"
)

// O algoritmo mais básico de ordenação de um array é conhecido como "Simple Sort". O mesmo consiste em varrer um array informado diversas vezes, sempre buscando o menor número e o armazenando em outro array.
//     Um exemplo de como o mesmo funciona pode ser analisado abaixo:
//     Array inicial: [ 9, 7, 3, 85, 6, 4, 71, 23, 12 ]

//     // INICIO PROGRAMA
//     DESORDENADO = [ 9, 7, 3, 85, 6, 4, 71, 23, 12 ]
//     ORDENADO = []
//     // PROCURA PELO MENOR NUMERO NO DESORDENADO, O RETIRA DO MESMO E O INSERE NO ARRAY DE ORDENADOS
//     DESORDENADO = [ 9, 7, 85, 6, 4, 71, 23, 12 ]
//     ORDENADO = [ 3 ]
//     // PROCURA PELO MENOR NUMERO NO DESORDENADO, O RETIRA DO MESMO E O INSERE NO ARRAY DE ORDENADOS
//     DESORDENADO = [ 9, 7, 85, 6, 71, 23, 12 ]
//     ORDENADO = [ 3, 4 ]
//     // PROCURA PELO MENOR NUMERO NO DESORDENADO, O RETIRA DO MESMO E O INSERE NO ARRAY DE ORDENADOS
//     DESORDENADO = [ 9, 7, 85, 71, 23, 12 ]
//     ORDENADO = [ 3, 4, 6 ]
//     // ... E ASSIM CONSECUTIVAMENTE ATÉ QUE AO FINAL DO PROCESSO TEMOS O SEGUINTE RESULTADO:
//     DESORDENADO = []
//     ORDENADO = [ 3, 4, 6, 7, 9, 12, 23, 71, 85 ]
//     Array Final: [ 3, 4, 6, 7, 9, 12, 23, 71, 85 ]
//     Dessa maneira, desenvolva um algoritmo que realize a ordenação de um array de números inteiros conforme o processo mencionado anteriormente e a retorne para o usuário.

// receber o array
// varrer o array
// buscando menor num

// armazendando esse num em outro array
//[ 9, 7, 3, 85, 6, 4, 71, 23, 12 ]

func main() {
	sequencia := []int{9, 7, 3, 85, 6, 4, 71, 23, 12}
	fmt.Println(sequencia)
	seq_ordenada := ordenare(sequencia)
	fmt.Println(buscaElemento(90, seq_ordenada))
}

func ordenare(seqDesordenada []int) []int { //O(n²)
	tamanho := len(seqDesordenada) //1
	seqOrdenada := []int{}         //1
	for i := 0; i < tamanho; i++ { //n
		menorNum := int(^uint(0) >> 1)
		indMenorNum := 0
		for ind, num := range seqDesordenada { //n
			if menorNum > num {
				menorNum = num
				indMenorNum = ind
			}
		}
		seqOrdenada = append(seqOrdenada, menorNum)
		seqDesordenada = remove(seqDesordenada, indMenorNum)
	}
	return seqOrdenada
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// func buscaElemento(elemento int, sequencia []int) int { //O(n)
// 	var indice int
// 	for ind, num := range sequencia { //n
// 		if num == elemento { //1
// 			indice = ind
// 		}
// 	}
// 	return indice
// }

func buscaElemento(elemento int, sequencia []int) bool { //O(n)
	tamanho := len(sequencia) - 1
	pontoMedio := tamanho / 2
	fmt.Println(tamanho, pontoMedio)
	fmt.Println(sequencia)
	if sequencia[pontoMedio] == elemento {
		return true
	}
	if tamanho == 0 {
		return false
	}
	if sequencia[pontoMedio] > elemento {
		return buscaElemento(elemento, sequencia[:pontoMedio-1])
	}
	if sequencia[pontoMedio] < elemento {
		return buscaElemento(elemento, sequencia[pontoMedio+1:])
	}
	return false
}
