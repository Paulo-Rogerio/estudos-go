package calculadora

import "errors"

// Como estou criando uma funcao fora do arquivo principal
// Sou obrigado a iniciar o nome da funcao com a primeira letra Maiuscula

// Meu dominador não pode ser 0, vamos considerar isso como Erro
func Divisao(a, b int) (int, error) {

	// Devo importar o pacote Errors
	if b == 0 {
		return 0, errors.New("Denominador não pode ser 0")
	}

	// Entregando os 2 fatores
	return (a / b), nil

}
