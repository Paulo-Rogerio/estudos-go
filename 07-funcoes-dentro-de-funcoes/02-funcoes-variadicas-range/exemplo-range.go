package main

import "fmt"

func main() {

	// Soma ( 1 + 2 + 3 + 4 + 5 = 15)
	resp := SomaTudo(1, 2, 3, 4, 5)
	fmt.Println(resp)
}

// Funcao recebe um array de inteiros, e retorna um inteiro
// Navegar em todo o x e somar item por item
func SomaTudo(x ...int) int {

	resultado := 0

	// Não quero usar key, entao uso um Blank Indentify represenado por (_,)
	// v represnta meu Valeu que nada mais e que o x
	// Depois Somo e Incremento
	// Loop

	for _, v := range x {
		resultado += v
	}

	return resultado
}
