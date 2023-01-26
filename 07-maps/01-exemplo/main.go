package main

import "fmt"

func main() {

	dados := make(map[string]int)
	fmt.Println(dados)

	init := make(map[string]int)
	fmt.Println(init)

	init["goias"] = 150000
	init["senador-canedo"] = 20000

	fmt.Println(init)

	valor, existe := init["goias"]

	if existe {
		fmt.Printf("O valor é: %v e status é: %v \n", valor , existe )
	} else {
		fmt.Printf("O valor é: %v e status é: %v \n", valor , existe )
	}

}



