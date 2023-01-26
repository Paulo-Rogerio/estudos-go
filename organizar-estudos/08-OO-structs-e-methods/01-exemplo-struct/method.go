package main

import "fmt"

// Methodos no Go é baseado em strutc, e uma struct é um type que vc cria.
// Criando uma Struct chamado Carro ( Method chamado Carro )
// Essa struct ( Method ), contem um atributo ( nome ) do tipo String.
type Carro struct {
	nome string
}

// Essa função precisa ser relacionar com a struct ( Method ) criado
// Chamado de Bind / Relacionamento
// Aqui ela é instanciada / atachada.
// Crio a funçao que Instancia c em Carro, e cria o method andar
func (c Carro) andar() {
	fmt.Println(c.nome, "andou !!!")
}

// Crio uma variavel fusca que é uma struct de Carro.
// Definio o nome do carro
// E chamo o method andar.
func main() {

	fusca := Carro{
		nome: "Fusca",
	}

	chevette := Carro{
		nome: "Chevette",
	}

	ferrari := Carro{
		nome: "Ferrari",
	}

	fusca.andar()
	chevette.andar()
	ferrari.andar()

}
