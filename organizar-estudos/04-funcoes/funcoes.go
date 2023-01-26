package main

import ("fmt")

func main(){
	result := Soma (1, 1)
	fmt.Printf("O resultado da soma é: %v \n", result)
	fmt.Printf("O Tipo da variável result é: %T \n", result)
}

// Criando Funcao Soma, que recebe 2 parametros inteiros e o retorno dela é inteiro
// Se eu não informar o tipo que desejo retornar, não posso informar	 o return

func Soma (a int, b int ) int {
	return a + b
}