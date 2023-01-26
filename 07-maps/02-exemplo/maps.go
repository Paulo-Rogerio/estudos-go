package main

import "fmt"

func main() {
	nome := make([]string, 0)
	nome = append(nome, "paulo")
	nome = append(nome, "camilla")
	fmt.Println("Nomes:", nome)

	// O que é a chave que é o valor
	// Chave ( paulo string )
	// Valor ( 40 int )
	idade := make(map[string]int)
	idade["paulo"] = 40
	idade["camilla"] = 33
	fmt.Println("Agrupado:", idade)

	fmt.Println("Idade Paulo:", idade["paulo"])

	val, ok := idade["Lucas"]
	fmt.Println("Existe Lucas?:", val, ok)
}
