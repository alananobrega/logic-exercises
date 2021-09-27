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
	fmt.Println(isPalindromo("Meu nome é El Kabong"))
}

func isPalindromo(sentence string) (string, bool) { //O(n)

	split_sentence := strings.Split(sentence, "") //1
	reverse_sentence := []string{}                //1

	fmt.Println(split_sentence) //1

	for _, caractere := range split_sentence { //n
		reverse_sentence = append(reverse_sentence, "x")
		copy(reverse_sentence[1:], reverse_sentence[0:])
		reverse_sentence[0] = caractere

		fmt.Println(reverse_sentence)
	}
	return sentence, true
}
