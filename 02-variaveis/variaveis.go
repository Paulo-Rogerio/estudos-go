package main

import (
	"fmt"
	"reflect"
)

// Tipando Variável
//var a string

func main (){

	// Uma vez declarado , não pode ser alterado
	const b = "Rogerio"

	// Declarando e deixando o Go inferir o tipo
	a := "Paulo "
	fmt.Printf("Bem vindo ao mundo Go %v \n", (a + b) )
	fmt.Printf("O tipo dessa variável é: %v \n", reflect.TypeOf(a))

}