package main

import "fmt"

func main() {

	// Endereco Memoria Computador ( xc00001e0b8 ) <---- a <----- 10
	// Imprimindo endereço memoria
	a := 10
	fmt.Println("Valor Variável (a):", a)
	fmt.Println("Endereço de Memória (a):", &a)
	fmt.Println("===================")

	// Esse endereço fica catalogado em nossos ponteiros
	// Criar uma variavel chamada ponteiro ( * )
	// Essa variavel ponteiro precisa bater com o tipo no caso int.
	// Ponteiro detem o endereço da memoria

	var ponteiro *int = &a

	// Uma vez que tenho o endereço da memoria o Go consegue fazer "Reference"
	// Esse Reference é consultar o valor que está vinculado naquele endereço de Memoria

	fmt.Println("Valor Variável (ponteiro):", ponteiro)
	fmt.Println("Endereço de Memória (ponteiro) (Acessando o Valor Definido na Memória):", *ponteiro)
	fmt.Println("===================")

	// Já que o ponteiro consegue acessar o endereço da memoria
	// Eu posso definir o valor que desejo naquele Endereço de Memoria

	*ponteiro = 50
	fmt.Println("Valor Variável (ponteiro):", ponteiro)
	fmt.Println("Endereço de Memória (ponteiro) (Acessando o Valor Definido na Memória):", *ponteiro)
	fmt.Println("Valor da Variável (a):", *ponteiro)
	fmt.Println("===================")

	// Criando variavel b com o mesmo endereco de memória de a
	b := &a
	fmt.Println("Valor Variável (b):", b)
	fmt.Println("Endereço de Memória (b) (Acessando o Valor Definido na Memória):", *ponteiro)
	fmt.Println("===================")

	// Redefinindo b e consultando as demais variaveis
	*b = 60
	fmt.Println("Valor Variável (b):", b)
	fmt.Println("Endereço de Memória (b) (Acessando o Valor Definido na Memória):", *b)
	fmt.Println("Valor Variável (ponteiro):", ponteiro)
	fmt.Println("Endereço de Memória (ponteiro) (Acessando o Valor Definido na Memória):", *ponteiro)
	fmt.Println("Endereço de Memória (a) (Acessando o Valor Definido na Memória):", a)
	fmt.Println("===================")

}
