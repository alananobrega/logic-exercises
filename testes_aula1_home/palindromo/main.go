package main

import (
	"fmt"
	"strings"
)

// Sabemos que existem determindadas frases, quando lidas ao contrário possuem o mesmo "resultado" de quando lidas em seu sentido natural, e a estas frases damos o nome de "Palíndromos".
// Exemplos de Palíndromos podem ser notados abaixo:
// 	"Amor a Roma"
// 	"Anotaram a data da maratona"
// Desenvolva um algoritmo que, a partir de determinada frase, retorne a mesma invertida e informe caso a mesma seja um palíndromo.
// Exemplo:
// inverteFrase("Meu nome é El Kabong") // Resultado = "gnobaK lE é emon ueM", False

//ver se a string contem algum acento se sim indice -2, menos quntos acentos tiver

func main() {
	fmt.Println(isPalindromo("Amor a Roma"))
}

func isPalindromo(sentence string) (string, bool) {
	//receber uma frase
	//receber a string e dividi-la em um array

	split_sentence := strings.Split(sentence, "")
	reverse_sentence := []string{}

	fmt.Println(len(sentence))

	for i := len(sentence) - 2; i >= 0; i-- {
		reverse_sentence = append(reverse_sentence, split_sentence[i])
	}

	fmt.Println(reverse_sentence)
	fmt.Println(split_sentence[1])

	//amor
	//1234

	//4321
	//roma

	//retorna a frase
	// phraseRune := []rune(sentence)
	return sentence, true
}

// func isPalindromo(sentence string) (string, bool) {
// 	//receber uma frase
// 	reverse_sentence := []byte{}

// 	for i := len(sentence) - 1; i >= 0; i-- {
// 		reverse_sentence = append(reverse_sentence, sentence[i])
// 	}

// 	fmt.Printf("%c\n", reverse_sentence)
// 	//amor
// 	//1234

// 	//4321
// 	//roma

// 	//retorna a frase
// 	// phraseRune := []rune(sentence)
// 	return sentence, true
// }
