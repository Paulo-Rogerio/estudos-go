package main

import "fmt"

func main() {

	abacate := 10

	fmt.Println("O endereço de memoria da variavel (abacate) é:", &abacate)
	fmt.Println("O valor da variavel (abacate) é:", abacate)
	fmt.Println("-----------------------------------")

	// Como a funçao pede um ponteiro, preciso informar o endereco de memoria como parametro.
	// Essa função não retorna nada, mas ela altera o valor da variavel contido
	// no endereço de memória da variavel abacate.
	abc(&abacate)

	fmt.Println("O endereço de memoria da variavel (abacate) é:", &abacate)
	fmt.Println("O valor da variavel (abacate) após chamada da função é:", abacate)
	fmt.Println(abacate)

}

// Função recebe um ponteiro inteiro e processa
// Observe que ela não retorna nada, apenas processa
func abc(a *int) {
	*a = 200
}
