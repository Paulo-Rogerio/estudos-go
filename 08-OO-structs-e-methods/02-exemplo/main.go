package main

import "fmt"

type Pessoal struct {
	Nome  string
	Idade int
}

type Categoria struct {
	Nome string
	Pai  *Categoria
}

// Criar um struct que valida se a Struct Categoria é pai
// Ou se ele herda isso de algum lugar
// Retorna um bool
func (c Categoria) Hasparent() bool {
	return c.Pai != nil
}

// Package Tipo String
func (p Pessoal) String() string {
	return fmt.Sprintf("Ola meu nome é %s e tenho %d anos", p.Nome, p.Idade)
}

func main() {

	p := Pessoal{
		Nome:  "Paulo",
		Idade: 40,
	}
	fmt.Println(p)
	fmt.Println(p.Nome)

	fmt.Println("-----")

	cat := Categoria{Nome: "Categoria 1"}
	if !cat.Hasparent() {
		fmt.Println("nao tem parente")
	} else {
		fmt.Println("TEM PAI")
	}

	fmt.Println("-----")

	cat = Categoria{Nome: "Categoria 1", Pai: &Categoria{Nome: "PAI"}}
	if !cat.Hasparent() {
		fmt.Println("nao tem parente")
	} else {
		fmt.Println("TEM PAI")
	}

	// Passando Objeto Completo
	fmt.Println("-----")
	p2 := Pessoal{
		Nome:  "Paulo",
		Idade: 40,
	}
	fmt.Println(p2)

}
