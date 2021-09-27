//exercicios livro algoritmo cap 10.1

package main

func main() {

}

//10.1_1
func primeira() {
	S := []int{1, 2, 3, 4, 5, 6}
	println()

}

func push(S []int, num int) {
	//append faz o mesmo
	//adicionar um num no ultimo index da pilha
	S[len(S)] = num
	//add um num na pilha
}

func pop(S []int, num int) {
	S[len(S)-1] -= 1
	return S[len(S)]
}

//TODO
//Caso nao coloque retorno nesse tipo de funcao retorna algo?
//tava prestes testar a funcao push
