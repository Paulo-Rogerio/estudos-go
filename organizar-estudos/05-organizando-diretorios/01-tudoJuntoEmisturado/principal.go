package main

// Como separei em diretorio, preciso informar o pacote que desejo importar
import (
	"fmt"
)

func main(){
	result := soma(1, 1)
	fmt.Printf("O resultado é: %v \n", result)
	fmt.Printf("O Tipo é: %T \n", result)
}

// Como estou criando uma funcao dentro do mesmo arquivo, posso usar
// letra minuscula no nome da funcao
func soma (a, b int ) int {
	return a + b
}
