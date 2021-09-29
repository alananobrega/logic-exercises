package main

import "fmt"

// Muitas vezes, quando estamos analisando determinada lista, surge-se a necessidade de remover os valores únicos dessa lista. E a essa nova lista de valores gerados, damos o nome de "valores únicos" (Unique).
// Um exemplo dessa transformação pode ser observado abaixo:
// Array:  [ 2, 3, 8, 9, 3, 2, 6, 4, 5, 1, 5, 5, 6, 5, 4, 7, 5, 2 ]
// Unique: [ 2, 3, 8, 9, 6, 4, 5, 1, 7 ]
// Dito isto, desenvolva um algoritmo que remova os valores repetidos de um array e o retorne para o usuário.
// Lembre-se os elementos não precisam estar na ordem!

func main() {
	seq := []int{2, 3, 8, 9, 3, 2, 6, 4, 5, 1, 5, 5, 6, 5, 4, 7, 5, 2}
	fmt.Println(unic(seq))
}

// recebe o array
//pega um elemento e compara com o proximo
//[a,b,a,c,d]
//i[0]==i[1]
//i[0]==i[2]
//i[0]==i[3]
//i[0]==i[4]

//i[1]==i[2]
//i[1]==i[3]
//i[1]==i[4]

//se i[0] == i[2]
//[a,b,c,d]
//indice - 1
//{0, 1, 2, 3, 4}
func unic(seq []int) []int {//O(n * log n)
	for ind, v := range seq {//n
		for i := ind + 1; i < len(seq); i++ {//log n
			if v == seq[i] {
				seq = remove(seq, i)
				i--
			}
		}

		if ind+2 == len(seq) {
			if v == seq[ind+1] {
				seq = remove(seq, ind)
			}
		}
	}
	return seq
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
