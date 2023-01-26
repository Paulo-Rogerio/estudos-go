package main

// Como separei em diretorio, preciso informar o pacote que desejo importar
import (
	"arquivosSeparados/calculadora"
	"fmt"
)

func main(){
	result := calculadora.Soma(1, 1)
	fmt.Printf("O resultado é: %v \n", result)
	fmt.Printf("O Tipo é: %T \n", result)
}

