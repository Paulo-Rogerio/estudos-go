package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// Se a soma for maior que 10 considero erro
	executa, err := Soma(1, 10)

	// Tratando o Error
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(executa)

	//==== Fazendo ByPass do Erro ====
	// Evitar usar assim
	// executa, _ := Soma(1, 10)
	// fmt.Println(executa)

}

// Funcao que recebe 2 inteiros e seu return e inteiro e error
func Soma(a, b int) (int, error) {

	res := (a + b)

	if res > 10 {
		return 0, errors.New("Soma dos fatores é maior que 10")
	}

	return res, nil
}
