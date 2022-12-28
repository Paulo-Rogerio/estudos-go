package main

// Como separei em diretorio, preciso informar o pacote que desejo importar
import (
	"fmt"
	"matematica"
)

func main(){
	result := matematica.Soma(1, 1)
	fmt.Printf("O resultado é: %v", result)
	fmt.Printf("O Tipo é: %T", result)
}
