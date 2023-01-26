package main

// Como separei em diretorio, preciso informar o pacote que desejo importar
import (
	"fmt"
	"tratandoErros/calculadora"
)

func main() {
	// O erro é provocado quando o denominador for 0
	// result, err := calculadora.Divisao(4, 2)
	result, err := calculadora.Divisao(4, 0)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("O resultado é: %v \n", result)
	fmt.Printf("O Tipo é: %T \n", result)
}
