package main

import (
	"fmt"
)

func main (){

	lista := []int{1,2,3,4}
	fmt.Println(lista)
	fmt.Println(lista[0])
	fmt.Println("Tamanho da Lista:", len(lista))
	
	fmt.Println("========")

	lista = append(lista, 8)
	fmt.Println(lista)

	fmt.Println("========")
	fmt.Println("Tamanho da Lista:", len(lista))

}