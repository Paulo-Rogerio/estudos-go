package main

import "fmt"

func main() {

	lista := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	um_a_cinco := 0
	seis_a_dez := 0

	for i := range lista {
		if lista[i] <= 5 {
			um_a_cinco += lista[i]
		} else {
			seis_a_dez += lista[i]
		}
	}

	fmt.Println("A soma dos itens de 1 a 5 é :", um_a_cinco)
	fmt.Println("A soma dos itens de 6 a 10 é:", seis_a_dez)
}
