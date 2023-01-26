package main

import "fmt"

func main() {

	// A variavel resultado recebe 2 funções func(x ...int) / func
	// E seu retorno é um inteiro
	resultado := func(x ...int) func() int {

		resposta := 0

		// Soma todos os valores passados e armazena na variavel resposta
		// Resultado 216
		for _, v := range x {
			resposta += v
		}

		// Multiplica 216 * 216
		return func() int {
			return resposta * resposta
		}
	}

	// Isso imprime o Objeto
	fmt.Println(resultado(54, 54, 54, 54))

	// Imprimindo a Multiplicacao do resultado
	fmt.Println(resultado(54, 54, 54, 54)())
}
